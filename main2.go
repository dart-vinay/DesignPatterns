package main

import (
	"DesignPatterns/Creational"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		Method()
		Method2()
		return nil
	})

	e.Logger.Fatal(e.Start(":20001"))
}

func Method() {
	singletonPattern := Creational.GetDBConnection("val")
	log.Info(singletonPattern)
}

func Method2() {
	singletonPattern := Creational.GetDBConnection("val2")
	log.Info(singletonPattern)
}
