package handler

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func ApplyHandler(e *echo.Echo, db *sql.DB) {
	// 회원가입 API
	e.POST("/api/sign-up", signUp(db))

	// 로그인 API(현재는 테스트용)
	e.POST("/api/sign-in", signIn(db))

	// 애플 로그인 API(현재는 테스트용)
	e.POST("/api/apple-sign-in", appleSignIn(db))

	// 목데이터로 테스트
	e.GET("/api/getlist", healthCheck(db))

	// 결심하기 리스트
	e.GET("/api/onboarding/goals", listGoals(db))

	// 온보딩 - 시간 생성
	e.GET("/api/onboarding/status", getOnboardingStatus(db))

	// 온보딩 - 목표 생성
	e.POST("/api/onboarding/goals", createExerciseGoal(db))

	// 온보딩 - 날짜 생성
	e.POST("/api/onboarding/dates", createExerciseDate(db))

	// 온보딩 - 시간 수정
	e.PUT("/api/onboarding/dates", updateExercisedate(db))

	// 온보딩 - 시간 수정
	e.GET("/api/onboarding/dates", listExercisedates(db))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
