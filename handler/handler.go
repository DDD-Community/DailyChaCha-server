package handler

import (
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func ApplyHandler(e *echo.Echo) {
	// 회원가입 API
	e.POST("/api/sign-up", signUp())

	// 로그인 API(현재는 테스트용)
	e.POST("/api/sign-in", signIn())

	// 애플 로그인 API(현재는 테스트용)
	e.POST("/api/apple-sign-in", appleSignIn())

	// 목데이터로 테스트
	e.GET("/api/getlist", healthCheck())

	// 결심하기 리스트
	e.GET("/api/onboarding/goals", listGoals())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
