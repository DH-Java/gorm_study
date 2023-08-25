package api

import (
	"github.com/gin-gonic/gin"
	"gorm_study/src/dao"
	"net/http"
	"strconv"
	"time"
)

func InsertUser(context *gin.Context) {
	var user dao.User
	user.CreateTime = time.Now().UnixMilli()

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	dao.InsertUser(&user)
	context.JSON(http.StatusOK, gin.H{"data": gin.H{"data": user}})
}

func SelectUserById(context *gin.Context) {
	param := context.Query("id")
	user := dao.SelectById(strconv.ParseInt(param, 10, 64))
	context.JSON(http.StatusOK, gin.H{"data": gin.H{"data": user}})
}

func SelectUserAll(context *gin.Context) {
	userArrays := dao.SelectByAll()
	context.JSON(http.StatusOK, gin.H{"data": gin.H{"data": userArrays}})
}

func UpdateUserById(context *gin.Context) {
	paramId := context.Query("id")
	userName := context.Query("username")
	id, _ := strconv.ParseInt(paramId, 10, 64)
	dao.UpdateUserById(id, userName)
	context.JSON(http.StatusOK, gin.H{"data": gin.H{"success": true}})
}

func DeleteUserById(context *gin.Context) {
	paramId := context.Query("id")
	id, _ := strconv.ParseInt(paramId, 10, 64)
	dao.DeleteUserById(id)
	context.JSON(http.StatusOK, gin.H{"data": gin.H{"success": true}})
}
