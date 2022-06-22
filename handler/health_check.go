package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DDD-Community/DailyChaCha-server/db"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// 테스트용 API
func healthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {

		header := c.Request().Header
		authv := header.Get("Authorization")

		// Get bearer token
		if !strings.HasPrefix(strings.ToLower(authv), "bearer") {
			return errors.New("invalid bearer token")
		}

		values := strings.Split(authv, " ")
		if len(values) < 2 {
			return errors.New("no bearer token")
		}

		token := values[1]
		fmt.Println("success", token)
		user := new(models.User)
		db := db.Connect()
		result := db.Find(user, "access_token=?", token)
		// 존재하지않는 아이디일 경우
		if result.RowsAffected == 0 {
			return echo.ErrBadRequest
		}
		if user.ExpiredAt.Before(time.Now()) {
			return echo.ErrUnauthorized
		}
		fmt.Print("success2", user.Email, user.AccessToken)

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
