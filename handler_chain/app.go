package main

import (
	"log"

	"github.com/go-amwk/core"
	"github.com/go-amwk/web"
)

func main() {
	app := web.Default()

	app.Use(middleware1).
		Use(middleware2).
		Use(middleware3).
		Use(handler).
		Use(handler_never_reached)

	app.Start()
}

func middleware1(c core.Context) error {
	log.Println("middleware1 start")
	err := c.Next()
	log.Println("middleware1 end")
	return err
}

func middleware2(c core.Context) error {
	log.Println("middleware2 start")
	err := c.Next()
	log.Println("middleware2 end")
	return err
}

func middleware3(c core.Context) error {
	log.Println("middleware3")
	return nil
}

func handler(c core.Context) error {
	log.Println("handler")
	c.Write([]byte("Hello, World!\n"))
	c.Abort()
	return nil
}

func handler_never_reached(c core.Context) error {
	log.Println("! This should never be printed!")
	return nil
}
