package leds

// Config stores configuration information for the LEDs.
type Config struct {
	Pin     int `yaml:"pin"`
	Count   int `yaml:"count"`
	Regions []struct {
		Name   string `yaml:"name"`
		Blocks []struct {
			Offset  int  `yaml:"offset"`
			Count   int  `yaml:"count"`
			Reverse bool `yaml:"reverse"`
		}
	} `yaml:"regions"`
}
