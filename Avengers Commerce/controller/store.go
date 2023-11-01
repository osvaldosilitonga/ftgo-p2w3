package controller

import (
	"net/http"
	api "ngc11/API"
	"ngc11/dto"
	"ngc11/repository"
	"ngc11/utils"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStoreController(db *gorm.DB) Store {
	return Store{
		DB: db,
	}
}

func (controller Store) GetAllStore(c echo.Context) error {
	stores, err := repository.GetAllStore(controller.DB)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	return c.JSON(http.StatusOK, stores)
}

func (controller Store) GetStoreDetailByID(c echo.Context) error {
	param := c.Param("id")
	storeID, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrInvalidParamID)
	}

	// get store data by id
	store, err := repository.GetStoreById(storeID, controller.DB)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrBadRequest)
	}

	// Get weather from third party API by passing lat and lot data
	weather, err := api.GetWeather(store.Lat, store.Lon)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ErrInternalServer)
	}

	// make response
	storeResponse := dto.StoreResponse{
		ID:     int(store.ID),
		Name:   store.Name,
		Rating: store.Rating,
	}

	storeDetailResponse := dto.StoreDetailResponse{
		Code:    http.StatusOK,
		Store:   storeResponse,
		Weather: *weather,
	}

	return c.JSON(http.StatusOK, storeDetailResponse)

}
