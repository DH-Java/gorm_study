package dao

import "log"

type User struct {
	ID         int64
	UserName   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

func InsertUser(user *User) {
	if err := DB.Create(user).Error; err != nil {
		log.Println("insert user error", err)
	}
}

func SelectById(id int64, err error) User {
	var user User
	if err := DB.Where("id=?", id).First(&user); err != nil {
		log.Println("selectBy ID error", err)
	}
	return user
}

func SelectByAll() []User {
	var user []User
	if err := DB.Find(&user).Error; err != nil {
		log.Println("select All error", err)
	}
	return user
}

func UpdateUserById(id int64, userName string) {
	if err := DB.Model(&User{}).Where("id", id).Update("username", userName).Error; err != nil {
		log.Println("UpdateUserById error", err)
	}
}
func DeleteUserById(id int64) {
	if err := DB.Where("id", id).Delete(&User{}); err != nil {
		log.Println("UpdateUserById error", err)
	}
}
