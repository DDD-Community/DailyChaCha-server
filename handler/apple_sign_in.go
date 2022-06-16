package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

type Token struct {
	Token string `json:"token" `
}

func appleSignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := new(Token)
		if err := c.Bind(token); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "bad request",
			})
		}

		teamID := os.Getenv("APPLE_TEAM_ID")
		clientID := os.Getenv("APPLE_CLIENT_ID")
		keyID := os.Getenv("APPLE_KEY_ID")
		fmt.Println("token", token.Token)
		appleUser, err := helper.ValidateAuthorizationToken(token.Token, os.Getenv("SECRET_KEY"), clientID, teamID, keyID)
		if err != nil {
			return errors.Wrap(err, "helper.ValidateAuthorizationToken")
		}

		accessToken, err := helper.CreateJWT(appleUser.Email)
		if err != nil {
			return echo.ErrInternalServerError
		}

		db := db.Connect()
		user := new(models.User)
		result := db.Find(&user, "email=?", appleUser.Email)
		expiredAt := time.Now().AddDate(0, 3, 0)
		// 이미 이메일이 존재할 경우의 처리
		if result.RowsAffected != 0 {
			user.AccessToken = &accessToken
			user.ExpiredAt = &expiredAt
			db.Save(&user)
		} else {
			if err := db.Create(&models.User{
				Email:       appleUser.Email,
				AccessToken: &accessToken,
				ExpiredAt:   &expiredAt,
			}); err.Error != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"message": "Failed insert user",
				})
			}
		}

		if err := c.JSON(http.StatusOK, map[string]string{
			"message":      "Login Success",
			"access_token": accessToken,
		}); err != nil {
			return errors.Wrap(err, "signIn")
		}
		return nil
	}
}
