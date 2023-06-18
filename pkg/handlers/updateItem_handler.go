package handlers

import (
	"net/http"

	"github.com/DaZZler12/MyRestServer/pkg/errorutil"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// UpdateItemByName handles the PUT /cars/{id} endpoint.
// @Summary Update a item by Name
// @Description Update an existing item with the provided details
// @Tags Items
// @Accept json
// @Produce json
// @Param name path string true "Item item_name"
// @Param itemData body models.Item true "Car details"
// @Success 200 {string} string "Updated successfully"
// @Router /api/items/{name} [put]
func (h *Handler) UpdateItemByName(c *gin.Context) {
	// id := c.Param("id")
	// carid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
	// 	return
	// }
	nameParam := c.Query("name")
	filter := bson.M{"item_name": nameParam}
	_, err := h.Service.GetItemByName(primitive.M(filter)) //checking for wheather the item exists or not
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve item"})
		return
	}
	var itemData models.Item
	if c.BindJSON(&itemData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}
	// cerr := h.Service.UpdateItemByID(itemData, carid)
	// if cerr != nil {
	// 	errorutil.HandleErrorResponse(c, err)
	// 	return
	// }
	cerr := h.Service.UpdateItemByName(itemData, nameParam)
	if cerr != nil {
		errorutil.HandleErrorResponse(c, cerr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}
