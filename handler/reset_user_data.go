package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// @Summary 유저 데이터 초기화 API
// @Description 유저의 모든 데이터를 초기화하는 API
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /reset [delete]
func resetUserData(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		_, err = models.UserExerciseHistories(
			models.UserExerciseHistoryWhere.UserID.EQ(int64(chaUser.ID)),
		).DeleteAll(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		_, err = models.UserObjects(
			models.UserObjectWhere.UserID.EQ(int64(chaUser.ID)),
		).DeleteAll(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		_, err = models.ExerciseGoals(
			models.ExerciseGoalWhere.UserID.EQ(chaUser.ID),
		).DeleteAll(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		_, err = models.ExerciseDates(
			models.ExerciseDateWhere.UserID.EQ(chaUser.ID),
		).DeleteAll(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		chaUser.IsOnboardingCompleted = null.BoolFrom(false)
		if _, err := chaUser.Update(ctx, db, boil.Infer()); err != nil {
			return errors.Wrap(err, "update")
		}

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
