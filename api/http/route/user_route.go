package route

import (
	"Skyline/api/http/controller"
	"Skyline/pkg/repository/role_repository"
	user_repository "Skyline/pkg/repository/user-repository"
	"Skyline/pkg/usecases/user_usecase"
	"github.com/gin-gonic/gin"
	"time"
)

func NewUserRouter(timeout time.Duration, router *gin.RouterGroup) {
	userRepository := user_repository.NewUserRepository()
	roleRepository := role_repository.NewRoleRepository()
	userUsecase := user_usecase.NewUserUsecase(userRepository, roleRepository)
	userController := &controller.UserController{
		UserUsecase: userUsecase,
	}
	router.POST("/Create", userController.Create)
	router.GET("/Get/:userId", userController.Get)
	router.POST("/Update", userController.Update)
	router.DELETE("/Delete/:userId", userController.Delete)
}
