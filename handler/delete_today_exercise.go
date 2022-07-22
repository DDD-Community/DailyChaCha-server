package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// @Summary 당일 운동삭제 API
// @Description 유저의 운동데이터를 삭제하는 API
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /exercises/today [delete]
func deleteTodayExercise(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		now := time.Now()
		y, m, d := now.In(kst).Date()

		nowDate := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		_, err = models.UserExerciseHistories(
			models.UserExerciseHistoryWhere.UserID.EQ(int64(chaUser.ID)),
			models.UserExerciseHistoryWhere.ExerciseDate.EQ(nowDate),
		).DeleteAll(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		_, err = models.UserObjects(
			models.UserObjectWhere.UserID.EQ(int64(chaUser.ID)),
			models.UserObjectWhere.ExerciseDate.EQ(nowDate),
		).DeleteAll(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
