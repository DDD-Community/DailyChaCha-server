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

type exerciseDate struct {
	ExerciseDate int  `json:"exercise_date"`
	ExerciseTime *int `json:"exercise_time"`
}

type UpdateExerciseDateRequest struct {
	ExerciseDates []exerciseDate `json:"exercise_dates"`
}

// @Summary 시간정하기 API
// @Description 유저의 온보딩 세번째 - 시간을 생성하는 API입니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Param ExerciseDate body UpdateExerciseDateRequest true "습관 일정"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /dates [put]
func updateExercisedate(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		exerciseDate := new(UpdateExerciseDateRequest)
		if err := c.Bind(exerciseDate); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}

		for _, e := range exerciseDate.ExerciseDates {
			date, err := models.ExerciseDates(
				models.ExerciseDateWhere.UserID.EQ(chaUser.ID),
				models.ExerciseDateWhere.ExerciseDate.EQ(e.ExerciseDate),
			).One(ctx, db)
			if errors.Cause(err) == sql.ErrNoRows {
				return c.JSON(http.StatusBadRequest, message{"invaild request"})
			}
			if err != nil && errors.Cause(err) != sql.ErrNoRows {
				return echo.ErrInternalServerError
			}
			date.ExerciseTime = null.IntFrom(*e.ExerciseTime)
			if _, err := date.Update(ctx, db, boil.Infer()); err != nil {
				return c.JSON(http.StatusInternalServerError, message{"Failed insert user"})
			}
		}

		dates, err := models.ExerciseDates(
			models.ExerciseDateWhere.UserID.EQ(chaUser.ID),
		).All(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		for _, d := range dates {
			if !d.ExerciseTime.Valid {
				return c.JSON(http.StatusBadRequest, message{"생성한 모든 요일에 대해 시간을 채워야 합니다."})
			}
		}

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
