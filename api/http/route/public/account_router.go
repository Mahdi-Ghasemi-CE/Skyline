package public

import (
	"Skyline/api/http/controller/public_controller"
	"Skyline/pkg/repository/role_repository"
	"Skyline/pkg/repository/session_repository"
	"Skyline/pkg/repository/user-repository"
	"Skyline/pkg/usecases/account_usecase"
	"Skyline/pkg/usecases/role_usecase"
	"Skyline/pkg/usecases/user_usecase"
	"github.com/gin-gonic/gin"
	"time"
)

func NewAccountRouter(timeout time.Duration, router *gin.RouterGroup) {
	userRepository := user_repository.NewUserRepository()
	sessionRepository := session_repository.NewSessionRepository()
	roleRepository := role_repository.NewRoleRepository()
	userUsecase := user_usecase.NewUserUsecase(userRepository, roleRepository)
	roleUsecase := role_usecase.NewRoleService(roleRepository)

	accountUsecase := account_usecase.NewAccountUsecase(userUsecase, sessionRepository, userRepository, roleUsecase)
	accountController := &public_controller.AccountController{
		AccountUsecase: accountUsecase,
		UserUsecase:    userUsecase,
	}
	//cityController := &public_controller.CityController{
	//	ContextTimeout: timeout,
	//}
	userGroup := router.Group("Account")
	userGroup.POST("/Login", accountController.Login)
	userGroup.POST("/SignUp", accountController.SignUp)
	userGroup.GET("/Verify/:email/:verifyCode", accountController.Verify)
	userGroup.GET("/ForgetPassword/:email", accountController.ForgetPassword)
	//userGroup.POST("/City", cityController.City)
	//userGroup.POST("/Country", cityController.Country)
	//userGroup.POST("/Continents", cityController.Continents)
}
