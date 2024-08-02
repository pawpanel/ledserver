package leds

// Config stores configuration information for the LEDs.
type Config struct {
	Count int `yaml:"count"`
	Pin   int `yaml:"pin"`
}
