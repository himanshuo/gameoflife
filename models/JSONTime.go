package models

import (
	"fmt"
	"strconv"
	"time"
)

type JSONTime time.Time

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	//cast as time.Time and call Unix() function
	//Unix() -> returns t as a Unix time (num seconds since epoch)
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)
	//convert to []byte which json thinks of as an array
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	//use the []byte as a string
	//convert string to integer Unix time
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	//convert the unix time from int to int64
	//convert int64 unix time into local time.
	//0 sec offset from epoch
	*t = JSONTime(time.Unix(int64(ts), 0))

	return nil
}

func (t *JSONTime) String() string {
	return time.Time(*t).String()
}