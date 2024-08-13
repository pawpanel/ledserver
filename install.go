package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli/v2"
)

const (
	systemdUnitFileName = "/lib/systemd/system/ledserver.service"
	systemdUnitFile     = `[Unit]
Description=ledserver

[Service]
ExecStart={{.self_path}} --config {{.config_path}}

[Install]
WantedBy=default.target
`

	configDefaultFileName = "/etc/pawplace/ledserver.yaml"
	configFile            = `# TODO: use this file to configure the application

leds:
  pin: 18
  count: 48
  regions:
    - name: all
      title: All
      blocks:
        - offset: 0
          count: 48
server:
  addr: 127.0.0.1:16123
`
)

var (
	configFlag = &cli.StringFlag{
		Name:    "config",
		Value:   configDefaultFileName,
		EnvVars: []string{"CONFIG"},
		Usage:   "filename of configuration file",
	}
	installCommand = &cli.Command{
		Name:   "install",
		Usage:  "install the application as a local service",
		Flags:  []cli.Flag{configFlag},
		Action: install,
	}
)

func writeTemplate(
	filename string,
	content string,
	perm fs.FileMode,
	data any,
) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}
	t, err := template.New("").Parse(content)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, data)
}

func install(c *cli.Context) error {

	// Determine the full path to the executable
	p, err := os.Executable()
	if err != nil {
		return err
	}

	// Write the unit file
	if err := writeTemplate(
		systemdUnitFileName,
		systemdUnitFile,
		0644,
		map[string]interface{}{
			"self_path":   p,
			"config_path": c.String("config"),
		},
	); err != nil {
		return err
	}

	// Write the configuration file
	if err := writeTemplate(
		c.String("config"),
		configFile,
		0600,
		nil,
	); err != nil {
		return err
	}

	fmt.Println("Service installed!")
	fmt.Println("")
	fmt.Println("Be sure to edit the configuration file:")
	fmt.Println(c.String("config"))
	fmt.Println("")
	fmt.Println("To enable the service and start it, run:")
	fmt.Println("  systemctl enable ledserver")
	fmt.Println("  systemctl start ledserver")

	return nil
}
