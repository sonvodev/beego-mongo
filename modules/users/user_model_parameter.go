package users

import (
	"wordbook/models"
)

type UserParameter struct {
	models.BaseParameterModel
	Name  string `form:"name"`
	Email string `form:"email"`
}
