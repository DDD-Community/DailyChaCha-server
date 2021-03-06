package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

var goalMap = map[string]int{
	"몸도 마음도 건강한 삶을 위해": 1,
	"루틴한 삶을 위해":        2,
	"멋진 몸매를 위해":        3,
}

type goal struct {
	Goal  string `json:"goal"`
	Index *int   `json:"index"`
}

// @Summary 결심하기 목록 API
// @Description 결심하기에서 사용할 목록들을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} []goal
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /goals [get]
func listGoals(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		_, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}
		firstIndex := 0
		secondIndex := 1
		thirdIndex := 2
		goalList := []goal{
			{"몸도 마음도 건강한 삶을 위해", &firstIndex},
			{"루틴한 삶을 위해", &secondIndex},
			{"멋진 몸매를 위해", &thirdIndex},
		}

		if err := c.JSON(http.StatusOK, goalList); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
