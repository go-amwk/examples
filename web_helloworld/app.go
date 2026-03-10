package main

import (
	"github.com/go-amwk/core"
	"github.com/go-amwk/web"
)

func main() {
	app := web.Default()

	app.Use(func(ctx core.Context) error {
		ctx.Write([]byte("Hello, World!"))
		return nil
	})

	app.Start()
}
