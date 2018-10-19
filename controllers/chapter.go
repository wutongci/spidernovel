package controllers

import (
	"spidernovel/models"
)

type ChapterController struct {
	BaseController
}

// @router /list [Get]
func (this *ChapterController) List() {
	list := models.GetAllList()
	this.ToJson(0,"success",list)
}
