package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/midoks/checkword/internal/cmd"
)

const Version = "0.0.2"
const AppName = "checkword"

func main() {

	app := cli.NewApp()
	app.Name = AppName
	app.Version = Version
	app.Usage = "A simple Check Word service"
	app.Commands = []cli.Command{
		cmd.Service,
	}

	app.Run(os.Args)
}
