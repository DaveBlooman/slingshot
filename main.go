package main

import (
	"os"

	"github.com/DaveBlooman/slingshot/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "DaveBlooman"
	app.Email = "david.blooman@gmail.co.uk"
	app.Usage = "Uploads a directory of static assets to S3"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
