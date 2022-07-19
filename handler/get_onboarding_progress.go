package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type GetOnboardingProgressResponse struct {
	Progress string `json:"progress"`
}

// @Summary 온보딩 진행상황 API
// @Description 유저의 온보딩 진행상황을 반환합니다. 결심하기가 완료됐다면 'date', 날짜정하기를 완료했다면 'time', 시간정하기를 완료했다면 'alert'을 보냅니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} GetOnboardingProgressResponse
// @Failure 500 {object} message
// @Router /progress [get]
func getOnboardingProgress(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}
		if chaUser.IsOnboardingCompleted.Valid && chaUser.IsOnboardingCompleted.Bool {
			return c.JSON(http.StatusOK, GetOnboardingProgressResponse{
				Progress: "done",
			})
		}

		progress := "goal"

		g, err := models.ExerciseGoals(
			models.ExerciseGoalWhere.UserID.EQ(chaUser.ID),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}
		if g != nil {
			progress = "date"
		}

		dates, err := models.ExerciseDates(
			models.ExerciseDateWhere.UserID.EQ(chaUser.ID),
		).All(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		if len(dates) != 0 {
			progress = "time"
			for _, d := range dates {
				if d.ExerciseTime.Valid {
					progress = "alert"
				}
			}
		}

		if err := c.JSON(http.StatusOK, GetOnboardingProgressResponse{
			Progress: progress,
		}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
