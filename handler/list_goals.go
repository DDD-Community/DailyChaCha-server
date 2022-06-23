package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// @Summary auth 토큰으로 테스트해볼 API입니다.
// @Description access token을 확인하여 해당 토큰 유저의 이메일을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} email
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /onboarding/goals [get]
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
