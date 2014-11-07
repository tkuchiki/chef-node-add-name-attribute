package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

var quiet bool

func Print(msg string) {
	if !quiet {
		fmt.Println(msg)
	}
}

func Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "quiet, q",
			Usage: "quiet mode",
		},
	}
}

func Command(c *cli.Context) {
	quiet = c.Bool("quiet")

	for _, location := range c.Args() {
		if FileNotExists(location) {
			Print(location + " is not found")
			continue
		}

		ch := GoWalk(location)

		for fpath := range ch {
			jsonData := DecodeJson(fpath)

			if _, ok := jsonData.(map[string]interface{})["name"]; ok {
				Print(fpath + ": already exists name attribute")
				continue
			} else {
				jsonData.(map[string]interface{})["name"] = FileWithoutExt(fpath)
			}

			indentedJson := EncodeJson(jsonData)
			WriteJSON(fpath, indentedJson)

			Print(fpath + ": Added name attribute")
		}
	}
}
