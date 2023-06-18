package handlers

import (
	"net/http"
	"strconv"

	"github.com/DaZZler12/MyRestServer/pkg/errorutil"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type GetAllItemsResponse struct {
	Items      []models.Item `json:"items"`
	TotalPages int64         `json:"totalPages"`
}

// GetAllItems retrieves all items with pagination.
// @Summary Get all items with pagination
// @Description Retrieve all items with pagination
// @Tags Items
// @Accept json
// @Produce json
// @Param _start query integer false "Start index for pagination"
// @Param _end query integer false "End index for pagination"
// @Success 200 {object} GetAllItemsResponse "success"
// @Router /api/items [get]
func (h *Handler) GetAllItems(c *gin.Context) {
	filters := bson.D{}            //defining the filter
	brandParam := c.Query("brand") //search by brand this are
	// if idParam != "" {
	// 	// Retrieve a specific car by ID
	// 	id, err := primitive.ObjectIDFromHex(idParam)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
	// 		return
	// 	}

	// 	car, err := h.Service.GetCarByID(id)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve car"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"car": car})
	// 	return
	// }
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

	end, err := strconv.Atoi(c.DefaultQuery("_end", "3"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end value"})
		return
	}
	pagination := utils.Pagination{
		PageNumber: (start / end) + 1,
		PageSize:   end - start,
	}
	// fmt.Println("filter for  handler: ", filters)
	cars, total, err := h.Service.GetAllItems(pagination, filters)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "items": cars, "totalPages": total})

}
