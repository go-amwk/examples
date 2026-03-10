package main

import (
	"fmt"
	"sync/atomic"

	"github.com/go-amwk/core"
	"github.com/go-amwk/web"
)

var cnt atomic.Int64

func main() {
	app := web.Default()

	app.Use(middleware).Use(handler)

	app.Start()
}

func middleware(c core.Context) error {
	c.Set("count", cnt.Add(1))
	return c.Next()
}

func handler(c core.Context) error {
	count, _ := c.Get("count")
	c.Write([]byte(fmt.Sprintf("Count: %d\n", count)))
	return nil
}
