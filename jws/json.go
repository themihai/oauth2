package jws

import (
	"strconv"
	"time"
)

// UnixTime implements encoding/json Marshaler
func (t UnixTime) Marshal() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t UnixTime) Unmarshal(b []byte) error {
	i, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		return err
	}
	t = UnixTime(time.Unix(i, 0))
	return nil
}
