package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type User struct {
	ID                  int64
	UserName            string              `gorm:"column:username"`
	Password            string              `gorm:"column:password"`
	DetailedInformation DetailedInformation `gorm:"embedded"`
	CreateTime          int64               `gorm:"column:createtime"`
	Admin               bool                `gorm:"-"`
	CreatedAt           time.Time
}
type DetailedInformation struct {
	Age     int    `json:"age gorm:column:age"`
	Sex     bool   `gorm:"column:sex"`
	Address string `gorm:"column:address"`
}

func (u User) TableName() string {
	return "users"
}

func UserTable(user User) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if user.Admin {
			return tx.Table("admin_users")
		}
		return tx.Table("users")
	}
}

//func InsertUser(user *User) {
//	user.Admin = false
//	 result := GetDB().Scopes(UserTable(*user)).Create(user)
//	 fmt.Printf("RowsAffecred:%d",result.RowsAffected)
//	if err := result.Error;err != nil {
//		//if err := DB.Table("users").Create(user).Error; err != nil {
//		log.Println("insert user error", err)
//	}
//}

// 密码md5加密
func InsertUser(user *User) {
	user.Admin = false
	//result := GetDB().Model(&User{}).Create(map[string]interface{}{
	//	"id":       user.ID,
	//	"username": user.UserName,
	//	"password": clause.Expr{SQL: "md5(?)", Vars: []interface{}{user.Password}},
	//})
	md5Password := clause.Expr{SQL: "md5(?)", Vars: []interface{}{user.Password}}
	//手写sql
	result := GetDB().Exec("insert into users (id,username,password) values (?,?,?)", user.ID, user.UserName, md5Password)
	fmt.Printf("RowsAffecred:%d\n", result.RowsAffected)
	if err := result.Error; err != nil {
		//if err := DB.Table("users").Create(user).Error; err != nil {
		log.Println("insert user error", err)
	}
}

func SelectById(id int64) User {
	var user User
	if err := GetDB().Where("id=?", id).First(&user); err != nil {
		log.Println("selectBy ID error", err)
	}
	return user
}

func SelectByAll() []User {
	var user []User
	if err := GetDB().Find(&user).Error; err != nil {
		log.Println("select All error", err)
	}
	return user
}

func UpdateUserById(id int64, userName string) {
	if err := GetDB().Model(&User{}).Where("id", id).Update("username", userName).Error; err != nil {
		log.Println("UpdateUserById error", err)
	}
}
func DeleteUserById(id int64) {
	if err := GetDB().Where("id", id).Delete(&User{}); err != nil {
		log.Println("UpdateUserById error", err)
	}
}
