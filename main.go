package main

import (
	"math/rand"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	_ "github.com/DDD-Community/DailyChaCha-server/docs"
)

// @title DailyChaCha Sample Swagger API
// @version 1.0
// @host ec2-13-209-98-22.ap-northeast-2.compute.amazonaws.com
// @BasePath /api
func main() {

	// godotenv는 로컬 개발환경에서 .env를 통해 환경변수를 읽어올 때 쓰는 모듈이다.
	// 프로덕션 환경에서는 필요하지 않음.
	if err := godotenv.Load(".env"); err != nil {
		logrus.Info(err)
	}
	rand.Seed(time.Now().UnixNano())
	e := echo.New()

	db := db.Connect()
	handler.ApplyHandler(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
