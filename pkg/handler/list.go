package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/saintvrai/Drom"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {

	//id, ok := c.Get(userCtx)
	//
	//if !ok {
	//	newErrorResponse(c, http.StatusInternalServerError, "user not found")
	//	return
	//}
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

func (h *Handler) getAllLists(c *gin.Context) {
	lists, err := h.services.CarList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{Data: lists})
}
func (h *Handler) getListByID(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
