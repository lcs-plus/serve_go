package services

import (
	"0121_1/app/models"
	"0121_1/global"
)

func GetList(listForm *models.GetListForm) *[]models.List {

	offset := int((listForm.Page - 1) * listForm.Limit)

	model := global.App.DB.Offset(offset).Limit(int(listForm.Limit))
	if listForm.Keyword != "" {
		model.Where("(title like ? Or content like ?)", "%"+listForm.Keyword+"%", "%"+listForm.Keyword+"%")
	}

	var list *[]models.List
	model.Find(&list)
	return list
}
