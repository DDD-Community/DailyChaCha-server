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
	e.GET("/api/user", getUser(db))

	// 결심하기 리스트
	e.GET("/api/goals", listGoals(db))

	// 온보딩 - 시간 생성
	e.GET("/api/status", getOnboardingStatus(db))

	// 온보딩 - 목표 생성
	e.POST("/api/goals", createExerciseGoal(db))

	// 온보딩 - 날짜 생성
	e.POST("/api/dates", createExerciseDate(db))

	// 온보딩 - 시간 수정
	e.PUT("/api/dates", updateExercisedate(db))

	// 온보딩 - 시간 수정
	e.GET("/api/dates", listExercisedates(db))

	// 온보딩 - 알림 설정
	e.POST("/api/alert", completeOnboardingAlert(db))

	// 온보딩 - 단계 확인
	e.GET("/api/progress", getOnboardingProgress(db))

	// 홈 - 오브젝트
	e.GET("/api/objects", listUserObjects(db))

	// 홈 - 레벨
	e.GET("/api/level", getUserLevel(db))

	// 홈 - 다음 운동정보
	e.GET("/api/next-exercise", getUserNextExercise(db))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
