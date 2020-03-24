package main

import (
	"bitbucket.org/warlorddevteam/conf/router"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	router.Router(e)
	e.Logger.Fatal(e.Start(":13200"))
}
