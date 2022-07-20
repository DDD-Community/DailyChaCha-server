package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
)

type UserExercise struct {
	// 사용자 ID
	UserID int64 `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// 운동일
	ExerciseDate time.Time `boil:"exercise_date" json:"exercise_date" toml:"exercise_date" yaml:"exercise_date"`
	// 운동 시작 시간
	ExerciseStartedAt time.Time `boil:"exercise_started_at" json:"exercise_started_at" toml:"exercise_started_at" yaml:"exercise_started_at"`
	// 운동 종료 시간
	ExerciseEndedAt null.Time `boil:"exercise_ended_at" json:"exercise_ended_at,omitempty" toml:"exercise_ended_at" yaml:"exercise_ended_at,omitempty"`
}

type GetTodayExerciseResponse struct {
	Exercise            *UserExercise `json:"exercise"`
	IsExerciseCompleted bool          `json:"is_exercise_completed"`
}

// @Summary 유저의 당일 운동정보를 가져오는 API
// @Description 유저의 운동 시점을 가져옵니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} GetTodayExerciseResponse
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /exercises/today [get]
func getTodayExercise(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		y, m, d := time.Now().In(kst).Date()
		nowDate := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		history, err := models.UserExerciseHistories(
			models.UserExerciseHistoryWhere.UserID.EQ(int64(chaUser.ID)),
			models.UserExerciseHistoryWhere.ExerciseDate.EQ(nowDate),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}
		resp := GetTodayExerciseResponse{
			Exercise:            nil,
			IsExerciseCompleted: false,
		}
		if history != nil {
			resp = GetTodayExerciseResponse{
				Exercise: &UserExercise{
					UserID:            history.UserID,
					ExerciseDate:      history.ExerciseDate,
					ExerciseStartedAt: history.ExerciseStartedAt.In(kst),
				},
				IsExerciseCompleted: history.ExerciseEndedAt.Valid,
			}
			if history.ExerciseEndedAt.Valid {
				resp.Exercise.ExerciseEndedAt = null.TimeFrom(history.ExerciseEndedAt.Time.In(kst))
			}
		}

		if err := c.JSON(http.StatusOK, resp); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
