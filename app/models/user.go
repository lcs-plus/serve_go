package models

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;size:15;index;comment:用户手机号"`
	Password string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	Role     string `json:"role" gorm:"not null;"`
	Timestamps
	SoftDeletes
}

type UserForm struct {
	Name     string `json:"name" form:"name"`
	Mobile   string `json:"mobile" form:"mobile"`
	Password string `json:"password" form:"password"`
}

type UserLoginResult struct {
	User        User   `json:"user"`
	AccessToken string `json:"access_token"`
}

func (User) TableName() string {
	return "users"
}
