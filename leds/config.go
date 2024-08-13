package leds

// Config stores configuration information for the LEDs.
type Config struct {
	Pin         int `yaml:"pin"`
	Count       int `yaml:"count"`
	RefreshRate int `yaml:"refresh_rate"`
	Regions     []struct {
		Name   string `yaml:"name"`
		Title  string `yaml:"title"`
		Blocks []struct {
			Offset  int  `yaml:"offset"`
			Count   int  `yaml:"count"`
			Reverse bool `yaml:"reverse"`
		}
	} `yaml:"regions"`
}
