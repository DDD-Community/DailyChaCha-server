package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// 테스트용 API
func listGoals() echo.HandlerFunc {
	return func(c echo.Context) error {

		_, err := helper.ValidateJWT(c)
		if err != nil {
			return err
		}

		goalList := []string{
			"몸도 마음도 건강한 삶을 위해",
			"루틴한 삶을 위해",
			"멋진 몸매를 위해",
		}

		list := map[string][]string{
			"goals": goalList,
		}
		if err := c.JSON(http.StatusOK, list); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
