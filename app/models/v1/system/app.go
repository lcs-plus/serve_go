package system

import "0121_1/app/models"

type Apps struct {
	models.ID
	Name            string `json:"name" gorm:"column:name;not null"`
	Config          string `json:"config" gorm:"type:text;not null"`
	Users           string `json:"users" gorm:"column:users;column:users;not null"`
	Brand           string `json:"brand" gorm:"column:brand;not null"`
	Status          int8   `json:"status" gorm:"column:status;not null"`
	IsPlan          int8   `json:"is_plan" gorm:"column:is_plan;not null"`
	Uid             int    `json:"uid" gorm:"column:uid;not null"`
	Sort            int    `json:"sort" gorm:"column:sort;not null"`
	IsUploadCookies int8   `json:"is_upload_cookies" gorm:"column:is_upload_cookies;not null"`
	BaseCurrency    string `json:"base_currency" gorm:"column:base_currency;not null"`
	models.Timestamps
	models.SoftDeletes
}

func (m *Apps) TableName() string {
	return "apps"
}
