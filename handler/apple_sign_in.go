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

// @Summary 애플 로그인 API
// @Description Token을 받아 access token을 반환합니다.
// @Accept json
// @Produce json
// @Param token body Token true "애플로그인 token"
// @Success 200 {object} models.Auth
// @Failure 400 {object} message
// @Failure 500 {object} message
// @Router /apple-sign-in [post]
func appleSignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := new(Token)
		if err := c.Bind(token); err != nil {
			return c.JSON(http.StatusBadRequest, message{"bad request"})
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
				return c.JSON(http.StatusInternalServerError, message{"Failed insert user"})
			}
		}

		if err := c.JSON(http.StatusOK, models.Auth{
			AccessToken: accessToken,
			ExpiredAt:   expiredAt.Format("2006-01-02"),
		}); err != nil {
			return errors.Wrap(err, "signIn")
		}
		return nil
	}
}
