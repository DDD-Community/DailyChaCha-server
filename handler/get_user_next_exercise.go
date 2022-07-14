package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type GetUserNextExerciseResponse struct {
	ContinuityExerciseDay int64   `json:"continuity_exercise_day"`
	ExerciseRemainingTime int64   `json:"exercise_remain_time"`
	ObjectImageURL        *string `json:"object_image_url"`
}

// @Summary 유저의 다음 운동정보 API
// @Description 홈에서 사용될 다음 운동정보들을 들을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} GetUserNextExerciseResponse
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /next-exercise [get]
func getUserNextExercise(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}
		abs := "https://dailychacha.s3.ap-northeast-2.amazonaws.com/object03.png"
		if err := c.JSON(http.StatusOK, GetUserNextExerciseResponse{
			ContinuityExerciseDay: 12,
			ExerciseRemainingTime: 36000,
			ObjectImageURL:        &abs,
		}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
