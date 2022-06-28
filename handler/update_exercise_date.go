package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type exerciseDate struct {
	ExerciseDate string `json:"exercise_date"`
	ExerciseTime *int   `json:"exercise_time"`
}

type exerciseDateForUpdate struct {
	ExerciseDates []exerciseDate `json:"exercise_dates"`
}

// @Summary 시간정하기 API
// @Description 유저의 온보딩 세번째 - 시간을 생성하는 API입니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Param ExerciseDate body exerciseDateForUpdate true "습관 일정"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /onboarding/dates [put]
func updateExercisedate() echo.HandlerFunc {
	return func(c echo.Context) error {
		chaUser, err := helper.ValidateJWT(c)
		if err != nil {
			return err
		}

		exerciseDate := new(exerciseDateForUpdate)
		if err := c.Bind(exerciseDate); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}

		db := db.Connect()

		for _, e := range exerciseDate.ExerciseDates {
			date := new(models.ExerciseDate)
			result := db.Where("exercise_date = ? AND user_id = ?", e.ExerciseDate, chaUser.Id).Find(&date)
			// 이미 이메일이 존재할 경우의 처리
			if result.RowsAffected != 0 {
				date.ExerciseTime = e.ExerciseTime
				db.Save(&date)
			} else {
				return c.JSON(http.StatusBadRequest, message{"invaild request"})
			}
		}

		chaUser.IsOnboardingCompleted = true
		db.Save(&chaUser)

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
