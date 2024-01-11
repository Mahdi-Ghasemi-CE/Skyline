package public_controller

import (
	custom_errors "Skyline/internal/custom-errors"
	"Skyline/internal/utils"
	"Skyline/pkg/models/address_models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CityController struct {
	ContextTimeout time.Duration
}

func (controller *CityController) City(ginContext *gin.Context) {
	var _, cancel = context.WithTimeout(ginContext, controller.ContextTimeout)
	defer cancel()

	var request []address_models.City
	err := ginContext.ShouldBind(&request)
	if err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}

	fmt.Println("1")

	utils.DB.Create(request)

	fmt.Println("2")
	ginContext.JSON(http.StatusOK, request)
	return
}

func (controller *CityController) Country(ginContext *gin.Context) {
	var _, cancel = context.WithTimeout(ginContext, controller.ContextTimeout)
	defer cancel()

	var request []address_models.Country
	err := ginContext.ShouldBind(&request)
	if err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}
	fmt.Println("1")

	utils.DB.Create(request)

	fmt.Println("2")
	ginContext.JSON(http.StatusOK, request)
	return
}

func (controller *CityController) Continents(ginContext *gin.Context) {
	var _, cancel = context.WithTimeout(ginContext, controller.ContextTimeout)
	defer cancel()

	var request []address_models.Continent
	err := ginContext.ShouldBind(&request)
	if err != nil {
		ginContext.JSON(http.StatusNotAcceptable, custom_errors.ErrorResponse(err))
		return
	}
	fmt.Println("1")

	utils.DB.Create(request)

	fmt.Println("2")
	ginContext.JSON(http.StatusOK, request)
	return
}
