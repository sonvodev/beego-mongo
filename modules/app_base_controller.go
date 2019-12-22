package modules

import (
	"math"
	"strconv"

	"github.com/astaxie/beego"
)

type AppBaseController struct {
	beego.Controller
}

func (this *AppBaseController) WrapListResponse(data interface{}, count int64, err error, index int64, size int64) {

	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = data
		totalPage := math.Ceil(float64(count) / float64(size))
		this.Ctx.Output.Header("X-Pagination-Index", strconv.FormatInt(index, 2))
		this.Ctx.Output.Header("X-Pagination-TotalPage", strconv.FormatFloat(totalPage, 'f', 6, 64))
		this.Ctx.Output.Header("X-Pagination-Total", strconv.FormatInt(count, 2))
	}

	this.ServeJSON()
}
