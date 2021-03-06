package godo

import (
	"bytes"
	"fmt"
	"os"

	"github.com/DaveBlooman/slingshot/Godeps/_workspace/src/github.com/mgutz/ansi"
)

type fileWrapper struct {
	file      *os.File
	buf       *bytes.Buffer
	readLines string

	recorder *bytes.Buffer

	// Adds color to stdout & stderr if terminal supports it
	colorStart string
	colorReset string
}

func newFileWrapper(file *os.File, recorder *bytes.Buffer, color string) *fileWrapper {
	streamer := &fileWrapper{
		file:       file,
		buf:        bytes.NewBuffer([]byte("")),
		recorder:   recorder,
		colorReset: ansi.ColorCode("reset"),
		colorStart: color,
	}

	return streamer
}

func (l *fileWrapper) Write(p []byte) (n int, err error) {
	if n, err = l.recorder.Write(p); err != nil {
		return
	}

	err = l.out(string(p))
	return
}

func (l *fileWrapper) WriteString(s string) (n int, err error) {
	if n, err = l.recorder.WriteString(s); err != nil {
		return
	}

	err = l.out(s)
	return
}

func (l *fileWrapper) Close() error {
	l.buf = bytes.NewBuffer([]byte(""))
	return nil
}

func (l *fileWrapper) out(str string) (err error) {

	if l.colorStart != "" {
		str = l.colorStart + str + l.colorReset
	}

	fmt.Fprint(l.file, str)
	return nil
}
