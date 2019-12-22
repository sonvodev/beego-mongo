package users

import (
	"wordbook/models"
	"wordbook/modules"
)

type UserController struct {
	modules.AppBaseController
}

var (
	service *UserService
)

func init() {
	service = &UserService{}
}

// @Title GetListUsers
// @Description get all users
// @Param   size        query    int   false        "The object content"
// @Param   index        query    int   false        "The object content"
// @Success 200 {object} UserDocument
// @router / [get]
func (this *UserController) Get() {

	filter := UserParameter{BaseParameterModel: models.BaseParameterModel{Index: 1, Size: 20}}
	words, count, err := service.GetListWords(filter)
	this.WrapListResponse(words, count, err, filter.Index, filter.Size)
}
