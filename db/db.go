package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	USER := os.Getenv("DBUSER") // DB 유저명
	PASS := os.Getenv("DBPASS") // DB 유저의 패스워드
	HOST := os.Getenv("DBHOST")
	PORT := 3306
	DBNAME := os.Getenv("DBNAME") // 사용할 DB 명을 입력
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		USER,
		PASS,
		HOST,
		PORT,
		DBNAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
