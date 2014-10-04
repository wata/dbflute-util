package main

import (
	"log"
	"os"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
    commandInit,
    commandDoc,
    commandLoadDataReverse,
    commandReplaceSchema,
}


var commandInit = cli.Command{
	Name:  "init",
	Usage: "",
	Description: `
`,
	Action: doInit,
}

var commandInit = cli.Command{
	Name:  "upgrade",
	Usage: "",
	Description: `
`,
	Action: doUpgrade,
}

var commandDoc = cli.Command{
	Name:  "doc",
	Usage: "",
	Description: `
`,
	Action: doDoc,
}

var commandLoadDataReverse = cli.Command{
	Name:  "load-data-reverse",
	Usage: "",
	Description: `
`,
	Action: doLoadDataReverse,
}

var commandReplaceSchema = cli.Command{
	Name:  "replace-schema",
	Usage: "",
	Description: `
`,
	Action: doReplaceSchema,
}


func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func doInit(c *cli.Context) {
}

func doUpgrade(c *cli.Context) {
}

func doDoc(c *cli.Context) {
}

func doLoadDataReverse(c *cli.Context) {
}

func doReplaceSchema(c *cli.Context) {
}
