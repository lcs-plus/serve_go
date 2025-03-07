package models

type List struct {
	ID
	Title   string `json:"title" gorm:"not null;comment:标题"`
	Content string `json:"content" gorm:"not null;comment:内容"`
	Sort    int64  `json:"sort" gorm:"not null;default:0;comment:排序"`
	Status  int64  `json:"status" gorm:"not null;default:0;comment:状态"`
	Timestamps
	SoftDeletes
}

type GetListForm struct {
	Page    int64  `json:"page" form:"page" binding:"required,min=1"`
	Limit   int64  `json:"limit" form:"limit" binding:"required,min=1,max=100"`
	Sort    string `json:"sort" form:"sort"`
	Status  int64  `json:"status" form:"status"`
	Keyword string `json:"keyword" form:"keyword"`
}

type AddListForm struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (List) TableName() string {
	return "lists"
}
