package users

import (
	"wordbook/models"
)

type UserParameter struct {
	models.BaseParameterModel
	Name  string
	Email string
}
