package main

import (
	"github.com/gin-gonic/gin"
	"gorm_study/src/router"
)

func main() {
	engine := gin.Default()

	router.InitRouter(engine)

	engine.Run(":8080")
}
