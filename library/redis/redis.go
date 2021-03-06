package redis

import (
	"context"
	"fmt"
	"gin_demo/utils"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var once = sync.Once{}
var CurrentClient *redis.Client

type Lock struct {
	Redis       *redis.Client
	Key         string
	Expiration  time.Duration
	CheckCancel chan bool
	LockId      string
	Mu          sync.Mutex
}

func GetLock(lockName string, expiration time.Duration) *Lock {
	lockId := utils.RandToken(10)

	return &Lock{
		Redis:      GetRedisClient(),
		Key:        "lock:" + lockName,
		Expiration: expiration,
		LockId:     lockId,
		Mu:         sync.Mutex{},
	}
}

func newRedisClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//read env variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// alternatively, read env var from config file
	// config := config.GetConfig()
	// redisHost := config.Redis.Host
	// redisPassword := config.Redis.Password

	conf := &redis.Options{
		Addr:     redisHost,
		DB:       1,
		Password: redisPassword,
	}
	newRedisClient := redis.NewClient(conf)
	resp := newRedisClient.Ping(ctx)
	if resp.Err() != nil {
		panic(resp.Err())
	}
	CurrentClient = newRedisClient

	// CurrentClient.Set(ctx, "gbg", 123, 30*time.Second)
}

// using sync.once to only init redis client once
func GetRedisClient() *redis.Client {
	if CurrentClient == nil {
		once.Do(func() { newRedisClient() })
	}
	return CurrentClient
}

// return lockId which used to identify locks
// this lock is blocked lock.
/* usage example:
lock := GetLock(lockName, 5*time.Second)
// lockId := lock.AcquireLock(c) // attemp to acquire a lock which has expiration time of 5 seconds
// defer lock.Release(c, lockName, lockId)
*/
func (lock *Lock) AcquireLock(c *gin.Context) string {
	// set tick interval to 100ns, which try to acquire lock every 100ns
	tick := time.NewTicker(time.Nanosecond * 100)

	// set time out to 10 second
	timer := time.NewTimer(10 * time.Second)

	for {
		select {

		// time out
		case <-timer.C:
			timer.Stop()
			return ""

		case <-tick.C:
			setNxCmd := lock.Redis.SetNX(c, lock.Key, lock.LockId, lock.Expiration)
			if ok, _ := setNxCmd.Result(); ok {
				go lock.checkLockIsRelease()
				return lock.LockId
			}
		}
	}
}

// check if lock is released, if not, renew lock
func (lock *Lock) checkLockIsRelease() {
	for {
		duration := time.Millisecond * time.Duration(lock.Expiration.Milliseconds()-lock.Expiration.Milliseconds()/10)
		checkCxt, _ := context.WithTimeout(context.Background(), duration)
		lock.CheckCancel = make(chan bool)
		select {
		case <-checkCxt.Done():
			// keep renewing lock until lock got released
			isLockRenewed := lock.renewLock()
			if !isLockRenewed {
				return
			}
		case <-lock.CheckCancel:
			fmt.Println("----lock has been released----")
			return
		}

	}
}

// renew lock, if success return true, otherwise return false
func (lock *Lock) renewLock() bool {
	lock.Mu.Lock()
	defer lock.Mu.Unlock()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := lock.Redis.Exists(ctx, lock.Key).Result()
	cancel()
	if err != nil {
		return false
	}
	if res == 1 {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		ok, err := lock.Redis.Expire(ctx, lock.Key, lock.Expiration).Result()
		cancel()
		if err != nil {
			return false
		}
		if ok {
			fmt.Println("--- lock is renewed-----")
			return true
		}
	}
	return false

}

func (lock *Lock) GetKey(c *gin.Context) string {
	return lock.Key
}

func (lock *Lock) GetRedis(c *gin.Context) *redis.Client {
	return lock.Redis
}

func (lock *Lock) Release(c *gin.Context) bool {
	lock.Mu.Lock()
	defer lock.Mu.Unlock()

	// alternatively, we can use redis pipeline
	// timer := time.NewTimer(50 * time.Second)
	// txf := func(tx *redis.Tx) error {
	// 	getCmd := tx.Get(c, key)
	// 	fn := func(pipe redis.Pipeliner) error {
	// 		if getCmd.Val() == lockId {
	// 			pipe.Del(c, key)
	// 		}
	// 		return nil
	// 	}
	// 	_, err := tx.TxPipelined(c, fn)
	// 	return err
	// }

	// for {
	// 	select {
	// 	case <-timer.C:
	// 		timer.Stop()
	// 		return false
	// 	default:
	// 		err := lock.redis.Watch(c, txf, key)
	// 		if err == nil {
	// 			return true
	// 		} else if err == redis.TxFailedErr {
	// 			// something wrong, we either lost the key or an unexpected thing happened, just try again
	// 			continue
	// 		} else {
	// 			return false
	// 		}
	// 	}
	// }

	// or we can use lua script to implement atomic operation
	const luaScript = `
	if redis.call('get', KEYS[1])==ARGV[1] then
		return redis.call('del', KEYS[1])
	else
		return 0
	end
	`

	script := redis.NewScript(luaScript)
	result, err := script.Run(c, lock.Redis, []string{lock.Key}, lock.LockId).Result()
	fmt.Println("-----lock release-----", err)
	if err == nil {
		if result.(int64) == 1 {
			// lock released (key deleted successfully)
			lock.CheckCancel <- true
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
