package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) GetItemByName(c *gin.Context) {
	// id := c.Param("id")
	// carid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
	// 	return
	// }
	nameParam := c.Query("name")
	filter := bson.M{"item_name": nameParam}
	item, err := h.Service.GetItemByName(primitive.M(filter))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"item": item})
}
