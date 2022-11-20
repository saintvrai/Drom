package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	userCtx :=
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user not found")
		return
	}
	var input }
func (h *Handler) getAllLists(c *gin.Context) {

}
func (h *Handler) getListByID(c *gin.Context) {

}
func (h *Handler) updateList(c *gin.Context) {

}
func (h *Handler) deleteList(c *gin.Context) {

}
