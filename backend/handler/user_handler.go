package handler

import (
	"jobsforgobe/dto"
	"jobsforgobe/entities"
	"jobsforgobe/repository"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RepositoryAuth struct {
	Repo repository.AuthUserRepository
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthController(db *gorm.DB) *RepositoryAuth {
	repo := repository.AuthUserInterface(db)
	return &RepositoryAuth{Repo: repo}
}

func (sql *RepositoryAuth) SignUpHandler(c echo.Context) error {
	var form dto.AuthUserDto

	if err := c.Bind(&form); err != nil {
		return c.JSON(400, entities.ErrorModel{
			StatusCode: 400,
			ErrMessage: err.Error(),
			Message:    "Something Went Wrong",
		})
	}

	hash, err := HashPassword(form.Password)

	if err != nil {
		return c.JSON(500, entities.ErrorModel{
			StatusCode: 500,
			ErrMessage: err.Error(),
			Message:    "Something Went Wrong",
		})
	}

	formData := entities.AuthUser{
		Name:     form.Name,
		Username: form.Username,
		Email:    form.Email,
		Password: hash,
	}

	errPost := sql.Repo.SignUpAccount(formData)

	if errPost != nil {
		return c.JSON(500, entities.ErrorModel{
			StatusCode: 500,
			ErrMessage: errPost.Error(),
			Message:    "Something Went Wrong",
		})
	}

	return c.JSON(201, map[string]interface{}{
		"status":  201,
		"message": "Successfully Register Your Account",
	})
}

func (sql *RepositoryAuth) SignInHandler(c echo.Context) error {
	var form dto.AuthUserLoginDto

	if err := c.Bind(&form); err != nil {
		return c.JSON(400, entities.ErrorModel{
			StatusCode: 400,
			ErrMessage: err.Error(),
			Message:    "Something Went Wrong",
		})
	}

	formData := entities.AuthUser{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}

	data, errPost := sql.Repo.SignInAccout(formData)
	if errPost != nil {
		return c.JSON(500, entities.ErrorModel{
			StatusCode: 500,
			ErrMessage: errPost.Error(),
			Message:    "Something Went Wrong",
		})
	}
	if !CheckPasswordHash(formData.Password, data.Password) {
		return c.JSON(400, entities.ErrorModel{
			StatusCode: 400,
			ErrMessage: "Password is Incorrect",
			Message:    "Something Went Wrong",
		})
	} else {
		return c.JSON(200, map[string]interface{}{
			"status":  200,
			"message": "Successfully Login",
		})
	}

}
