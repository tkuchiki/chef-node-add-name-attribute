package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "chef-node-add-name-attribute"
	app.Usage = "add name attribute to node json file"
	app.Flags = Flags()
	app.Action = Command
	app.Run(os.Args)
}
