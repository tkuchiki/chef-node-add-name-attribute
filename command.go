package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "overwrite, o",
			Usage: "set overwrite",
		},
	}
}

func Command(c *cli.Context) {
	for _, location := range c.Args() {
		if FileNotExists(location) {
			fmt.Println(location + " is not found")
			continue
		}

		ch := GoWalk(location)

		for filepath := range ch {
			jsonData := DecodeJson(filepath)
			indentedJson := EncodeJson(jsonData, filepath)
			WriteJSON(filepath, indentedJson)
		}
	}
}
