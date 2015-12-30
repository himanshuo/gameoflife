package models

import (
	"fmt"
	"strconv"
	"time"
)

type JSONTime time.Time

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = JSONTime(time.Unix(int64(ts), 0))

	return nil
}

func (t *JSONTime) String() string {
	return time.Time(*t).String()
}