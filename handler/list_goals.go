package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type goal struct {
	Goal string `json:"goal"`
}

// @Summary 결심하기 목록 API
// @Description 결심하기에서 사용할 목록들을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} "goals":[]goal
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /onboarding/goals [get]
func listGoals(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		_, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		goalList := map[string][]goal{
			"goals": {
				{"몸도 마음도 건강한 삶을 위해"},
				{"루틴한 삶을 위해"},
				{"멋진 몸매를 위해"},
			},
		}

		if err := c.JSON(http.StatusOK, goalList); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
