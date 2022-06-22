package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// 테스트용 API
func healthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {

		user, err := helper.ValidateJWT(c)
		if err != nil {
			return nil
		}

		// Mock Data를 생성한다.
		list := map[string]string{
			"로그인한 유저": user.Email,
		}
		if err := c.JSON(http.StatusOK, list); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
