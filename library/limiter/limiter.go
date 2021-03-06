package limiter

import (
	"gin_demo/library/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type LockOptions struct {
	Attempts        int
	Ip              string
	Duration        time.Duration
	EndPoint        string
	BlockedDuration time.Duration
}

type LockResult struct {
	Block         bool
	TimeLeft      int
	AttemptsCount int
	AttemptsLeft  int
}

func RateLimiter(c *gin.Context, opts LockOptions) (LockResult, error) {
	lockName := opts.Ip + ":" + opts.EndPoint
	blockedKey := "isBlocked:" + lockName
	lock := redis.GetLock(lockName, opts.Duration)

	// example usage of lock acquire and lock release
	// lockId := lock.AcquireLock(c) // attemp to acquire a lock which has expiration time of 5 seconds
	// defer lock.Release(c)
	key := lock.Key // unique key for
	getCmd := lock.Redis.Get(c, key)
	AttemptsLeft, _ := getCmd.Int()
	timeLeft := lock.Redis.PTTL(c, key).Val() // in milliseconds

	if AttemptsLeft <= 0 && timeLeft < 0 {
		// either key expired or first time using
		AttemptsLeft := opts.Attempts - 1
		setResult := lock.Redis.Set(c, key, AttemptsLeft, opts.Duration)
		if err := setResult.Err(); err != nil {
			return LockResult{}, err
		}
		result := LockResult{
			AttemptsLeft:  AttemptsLeft,
			TimeLeft:      int(opts.Duration.Milliseconds()),
			AttemptsCount: 1,
			Block:         false,
		}
		return result, nil // allow request go through
	} else if AttemptsLeft <= 0 && timeLeft >= 0 {
		// block request

		isAlreadyBlocked, _ := lock.Redis.Get(c, blockedKey).Int()
		if isAlreadyBlocked < 0 {
			lock.Redis.Set(c, blockedKey, 1, opts.Duration)
			lock.Redis.PExpire(c, blockedKey, opts.BlockedDuration)
		}
		return LockResult{
			AttemptsLeft:  0,
			TimeLeft:      int(timeLeft),
			AttemptsCount: 1,
			Block:         true,
		}, nil
	} else {
		decrCmd := lock.Redis.Decr(c, key)
		AttemptsLeft := int(decrCmd.Val())
		return LockResult{
			AttemptsLeft:  AttemptsLeft,
			TimeLeft:      int(timeLeft),
			AttemptsCount: opts.Attempts - AttemptsLeft,
			Block:         false,
		}, nil
	}

}
