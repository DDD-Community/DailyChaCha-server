package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type exerciseDateForCreate struct {
	ExerciseDates []string `json:"exercise_dates"`
}

// @Summary 날짜정하기 생성 API
// @Description 유저의 온보딩 두번째 - 날짜를 생성하는 API입니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Param Weekday body exerciseDateForCreate true "요일"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /onboarding/dates [post]
func createExerciseDate() echo.HandlerFunc {
	return func(c echo.Context) error {
		chaUser, err := helper.ValidateJWT(c)
		if err != nil {
			return err
		}

		exerciseDate := new(exerciseDateForCreate)
		if err := c.Bind(exerciseDate); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}

		db := db.Connect()

		for _, e := range exerciseDate.ExerciseDates {
			if _, ok := helper.WeekdayToID[e]; !ok {
				return c.JSON(http.StatusBadRequest, message{"invaild date"})
			}
			if err := db.Create(&models.ExerciseDate{
				UserID:       chaUser.Id,
				ExerciseDate: e,
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
