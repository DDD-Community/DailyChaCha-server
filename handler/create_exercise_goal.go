package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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
func createExerciseGoal() echo.HandlerFunc {
	return func(c echo.Context) error {
		chaUser, err := helper.ValidateJWT(c)
		if err != nil {
			return err
		}

		goal := new(goal)
		if err := c.Bind(goal); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}

		db := db.Connect()
		exerciseGoal := new(models.ExerciseGoal)
		result := db.Find(&exerciseGoal, "user_id=?", chaUser.Id)
		if result.RowsAffected != 0 {
			return c.JSON(http.StatusBadRequest, message{"이미 결심이 생성되었습니다."})
		} else {
			if err := db.Create(&models.ExerciseGoal{
				UserID:       chaUser.Id,
				ExerciseGoal: goal.Goal,
			}); err.Error != nil {
				return c.JSON(http.StatusInternalServerError, message{"Failed insert user"})
			}
		}

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
