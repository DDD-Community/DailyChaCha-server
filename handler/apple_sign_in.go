package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

func appleSignIn() echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.FormValue("token")
		teamID := os.Getenv("APPLE_TEAM_ID")
		clientID := os.Getenv("APPLE_CLIENT_ID")
		keyID := os.Getenv("APPLE_KEY_ID")

		a, err := helper.ValidateAuthorizationToken(token, os.Getenv("SECRET_KEY"), clientID, teamID, keyID)
		if err != nil {
			return errors.Wrap(err, "helper.ValidateAuthorizationToken")
		}

		accessToken, err := helper.CreateJWT(a.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		cookie := new(http.Cookie)
		cookie.Name = "access-token"
		cookie.Value = accessToken
		cookie.HttpOnly = true
		cookie.Expires = time.Now().Add(time.Hour * 24)

		c.SetCookie(cookie)

		if err := c.JSON(http.StatusOK, map[string]string{
			"message": "Login Success",
		}); err != nil {
			return errors.Wrap(err, "signIn")
		}
		return nil
	}
}
