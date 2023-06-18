package privateroutes

import (
	"github.com/DaZZler12/MyRestServer/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func Privateroutes(privateRoutes *gin.RouterGroup, h *handlers.Handler) {
	privateRoutes.GET("", h.GetAllItems) //api/items [and get method]
	privateRoutes.POST("", h.InsertItem)
	//privateRoutes.GET(":id", h.GetAllItems) // get item by id
	//privateRoutes.PUT(":id", h.UpdateItemByID) // get item by id
	privateRoutes.PUT("", h.UpdateItemByName)
	privateRoutes.DELETE("", h.DeleteItemByName)
}
