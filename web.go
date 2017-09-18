package main

import (
	"io/ioutil"
	"log"

	"ireul.com/cli"
	"ireul.com/neowx/routes"
	"ireul.com/neowx/types"
	"ireul.com/web"
	"ireul.com/yaml"
)

// webCommand 用来启动 Web 服务
var webCommand = cli.Command{
	Name:   "web",
	Usage:  "start the web server",
	Action: execWebCommand,
}

func execWebCommand(c *cli.Context) (err error) {
	// decode Config
	bytes, err := ioutil.ReadFile(c.GlobalString("config"))
	if err != nil {
		log.Fatal(err)
		return
	}
	cfg := types.Config{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%v\n", cfg)
	// build web.Web
	w := web.New()
	w.Use(web.Logger())
	w.Use(web.Recovery())
	w.Use(web.Static("public"))
	w.Use(web.Renderer())
	// mount
	routes.Mount(w)
	// run web instance
	w.Run(cfg.Host, cfg.Port)
	return
}
