package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/OhMinsSup/pin-server/lib"

	"github.com/OhMinsSup/pin-server/database/models"
	shortid "github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func checkKeyword(k interface{}) (string, bool) {
	if k != nil {
		return "회원가입", false
	} else {
		return "로그인", true
	}
}

func createURI(u interface{}, code string) string {
	if u != nil {
		return "http://localhost:3000/email-register?code=" + code
	} else {
		return "http://localhost:3000/email-login?code=" + code
	}
}

func sendAuthEmail(c echo.Context) error {
	type SendEmail struct {
		Email string `json:"email" form:"email" validate:"required,email"`
	}

	db := c.Get("db").(*gorm.DB)
	u := new(SendEmail)
	g := shortid.Generator()

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request Error: "+err.Error())
	}

	// validate check
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validate Error:"+err.Error())
	}

	// email exists check
	var exists models.User
	db.Where("email = ?", u.Email).First(&exists)

	// emailAuth model created
	emailAuth := models.EmailAuth{
		Code:  g.Generate(),
		Email: u.Email,
	}

	db.NewRecord(emailAuth)
	db.Create(&emailAuth)

	keyword, ok := checkKeyword(exists)

	// templateData init
	templateData := struct {
		Keyword string
		URI     string
	}{
		Keyword: keyword,
		URI:     createURI(exists, emailAuth.Code),
	}

	// 이메일 템플릿 생성및 전송
	// Interpret and email html files
	m := lib.NewRequest([]string{u.Email}, keyword, "veloss<verification@gmail.com>")
	if err := m.ParseTemplate("statics/email.html", templateData); err == nil {
		ok := m.SendEmail()
		fmt.Println(ok)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"registered": ok,
	})
}

func code(c echo.Context) error {
	code := c.Param("code")

	db := c.Get("db").(*gorm.DB)

	//  발급된 code값 유효성 체크
	var emailAuth models.EmailAuth
	if err := db.Where("code = ?", code).First(&emailAuth).Error; err != nil {
		return c.JSON(http.StatusGone, echo.Map{
			"name": "Check_Email_Gone",
		})
	}

	// emailAuth가 생성된지(Time형) 24시간 전인지 조사
	target := emailAuth.CreatedAt.AddDate(0, 0, -1)
	expireDate := time.Duration(24) * time.Hour
	diff := (time.Since(target) >= expireDate)

	// 24시간이 지나면 유효시간을 초과
	if diff || emailAuth.Logged {
		return c.JSON(http.StatusGone, echo.Map{
			"name": "EXPIRED_CODE",
		})
	}

	// 유저 체크
	var user models.User
	if err := db.Where("email = ?", emailAuth.Email).First(&user).Error; err != nil {
		type registerTokenJwt struct {
			ID    string `json:id`
			Email string `json:"email"`
		}

		// 토큰으로 생성할 것 생성
		claims := &registerTokenJwt{
			emailAuth.ID,
			emailAuth.Email,
		}

		//  token is valid for 1hour
		exp := time.Now().Add(time.Hour * 24)

		// 토큰 생성
		registerToken, _ := lib.GenerateToken(claims, exp, "email-register")

		return c.JSON(http.StatusOK, echo.Map{
			"email": emailAuth.Email,
			"register_token": registerToken
		})
	}
	// 로그인 로직

	return c.String(http.StatusOK, "로그인 이다")
}

func locaRegister(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
