package handlers

import (
	"net/http"

	"github.com/DaZZler12/MyRestServer/pkg/errorutil"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/gin-gonic/gin"
)

// InsertItem handles the POST /cars endpoint.

// InsertItem godoc
// @Summary Insert a new item
// @Description Insert a new item with the provided details
// @Tags Items
// @Accept json
// @Produce json
// @Param itemData body models.Item true "Car details"
// @Success 201 {object} gin.H "success"
// @Failure      400  {object}  gin.H    "message: Bad Request"
// @Router /api/items/add [post]
func (h *Handler) InsertItem(c *gin.Context) {
	var itemData models.Item
	if c.BindJSON(&itemData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}
	err := h.Service.InsertItem(itemData)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "New Item Created"})

}
