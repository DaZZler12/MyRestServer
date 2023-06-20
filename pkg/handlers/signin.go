package handlers

import (
	"net/http"

	"github.com/DaZZler12/MyRestServer/pkg/errorutil"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/DaZZler12/MyRestServer/pkg/utils"
	"github.com/gin-gonic/gin"
)

// SignIn godoc
// @Summary      SignIn
// @Description  Sign in user with email and generate token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        user body models.UserInput true "User input data"
// @Success 200 {object} gin.H "message": "Log in success"
// @Failure 406 {object} gin.H "message": "Provide necessery details"
// @Failure 500 {object} gin.H "message": "Internal Server Error"
// @Router       /api/signin [post]
func (h *Handler) SignIn(c *gin.Context) {
	var userData models.UserInput
	if c.BindJSON(&userData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide necessery details"})
		c.Abort()
		return
	}

	user, err := h.Service.SignIn(userData.Email, userData.Password)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	token, _err := utils.GenerateToken(user.ID.String())
	if _err != nil {
		errorutil.HandleErrorResponse(c, _err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Log in success", "token": token})
}
