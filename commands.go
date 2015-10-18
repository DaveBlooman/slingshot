package main

import (
	"fmt"
	"os"

	"github.com/DaveBlooman/slingshot/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DaveBlooman/slingshot/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{

	{
		Name:   "upload",
		Usage:  "",
		Action: command.CmdUpload,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "bucket, b",
				Usage: "Bucket name",
			},
			cli.StringFlag{
				Name:  "path, p",
				Usage: "Objects path",
			},
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Region",
			},
			cli.StringFlag{
				Name:  "directory, d",
				Usage: "Directory",
			},
		},
	},

	{
		Name:   "current",
		Usage:  "",
		Action: command.CmdCurrent,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "bucket, b",
				Usage: "Bucket name",
			},
			cli.StringFlag{
				Name:  "path, p",
				Usage: "Objects path",
			},
			cli.StringFlag{
				Name:  "region, r",
				Usage: "Region",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
