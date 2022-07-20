package handler

import (
	"database/sql"
	"fmt"
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

type CompleteTodayExerciseRequest struct {
	Code int64 `json:"code"`
}

var kst, _ = time.LoadLocation("Asia/Seoul")

// @Summary 당일 운동시작, 종료 API
// @Description 유저의 운동의 시작과 종료 시간을 기록하는 API
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Param req body CompleteTodayExerciseRequest true "요청"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /exercises/today [post]
func completeTodayExercise(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		req := new(CompleteTodayExerciseRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
		}
		y, m, d := time.Now().In(kst).Date()

		nowDate := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		history, err := models.UserExerciseHistories(
			models.UserExerciseHistoryWhere.UserID.EQ(int64(chaUser.ID)),
			qm.OrderBy(fmt.Sprintf("%s DESC", models.UserExerciseHistoryColumns.ExerciseStartedAt)),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}

		if req.Code == 1 { // 운동 시작
			if history.ExerciseDate == nowDate {
				return c.JSON(http.StatusBadRequest, message{"이미 시작된 당일 운동이 있습니다."})
			}
			if err := (&models.UserExerciseHistory{
				UserID:            int64(chaUser.ID),
				ExerciseDate:      nowDate,
				ExerciseStartedAt: time.Now(),
			}).Insert(ctx, db, boil.Infer()); err != nil {
				return c.JSON(http.StatusInternalServerError, message{"Failed insert user exercise history"})
			}
		}

		if req.Code == 2 { // 운동 종료
			if history.ExerciseEndedAt.Valid {
				return c.JSON(http.StatusBadRequest, message{"이미 당일 운동이 종료되었습니다."})
			}
			history.ExerciseEndedAt = null.TimeFrom(time.Now())
			if _, err := history.Update(ctx, db, boil.Infer()); err != nil {
				return errors.Wrap(err, "update")
			}
		}

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}