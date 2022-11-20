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
	c.JSON(http.StatusOK, map[string]interface{}{"id": err.Error()})
}
func (h *Handler) getAllLists(c *gin.Context) {

}
func (h *Handler) getListByID(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
