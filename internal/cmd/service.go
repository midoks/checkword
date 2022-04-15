package cmd

import (
	"github.com/urfave/cli"

	"github.com/midoks/checkword/internal/app/libs"
	_ "github.com/midoks/checkword/internal/app/routers"
)

var Service = cli.Command{
	Name:        "service",
	Usage:       "This command starts all services",
	Description: `Start Check Word services`,
	Action:      runAllService,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runAllService(c *cli.Context) error {

	libs.Init()

	return nil
}
