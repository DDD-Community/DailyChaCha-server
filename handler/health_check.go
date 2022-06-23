package handler

import (
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type user struct {
	Email string `json:"email"`
}

type message struct {
	Message string `json:"message"`
}

// @Summary Get test list
// @Description Get auth test api
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @name get-my-email
// @param Authorization header string true "Authorization"
// @Success 200 {object} user
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /api/getlist [get]
func healthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {

		chaUser, err := helper.ValidateJWT(c)
		if err != nil {
			return err
		}

		if err := c.JSON(http.StatusOK, user{
			Email: chaUser.Email,
		}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
