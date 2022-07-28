package utils

import (
	"encoding/json"
	"fmt"
	"gin_demo/library/logger"
	"time"

	"go.uber.org/zap"
)

type MyTime time.Time

func (mt *MyTime) UnmarshalJSON(bs []byte) error {
	var timestamp int64
	err := json.Unmarshal(bs, &timestamp)
	if err != nil {
		return err
	}

	*mt = MyTime(time.Unix(timestamp/1000, timestamp%1000*1e6))
	fmt.Println("---mt---", *mt)
	return nil
}

func (mt MyTime) TimeStamp() (ts int64) {
	ts = time.Time(mt).Unix()
	return
}

func (mt MyTime) MarshalJSON() ([]byte, error) {
	timestamp := time.Time(mt).UnixNano() / 1e6
	logger.ZapLog_V1.Info("timestamp: ", zap.Int64("time", timestamp))
	return json.Marshal(timestamp)
}
