package errorutil

import (
	"net/http"

	serror "github.com/DaZZler12/MyRestServer/pkg/serror"
	"github.com/gin-gonic/gin"
)

func HandleErrorResponse(c *gin.Context, err error) {
	switch err := err.(type) {
	case serror.AppError:
		c.JSON(err.StatusCode, gin.H{"message": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.Abort()
}
