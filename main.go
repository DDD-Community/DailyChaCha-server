package main

import (
	"github.com/DDD-Community/DailyChaCha-server/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	handler.ApplyHandler(e)

	e.Logger.Fatal(e.Start(":1323"))
}
