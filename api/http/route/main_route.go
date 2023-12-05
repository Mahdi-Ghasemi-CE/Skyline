package route

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("user")

	//	All Public APIs
	NewUserRouter(timeout, publicRouter)
}
