package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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
	log.Println(time.Time(mt).UnixNano())
	return json.Marshal(timestamp)
}
