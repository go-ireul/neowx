package main

import "ireul.com/cli"

// workerCommand 用来启动 worker 服务
var workerCommand = cli.Command{
	Name:   "worker",
	Usage:  "start the worker server",
	Action: execworkerCommand,
}

func execworkerCommand(c *cli.Context) (err error) {
	return
}
