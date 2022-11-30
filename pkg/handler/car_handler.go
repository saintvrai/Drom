package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/saintvrai/Drom"
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

	var input Drom.Car
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	_, err := h.services.CarList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"create": input})
}

type getAllListsResponse struct {
	Data []Drom.Car `json:"data"`
}

// @Summary Get All Cars
// @Tags cars
// @Description get all cars from database
// @ID get-all-cars
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getCarsList(c *gin.Context) {
	lists, err := h.services.CarList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}

// @Summary Get Car By Id
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
	list, err := h.services.CarList.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update Car By Id
// @Tags cars
// @Description update car by id from database
// @ID update-car-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Car ID"
// @Param name path string true "NameCar"
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
	var input Drom.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.CarList.Update(id, input); err != nil {
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
	err = h.services.CarList.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
