package main

import (
	"github.com/go-amwk/core"
	"github.com/go-amwk/router"
	"github.com/go-amwk/web"
)

func main() {
	app := web.Default()
	r := router.New()

	r.Any("/ping", func(ctx core.Context) error {
		_, err := ctx.Write([]byte("pong"))
		return err
	}).GET("/secret", func(ctx core.Context) error {
		_, err := ctx.Write([]byte("Secret data"))
		return err
	})

	app.Use(r.Route())

	app.Start()
}
