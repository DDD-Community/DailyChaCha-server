package handler

import (
	"database/sql"
	"net/http"

	"github.com/DDD-Community/DailyChaCha-server/helper"
	"github.com/DDD-Community/DailyChaCha-server/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type UserObject struct {
	ID int64 `json:"id" toml:"id" yaml:"id"`
	// 정렬 순서
	DisplayOrder int `json:"display_order" toml:"display_order" yaml:"display_order"`
	// 오브젝트 유형
	ObjectType string `json:"object_type" toml:"object_type" yaml:"object_type"`
	// 오브젝트 이름
	ObjectName string `json:"object_name" toml:"object_name" yaml:"object_name"`
	// 이미지 URL
	ImageURL string `json:"image_url" toml:"image_url" yaml:"image_url"`
}

type UserBackground struct { // ID
	ID int64 `json:"id" toml:"id" yaml:"id"`
	// 정렬 순서
	DisplayOrder int `json:"display_order" toml:"display_order" yaml:"display_order"`
	// 이미지 URL
	ImageURL string `json:"image_url" toml:"image_url" yaml:"image_url"`
}

type ListUserObjectsResponse struct {
	CharacterImageURL string            `json:"character_image_url"`
	Objects           []*UserObject     `json:"objects"`
	Backgrounds       []*UserBackground `json:"backgrounds"`
	HasBrokenObject   bool              `json:"has_broken_object"`
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
			CharacterImageURL: "https://dailychacha.s3.ap-northeast-2.amazonaws.com/character.gif",
			HasBrokenObject:   false,
		}

		objects, err := models.Objects().All(ctx, db)
		if err != nil {
			return echo.ErrInternalServerError
		}
		for i, o := range objects {
			resp.Objects = append(resp.Objects, &UserObject{
				ID:           o.ID,
				ImageURL:     o.ImageURL,
				ObjectType:   o.ObjectType,
				ObjectName:   o.ObjectName,
				DisplayOrder: i + 1,
			})
		}

		resp.Backgrounds = []*UserBackground{
			{
				ID:           2,
				DisplayOrder: 1,
				ImageURL:     "https://dailychacha.s3.ap-northeast-2.amazonaws.com/img_bg_gym_2.png",
			},
			{
				ID:           1,
				DisplayOrder: 2,
				ImageURL:     "https://dailychacha.s3.ap-northeast-2.amazonaws.com/img_bg_gym_1.png",
			},
			{
				ID:           2,
				DisplayOrder: 3,
				ImageURL:     "https://dailychacha.s3.ap-northeast-2.amazonaws.com/img_bg_gym_2.png",
			},
			{
				ID:           3,
				DisplayOrder: 4,
				ImageURL:     "https://dailychacha.s3.ap-northeast-2.amazonaws.com/img_bg_gym_3.png",
			},
		}

		if err := c.JSON(http.StatusOK, resp); err != nil {
			return errors.Wrap(err, "healthCheck")
		}
		return nil
	}
}
