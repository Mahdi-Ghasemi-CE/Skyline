package admin

import (
	"Skyline/api/http/controller/admin_controller"
	"Skyline/pkg/repository/role_repository"
	"Skyline/pkg/repository/user-repository"
	"Skyline/pkg/usecases/user_usecase"
	"github.com/gin-gonic/gin"
	"time"
)

func NewUserAdminRouter(timeout time.Duration, router *gin.RouterGroup) {
	userRepository := user_repository.NewUserRepository()
	roleRepository := role_repository.NewRoleRepository()
	userUsecase := user_usecase.NewUserUsecase(userRepository, roleRepository)
	userController := &admin_controller.UserController{
		UserUsecase:    userUsecase,
		ContextTimeout: timeout,
	}
	userGroup := router.Group("user")
	userGroup.POST("/Create", userController.Create)
	userGroup.GET("/Get/:userId", userController.Get)
	userGroup.POST("/Update", userController.Update)
	userGroup.DELETE("/Delete/:userId", userController.Delete)
}
