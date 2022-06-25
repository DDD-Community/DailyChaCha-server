package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type exerciseDateForList struct {
	Goal          *string         `json:"goal"`
	ExerciseDates []*exerciseDate `json:"exercise_dates"`
}

// @Summary 온보딩 운동일정 가져오는 API
// @Description 유저의 온보딩 여부를 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} exerciseDateForList
// @Failure 500 {object} message
// @Router /onboarding/dates [get]
func listExercisedates() echo.HandlerFunc {
	return func(c echo.Context) error {

		chaUser, err := helper.ValidateJWT(c)
		if err != nil {
			return err
		}

		db := db.Connect()

		exerciseDates := exerciseDateForList{}

		exerciseGoal := new(models.ExerciseGoal)
		result := db.Find(&exerciseGoal, "user_id=?", chaUser.Id)
		if result.RowsAffected != 0 {
			exerciseDates.Goal = &exerciseGoal.ExerciseGoal
		}

		rows, err := db.Table("exercise_date").Where("user_id=?", chaUser.Id).Rows()
		if err != nil {
			return err
		}
		for rows.Next() {
			e := new(models.ExerciseDate)
			if err := rows.Scan(&e); err != nil {
				return c.JSON(http.StatusInternalServerError, message{"Failed scan date"})
			}
			exerciseDates.ExerciseDates = append(exerciseDates.ExerciseDates, &exerciseDate{
				ExerciseDate: e.ExerciseDate,
				ExerciseTime: e.ExerciseTime,
			})
		}

		if err := c.JSON(http.StatusOK, exerciseDates); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
