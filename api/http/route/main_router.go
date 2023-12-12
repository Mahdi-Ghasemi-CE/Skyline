package route

import (
	"Skyline/api/http/route/admin"
	"Skyline/api/http/route/public"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(timeout time.Duration, gin *gin.Engine) {
	version1 := gin.Group("v1")

	//	All Admin APIs
	adminRouter := version1.Group("Admin")
	admin.NewUserAdminRouter(timeout, adminRouter)

	//	All public_controller APIs
	publicRouter := version1.Group("")
	public.NewAccountRouter(timeout, publicRouter)
}
