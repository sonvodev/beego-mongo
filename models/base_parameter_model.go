package models

type BaseParameterModel struct {
	Index   int64  `form:"index,int"`
	Size    int64  `form:"size,int"`
	Keyword string `form:"keyword"`
}
