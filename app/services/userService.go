package services

import (
	"0121_1/app/models"
	"0121_1/global"
	"0121_1/utils"
)

// Register /**
func Register(UserForm *models.UserForm) {
	var user models.User

	user.Name = UserForm.Name
	user.Password = utils.Md5(utils.Base64Encode(UserForm.Password))
	user.Mobile = UserForm.Mobile

	//parsedTime, err := time.Parse("2006-01-02T15:04:05", time.Now().Format("2006-01-02T15:04:05"))
	//if err != nil {
	//	panic(err)
	//}
	//user.CreatedAt = parsedTime
	global.App.DB.Table("users").Create(&user)

	global.App.Log.Debug(user.Name)
}

func Login(userForm *models.UserForm) *models.User {

	var user *models.User
	findUser := global.App.DB.Where("name=?", userForm.Name).Find(&user)

	userForm.Password = utils.Md5(utils.Base64Encode(userForm.Password))
	if findUser.RowsAffected != 0 && userForm.Password == user.Password {
		return user
	} else {
		return nil
	}
}
