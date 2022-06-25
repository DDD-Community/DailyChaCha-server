package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type onboardingStatus struct {
	IsOnboardingCompleted bool `json:"is_onboarding_completed"`
}

// @Summary 온보딩 상태 API
// @Description 유저의 온보딩 여부를 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} onboardingStatus
// @Failure 500 {object} message
// @Router /onboarding/status [get]
func getOnboardingStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		chaUser, err := helper.ValidateJWT(c)
		if err != nil {
			return err
		}
		if err := c.JSON(http.StatusOK, onboardingStatus{
			IsOnboardingCompleted: chaUser.IsOnboardingCompleted,
		}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
