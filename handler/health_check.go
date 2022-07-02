package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type GetUserResponse struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

type message struct {
	Message string `json:"message"`
}

// @Summary 유저정보를 가져오는 API입니다.
// @Description access token을 확인하여 해당 토큰 유저의 이메일과 user id를 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} GetUserResponse
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /user [get]
func getUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		chaUser, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		if err := c.JSON(http.StatusOK, GetUserResponse{
			UserID: chaUser.ID,
			Email:  chaUser.Email,
		}); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
