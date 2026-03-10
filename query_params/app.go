package main

import (
	"fmt"

	"github.com/go-amwk/core"
	"github.com/go-amwk/web"
)

func main() {
	app := web.Default()

	app.Use(func(ctx core.Context) error {
		name := ctx.Query("name")
		if name == "" {
			name = "World"
		}
		ctx.Write([]byte(fmt.Sprintf("Hello, %s!\n", name)))
		return nil
	})

	app.Start()
}
