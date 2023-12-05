package controller

import (
	custom_errors "Skyline/internal/custom-errors"
	user_models "Skyline/pkg/models/user-models"
	"Skyline/pkg/usecases/user_usecase"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"time"
)

type UserController struct {
	UserUsecase    user_usecase.UserUsecaseInterface
	contextTimeout time.Duration
}

func (controller *UserController) Create(ginContext *gin.Context) {
	var request user_models.UserRequest
	_, cancel := context.WithTimeout(ginContext, controller.contextTimeout)
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

	user, err := controller.UserUsecase.Create(&request)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}

	response := &user_models.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	ginContext.JSON(http.StatusOK, response)
	return
}

func (controller *UserController) Get(ginContext *gin.Context) {
	_, cancel := context.WithTimeout(ginContext, controller.contextTimeout)
	defer cancel()

	param := ginContext.Param("userId")
	userId, err := strconv.Atoi(param)

	if err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}

	user, err := controller.UserUsecase.Get(userId)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}
	if user.UserId == 0 {
		ginContext.JSON(http.StatusNotFound, custom_errors.ErrorResponse(custom_errors.CustomDataNotFoundError("user")))
		return
	}

	response := &user_models.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	ginContext.JSON(http.StatusOK, response)
	return
}

func (controller *UserController) Update(ginContext *gin.Context) {
	var request user_models.UpdateUserRequest
	_, cancel := context.WithTimeout(ginContext, controller.contextTimeout)
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

	user, err := controller.UserUsecase.Update(&request)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}

	response := &user_models.UserResponse{
		UserId:    user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	ginContext.JSON(http.StatusOK, response)
	return
}

func (controller *UserController) Delete(ginContext *gin.Context) {
	_, cancel := context.WithTimeout(ginContext, controller.contextTimeout)
	defer cancel()

	param := ginContext.Param("userId")
	userId, err := strconv.Atoi(param)

	if err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}

	response, err := controller.UserUsecase.Delete(userId)
	if err != nil {
		ginContext.JSON(http.StatusBadRequest, custom_errors.ErrorResponse(err))
		return
	}
	if response == false {
		ginContext.JSON(http.StatusInternalServerError, custom_errors.ErrorResponse(err))
		return
	}

	ginContext.JSON(http.StatusOK, nil)
	return
}
