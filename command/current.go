package command

import (
	"fmt"

	"github.com/DaveBlooman/slingshot/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DaveBlooman/slingshot/storage"
)

func CmdCurrent(c *cli.Context) {

	client, err := storage.List(c.String("region"), c.String("bucket"), c.String("path"))
	if err != nil {
		fmt.Println("Unable to complete request")
	}
	for _, element := range client.Contents {
		fmt.Println(*element.Key)
	}
}
