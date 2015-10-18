package main

import . "github.com/DaveBlooman/slingshot/Godeps/_workspace/src/gopkg.in/godo.v1"

func tasks(p *Project) {
	p.Task("build", func(c *Context) error {
		return Run("go build")
	}).Watch("**/*.go")
}

func main() {
	Godo(tasks)
}
