package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// 테스트용 API
func healthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Mock Data를 생성한다.
		list := map[string]string{
			"1": "고양이",
			"2": "사후자",
			"3": "호랑이",
		}
		if err := c.JSON(http.StatusOK, list); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
