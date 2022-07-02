package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ListExercisedatesResponse struct {
	Goal          *string         `json:"goal"`
	ExerciseDates []*exerciseDate `json:"exercise_dates"`
}

// @Summary 온보딩 운동일정 가져오는 API
// @Description 유저의 온보딩 여부를 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} ListExercisedatesResponse
// @Failure 500 {object} message
// @Router /onboarding/dates [get]
func listExercisedates(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		resp := ListExercisedatesResponse{}

		g, err := models.ExerciseGoals(
			models.ExerciseGoalWhere.UserID.EQ(chaUser.ID),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}
		if g != nil {
			resp.Goal = &g.ExerciseGoal
		}

		dates, err := models.ExerciseDates(
			models.ExerciseDateWhere.UserID.EQ(chaUser.ID),
		).All(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		for _, d := range dates {
			date := &exerciseDate{
				ExerciseDate: d.ExerciseDate,
			}
			if d.ExerciseTime.Valid {
				date.ExerciseTime = d.ExerciseTime.Ptr()
			}
			resp.ExerciseDates = append(resp.ExerciseDates, date)
		}

		if err := c.JSON(http.StatusOK, resp); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
