package controllers

import (
	"github.com/revel/revel"
//	. "github.com/leanote/leanote/app/lea"
)

// admin 首页

type AdminSetting struct {
	AdminBaseController
}

// email配置
func (c AdminSetting) Email() revel.Result {
	return nil
}