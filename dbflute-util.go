package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "dbflute-util"
	app.Version = Version
	app.Usage = ""
	app.Author = "wata"
	app.Email = "nagasawa@junkapp.com"
    app.Commands = Commands
    
	app.Run(os.Args)
}


