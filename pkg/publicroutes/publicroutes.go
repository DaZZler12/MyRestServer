package publicroutes

import (
	"github.com/DaZZler12/MyRestServer/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(public *gin.RouterGroup, h *handlers.Handler) {
	public.POST("/signin", h.SignIn)
	public.POST("/signup", h.SignUp)
}
