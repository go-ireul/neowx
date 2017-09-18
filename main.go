package main

import (
	"os"

	"ireul.com/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "neowx"
	app.Usage = "Neo Wechat MP Gateway"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "config.yaml",
			Usage: "config file",
		},
	}
	app.Commands = []cli.Command{
		webCommand,
		workerCommand,
	}
	app.Run(os.Args)
}
