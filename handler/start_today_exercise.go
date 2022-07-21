package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type StartTodayExerciseResponse struct {
	StartedAt time.Time `json:"started_at"`
}

// @Summary 당일 운동시작, 종료 API
// @Description 유저의 운동의 시작과 종료 시간을 기록하는 API
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} StartTodayExerciseResponse
// @Failure 500 {object} message
// @Router /exercises/today/start [post]
func startTodayExercise(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}
		now := time.Now()
		y, m, d := now.In(kst).Date()

		nowDate := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		history, err := models.UserExerciseHistories(
			models.UserExerciseHistoryWhere.UserID.EQ(int64(chaUser.ID)),
			qm.OrderBy(fmt.Sprintf("%s DESC", models.UserExerciseHistoryColumns.ExerciseStartedAt)),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}

		if history != nil && history.ExerciseDate == nowDate {
			return c.JSON(http.StatusBadRequest, message{"이미 시작된 당일 운동이 있습니다."})
		}
		if err := (&models.UserExerciseHistory{
			UserID:            int64(chaUser.ID),
			ExerciseDate:      nowDate,
			ExerciseStartedAt: now,
		}).Insert(ctx, db, boil.Infer()); err != nil {
			return c.JSON(http.StatusInternalServerError, message{"Failed insert user exercise history"})
		}

		if err := c.JSON(http.StatusOK, StartTodayExerciseResponse{
			StartedAt: now.In(kst),
		}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
