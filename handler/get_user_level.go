package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type GetUserLevelResponse struct {
	Level int64   `json:"level"`
	Goal  *string `json:"goal"`
}

// @Summary 유저의 레벨정보를 가져오는 API
// @Description 홈에서 사용될 레벨, 목표를 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} GetUserLevelResponse
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /level [get]
func getUserLevel(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}
		resp := GetUserLevelResponse{
			Level: chaUser.Level,
		}
		g, err := models.ExerciseGoals(
			models.ExerciseGoalWhere.UserID.EQ(chaUser.ID),
		).One(ctx, db)
		if err != nil && errors.Cause(err) != sql.ErrNoRows {
			return echo.ErrInternalServerError
		}
		if g != nil {
			resp.Goal = &g.ExerciseGoal
		}

		if err := c.JSON(http.StatusOK, resp); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
