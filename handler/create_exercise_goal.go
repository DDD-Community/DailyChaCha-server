package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// @Summary 결심하기 생성 API
// @Description 유저의 온보딩 첫번째 - 결심을 생성하는 API입니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Param Goal body goal true "결심"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /onboarding/goals [post]
func createExerciseGoal(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		goal := new(goal)
		if err := c.Bind(goal); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}

		g, err := models.ExerciseGoals(
			models.ExerciseGoalWhere.UserID.EQ(chaUser.ID),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}
		if g != nil {
			g.ExerciseGoal = goal.Goal
			if _, err := g.Update(ctx, db, boil.Infer()); err != nil {
				return echo.ErrInternalServerError
			}
			return c.JSON(http.StatusOK, message{"success"})
		} else {
			if err := (&models.ExerciseGoal{
				UserID:       chaUser.ID,
				ExerciseGoal: goal.Goal,
			}).Insert(ctx, db, boil.Infer()); err != nil {
				return c.JSON(http.StatusInternalServerError, message{"Failed insert user"})
			}
		}

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
