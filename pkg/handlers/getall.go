package handlers

import (
	"net/http"
	"strconv"

	"github.com/DaZZler12/MyRestServer/pkg/errorutil"
	"github.com/DaZZler12/MyRestServer/pkg/models"

	// "github.com/DaZZler12/MyRestServer/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type GetAllItemsResponse struct {
	Items      []models.Item `json:"items"`
	TotalPages int64         `json:"totalPages"`
}

// GetAllItems godoc
// @Summary Get all items with pagination and optional query param as brand
// @Description Retrieve all items with pagination and also an optional query param added as brand
// @Tags Items
// @Accept json
// @Produce json
// @Param _start query integer false "Start index for pagination"
// @Param _end query integer false "End index for pagination"
// @Param brand query string false "Brand filter"
// @Success 200 {object} GetAllItemsResponse "success"
// @Failure 400 {object} gin.H "message: Invalid Start or End Value"
// @Failure 500 {object} gin.H "message": "Internal Server Error"
// @Router /api/items [get]
func (h *Handler) GetAllItems(c *gin.Context) {
	filters := bson.D{}            //defining the filter
	brandParam := c.Query("brand") //search by brand this are
	if brandParam != "" {
		filters = append(filters, bson.E{
			Key: "brand", Value: brandParam,
		})
	}
	start, err := strconv.Atoi(c.DefaultQuery("_start", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start value"})
		return
	}

	end, err := strconv.Atoi(c.DefaultQuery("_end", "4"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end value"})
		return
	}
	if start < 0 || end < 0 || start > end {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Start or End Value"})
		return
	}
	items, total, err := h.Service.GetAllItems(start, end, filters)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.Header("X-Total-Count", strconv.Itoa(int(total)))
	c.JSON(http.StatusOK, gin.H{"message": "success", "items": items, "totalPages": total})
}
