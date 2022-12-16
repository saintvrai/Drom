package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/saintvrai/Drom/internal/car"
	"github.com/saintvrai/Drom/internal/client"
	"net/http"
	"strconv"
)

// @Summary Create Car
// @Tags cars
// @Description create one car to Drom database
// @ID create-car
// @Accept  json
// @Produce  json
// @Param input body Drom.Car true "car info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createCar(c *gin.Context) {
	_, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}
	var input car.Car
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	_, err := h.services.Car.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Get All Cars
// @Tags cars
// @Description get all cars from database
// @ID get-all-cars
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllCarsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
type getAllCarsResponse struct {
	Data []car.Car `json:"data"`
}

func (h *Handler) getCarsList(c *gin.Context) {
	lists, err := h.services.Car.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllCarsResponse{Data: lists})
}

// @Summary Get Car By ID
// @Tags cars
// @Description get car by id from database
// @ID get-car-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Car ID"
// @Success 200 {object} Drom.Car
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getCarById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	list, err := h.services.Car.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update Car By ID
// @Tags cars
// @Description update car by id from database
// @ID update-car-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Car ID"
// @Param car body Drom.Car true "DromCar"
// @Success 200 {object} Drom.Car
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateCarById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input car.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Car.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Car By Id
// @Tags cars
// @Description Delete car by id from database
// @ID delete-car-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Car ID"
// @Success 200 {object} Drom.Car
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Car.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

type getAllCarsAndClientsResponse struct {
	CarData    []car.Car       `json:"carData"`
	ClientData []client.Client `json:"clientData"`
}

func (h *Handler) getAllCarsAndClients(c *gin.Context) {
	list, err := h.services.Car.GetAllCarsAndClients()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list[0])
}
