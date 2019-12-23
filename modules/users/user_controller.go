package users

import (
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
	condition := UserParameter{}
	if err := this.ParseForm(&condition); err != nil {
		this.WrapListResponse([0]UserDocument{}, 0, err, 1, 20)
		return
	}
	words, count, err := service.GetListWords(&condition)
	this.WrapListResponse(words, count, err, 1, 20)
}
