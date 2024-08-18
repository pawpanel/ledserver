package effects

import (
	"encoding/json"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	var strDuration string
	if err := json.Unmarshal(b, &strDuration); err != nil {
		return err
	}
	var err error
	v, err := time.ParseDuration(strDuration)
	if err != nil {
		return err
	}
	*d = Duration(v)
	return nil
}
