package public_controller

import (
	custom_errors "Skyline/internal/custom-errors"
	user_models "Skyline/pkg/models/user-models"
	"Skyline/pkg/usecases/account_usecase"
	"Skyline/pkg/usecases/user_usecase"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
)

type AccountController struct {
	AccountUsecase account_usecase.AccountUsecaseInterface
	UserUsecase    user_usecase.UserUsecaseInterface
	ContextTimeout time.Duration
}

func (controller *AccountController) Login(ginContext *gin.Context) {
	var request user_models.LoginRequest
	_, cancel := context.WithTimeout(ginContext, controller.ContextTimeout)
	defer cancel()

	err := ginContext.ShouldBind(&request)
	if err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}

	validator := validator.New()
	if err := validator.Struct(request); err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ValidationErrorMessage(err))
		return
	}

	user, err := controller.AccountUsecase.Login(&request, ginContext.ClientIP(), ginContext.Request.UserAgent())
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}

	ginContext.JSON(http.StatusOK, user)
	return
}

func (controller *AccountController) SignUp(ginContext *gin.Context) {
	var request user_models.UserRequest
	_, cancel := context.WithTimeout(ginContext, controller.ContextTimeout)
	defer cancel()

	err := ginContext.ShouldBind(&request)
	if err != nil || request.Password != request.RePassword {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}

	validator := validator.New()
	if err := validator.Struct(request); err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ValidationErrorMessage(err))
		return
	}

	_, err = controller.UserUsecase.Create(&request)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}

	ginContext.JSON(http.StatusOK, nil)
	return
}

func (controller *AccountController) Verify(ginContext *gin.Context) {
	_, cancel := context.WithTimeout(ginContext, controller.ContextTimeout)
	defer cancel()

	email := ginContext.Param("email")

	verifyCodeParam := ginContext.Param("verifyCode")
	verifyCode, err := strconv.Atoi(verifyCodeParam)
	if err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}

	_, err = controller.UserUsecase.ActivateUser(email, verifyCode)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}

	ginContext.JSON(http.StatusOK, nil)
	return
}

func (controller *AccountController) ForgetPassword(ginContext *gin.Context) {
	_, cancel := context.WithTimeout(ginContext, controller.ContextTimeout)
	defer cancel()

	email := ginContext.Param("email")

	_, err := controller.AccountUsecase.ForgetPassword(email)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}

	ginContext.JSON(http.StatusOK, nil)
	return
}
