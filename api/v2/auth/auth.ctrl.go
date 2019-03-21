package auth

import (
	"fmt"
	"net/http"

	"github.com/OhMinsSup/pin-server/lib"

	"github.com/OhMinsSup/pin-server/database/models"
	shortid "github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func checkKeyword(k interface{}) string {
	if k != nil {
		return "register"
	} else {
		return "login"
	}
}

func createUrl(u interface{}, code string) string {
	if u != nil {
		return "http://localhost:3000/email-register/?code=" + code
	} else {
		return "http://localhost:3000/email-login/?code=" + code
	}
}

func sendAuthEmail(c echo.Context) error {
	type SendEmail struct {
		Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	}

	db := c.Get("db").(*gorm.DB)
	u := new(SendEmail)
	g := shortid.Generator()

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request Error: "+err.Error())
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validate Error:"+err.Error())
	}

	var exists models.User
	db.Debug().Where("email = ?", u.Email).First(&exists)

	emailAuth := models.EmailAuth{
		Code:  g.Generate(),
		Email: u.Email,
	}

	db.NewRecord(emailAuth)
	db.Create(&emailAuth)

	keyword := checkKeyword(exists)

	templateData := struct {
		Keyword string
		URI     string
	}{
		Keyword: keyword,
		URI:     createUrl(exists, emailAuth.Code),
	}

	m := lib.NewRequest([]string{u.Email}, "veloss"+keyword, "Hello world")
	if err := m.ParseTemplate("/statics/email.html", templateData); err == nil {
		ok := m.SendEmail()
		fmt.Println(ok)
	}

	return c.JSON(http.StatusOK, u)
}

func locaRegister(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
