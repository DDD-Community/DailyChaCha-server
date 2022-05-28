package handler

import (
	"github.com/labstack/echo/v4"
)

func ApplyHandler(e *echo.Echo) {
	// 회원가입 API
	e.POST("/api/signup", signUp())

	// 로그인 API(현재는 테스트용)
	e.POST("/api/signin", signIn())

	// 목데이터로 테스트
	e.GET("/api/getlist", healthCheck())
}
