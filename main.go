package main

import (
	"github.com/WarlordDev/configuration_service/router"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	router.Router(e)
	e.Logger.Fatal(e.Start(":13200"))
}
