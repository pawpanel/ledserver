package server

import (
	"encoding/json"
	"time"
)

type stringDuration time.Duration

func (s *stringDuration) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var err error
	d, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*s = stringDuration(d)
	return nil
}
