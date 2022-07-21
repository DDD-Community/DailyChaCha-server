package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type GetUserNextExerciseResponse struct {
	ContinuityExerciseDay int64   `json:"continuity_exercise_day"`
	ExerciseRemainingTime int64   `json:"exercise_remain_time"`
	ObjectImageURL        *string `json:"object_image_url"`
}

// @Summary 유저의 다음 운동정보 API
// @Description 홈에서 사용될 다음 운동정보들을 들을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} GetUserNextExerciseResponse
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /next-exercise [get]
func getUserNextExercise(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}
		now := time.Now()
		y, m, d := now.In(kst).Date()

		nowDate := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		o, err := models.UserObjects(
			models.UserObjectWhere.UserID.EQ(int64(chaUser.ID)),
			models.UserObjectWhere.ExerciseDate.EQ(nowDate),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}

		if o != nil {
			object, err := models.Objects(
				models.ObjectWhere.ID.EQ(o.ObjectID),
			).One(ctx, db)
			if err != nil && errors.Cause(err) != sql.ErrNoRows {
				return echo.ErrInternalServerError
			}

			return c.JSON(http.StatusOK, GetUserNextExerciseResponse{
				ContinuityExerciseDay: 12,
				ExerciseRemainingTime: 36000,
				ObjectImageURL:        &object.ImageURL,
			})
		}
		return c.JSON(http.StatusOK, GetUserNextExerciseResponse{
			ContinuityExerciseDay: 12,
			ExerciseRemainingTime: 36000,
			ObjectImageURL:        nil,
		})
	}
}
