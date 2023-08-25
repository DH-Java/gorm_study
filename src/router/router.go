package router

import (
	"github.com/gin-gonic/gin"
	"gorm_study/src/api"
)

func InitRouter(router *gin.Engine) {
	api.RegisterRouter(router)
}
