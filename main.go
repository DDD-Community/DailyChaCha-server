package main

import (
	"github.com/DDD-Community/DailyChaCha-server/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	// godotenv는 로컬 개발환경에서 .env를 통해 환경변수를 읽어올 때 쓰는 모듈이다.
	// 프로덕션 환경에서는 필요하지 않음.
	if err := godotenv.Load(".env"); err != nil {
		logrus.Info(err)
	}
	e := echo.New()
	handler.ApplyHandler(e)

	e.Logger.Fatal(e.Start(":1323"))
}
