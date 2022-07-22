package handler

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CompleteTodayExerciseResponse struct {
	Object      *models.Object `json:"object"`
	CompletedAt time.Time      `json:"completed_at"`
}

var kst, _ = time.LoadLocation("Asia/Seoul")

// @Summary 당일 운동종료 API
// @Description 유저의 운동종료를 기록하는 API
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} CompleteTodayExerciseResponse
// @Failure 500 {object} message
// @Router /exercises/today/complete [post]
func completeTodayExercise(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		history, err := models.UserExerciseHistories(
			models.UserExerciseHistoryWhere.UserID.EQ(int64(chaUser.ID)),
			qm.OrderBy(fmt.Sprintf("%s DESC", models.UserExerciseHistoryColumns.ExerciseStartedAt)),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}

		if history == nil {
			return c.JSON(http.StatusBadRequest, message{"종료할 운동이 없습니다."})
		}
		if history.ExerciseEndedAt.Valid {
			return c.JSON(http.StatusBadRequest, message{"이미 최근 운동이 종료되었습니다."})
		}
		now := time.Now().Truncate(time.Second)
		history.ExerciseEndedAt = null.TimeFrom(now)
		if _, err := history.Update(ctx, db, boil.Infer()); err != nil {
			return errors.Wrap(err, "update")
		}
		objects, err := models.Objects(
			models.ObjectWhere.ObjectType.EQ("normal"),
		).All(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}

		userObject := objects[rand.Intn(len(objects))]
		if err := (&models.UserObject{
			UserID:       int64(chaUser.ID),
			ObjectID:     userObject.ID,
			ExerciseDate: history.ExerciseDate,
		}).Insert(ctx, db, boil.Infer()); err != nil {
			return c.JSON(http.StatusInternalServerError, message{"Failed insert user exercise history"})
		}
		return c.JSON(http.StatusOK, CompleteTodayExerciseResponse{
			Object:      userObject,
			CompletedAt: now.In(kst),
		})
	}
}
