package handlers

import (
	"net/http"

	"github.com/DaZZler12/MyRestServer/pkg/errorutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// DeleteItemByName handles the DELETE /items/{name} endpoint.
// @Summary Delete item by Item_Name
// @Description Delete an item based on its Item_Name
// @Tags Items
// @Accept json
// @Produce json
// @Param nameParam path string true "Item_Name"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /items/{id} [delete]
func (h *Handler) DeleteItemByName(c *gin.Context) {
	// id := c.Param("id")
	nameParam := c.Query("name")
	// carid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
	// 	return
	// }
	filter := bson.M{"item_name": nameParam}
	_, err := h.Service.GetItemByName(primitive.M(filter)) //checking for wheather the item exists or not
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve item"})
		return
	}
	// ierr := h.Service.DeleteItemById(carid)
	ierr := h.Service.DeleteItemByName(nameParam)
	if ierr != nil {
		errorutil.HandleErrorResponse(c, ierr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
