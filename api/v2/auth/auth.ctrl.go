package auth

import (
	"net/http"
	"time"

	"github.com/OhMinsSup/pin-server/database/models"
	"github.com/OhMinsSup/pin-server/lib"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// RequestLocalRegister localRegister 요청 Body 데이터값
type RequestLocalRegister struct {
	Username    string `json:"username" validate:"required,max=16,min=1"`
	DisplayName string `json:"display_name" validate:"required,max=45,min=1"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required, min=6"`
}

func locaRegister(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	body := new(RequestLocalRegister)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request Error: "+err.Error())
	}

	if err := c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validate Error:"+err.Error())
	}

	var exists models.User
	if err := db.Where("username = ?", body.Username).Or("email = ?", body.Email).First(&exists).Error; err == nil {
		return c.JSON(http.StatusConflict, echo.Map{
			"message": "유저명또는 이메일이 이미 존재합니다.",
		})
	}

	hash, hashErr := lib.Hash(body.Password)
	if hashErr != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	user := models.User{
		Username:    body.Username,
		Email:       body.Email,
		DisplayName: body.DisplayName,
		Password:    hash,
	}

	db.NewRecord(user)
	db.Create(&user)

	serialized := user.Serialize()
	token, _ := lib.GenerateToken(serialized)

	cookie := new(http.Cookie)
	cookie.Name = "access_token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24 * 7)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"user":  serialized,
		"token": token,
	})
}
