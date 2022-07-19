package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// @Summary 알림설정 완료 API
// @Description 유저의 온보딩 네번째 - 알림설정 완료하는 API입니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} message
// @Failure 500 {object} message
// @Router /alert [post]
func completeOnboardingAlert(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		chaUser.IsOnboardingCompleted = null.BoolFrom(true)

		if _, err := chaUser.Update(ctx, db, boil.Infer()); err != nil {
			return errors.Wrap(err, "update")
		}

		if err := c.JSON(http.StatusOK, message{"success"}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
