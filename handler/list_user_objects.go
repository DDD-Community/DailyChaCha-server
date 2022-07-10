package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ListUserObjectsResponse struct {
	CharacterImageURL string               `json:"character_image_url"`
	Objects           []*models.Object     `json:"objects"`
	Backgrounds       []*models.Background `json:"backgrounds"`
	HasBrokenObject   bool                 `json:"has_broken_object"`
}

// @Summary 유저의 오브젝트 목록 API
// @Description 홈에서 사용될 배경, 오브젝트 목록들을 반환합니다.
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @param Authorization header string true "bearer {token}"
// @Success 200 {object} ListUserObjectsResponse
// @Failure 401 {object} message
// @Failure 400 {object} message
// @Router /objects [get]
func listUserObjects(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		_, err := helper.ValidateJWT(c, db)
		if err != nil {
			return err
		}

		resp := ListUserObjectsResponse{
			CharacterImageURL: "https://dailychacha.s3.ap-northeast-2.amazonaws.com/character.png",
			HasBrokenObject:   false,
		}

		objects, err := models.Objects().All(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}
		for _, o := range objects {
			resp.Objects = append(resp.Objects, o)
		}

		backgrounds, err := models.Backgrounds().All(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}
		for _, b := range backgrounds {
			resp.Backgrounds = append(resp.Backgrounds, b)
		}

		if err := c.JSON(http.StatusOK, resp); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
