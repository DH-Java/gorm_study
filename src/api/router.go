package api

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	r.POST("/insert", InsertUser)
	r.GET("/selectUserById", SelectUserById)
	r.GET("/selectUserAll", SelectUserAll)
	r.GET("/updateUserById", UpdateUserById)
	r.DELETE("/DeleteUserById", DeleteUserById)
}
