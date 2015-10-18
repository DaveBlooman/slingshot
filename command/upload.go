package command

import (
	"bytes"
	"fmt"
	"log"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/DaveBlooman/slingshot/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DaveBlooman/slingshot/storage"
)

func CmdUpload(c *cli.Context) {

	fileList := getFiles(c.String("directory"))

	for _, f := range fileList {
		file, err := os.Open(f)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		fileInfo, _ := file.Stat()
		var size int64 = fileInfo.Size()

		buffer := make([]byte, size)

		file.Read(buffer)

		fileBytes := bytes.NewReader(buffer)

		fileName := strings.Replace(file.Name(), c.String("directory"), "", -1)
		fileType := mime.TypeByExtension(path.Ext(f))
		path := fmt.Sprintf("/%s/%s", c.String("path"), fileName)

		_, putErr := storage.Put(c.String("region"), c.String("bucket"), path, fileType, fileBytes, size)

		if putErr != nil {
			fmt.Println("%s", putErr)
			return
		}
		fmt.Println("Uploading file:" + file.Name())
	}
	fmt.Println(fmt.Sprintf("Upload of directory %s complete", c.String("directory")))
}

func getFiles(filesDir string) []string {

	fileList := []string{}
	err := filepath.Walk(filesDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Error:", err)
		}
		if f.IsDir() {
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return fileList
}
