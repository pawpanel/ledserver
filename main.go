package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pawplace/ledserver/config"
	"github.com/pawplace/ledserver/leds"
	"github.com/pawplace/ledserver/server"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func main() {
	app := &cli.App{
		Name:  "ledserver",
		Usage: "HTTP server for controlling ws2812b LEDs",
		Flags: []cli.Flag{configFlag},
		Commands: []*cli.Command{
			installCommand,
		},
		Action: func(c *cli.Context) error {

			// Load the configuration file
			f, err := os.Open(c.String("config"))
			if err != nil {
				return err
			}
			defer f.Close()

			// Parse the configuration file
			cfg := &config.Config{}
			if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
				return err
			}

			// Initialize the LEDs
			l, err := leds.New(&cfg.Leds)
			if err != nil {
				return err
			}

			// Create the server
			s := server.New(&cfg.Server, l)
			defer s.Close()

			// Wait for SIGINT or SIGTERM
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			<-sigChan

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %s\n", err.Error())
	}
}
