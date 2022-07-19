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

type CreateExerciseDateRequest struct {
	ExerciseDates []int `json:"exercise_dates"`
}

// @Summary 날짜정하기 생성 API
// @Description 유저의 온보딩 두번째 - 날짜를 생성하는 API입니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Param Weekday body CreateExerciseDateRequest true "요일"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /dates [post]
func createExerciseDate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		exerciseDate := new(CreateExerciseDateRequest)
		if err := c.Bind(exerciseDate); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}

		if _, err := models.ExerciseDates(
			models.ExerciseDateWhere.UserID.EQ(chaUser.ID),
		).DeleteAll(ctx, db); err != nil {
			return c.JSON(http.StatusInternalServerError, message{"delete error"})
		}

		for _, e := range exerciseDate.ExerciseDates {
			if _, ok := helper.WeekdayToString[helper.Weekday(e)]; !ok {
				return c.JSON(http.StatusBadRequest, message{"invaild date"})
			}
			if err := (&models.ExerciseDate{
				UserID:         chaUser.ID,
				ExerciseDate:   e,
				ExerciseDateEn: helper.WeekdayToString[helper.Weekday(e)],
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
