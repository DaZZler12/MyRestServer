package handlers

import (
	"net/http"

	"github.com/DaZZler12/MyRestServer/pkg/errorutil"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/gin-gonic/gin"
)

// SignUp handles the POST /signup endpoint.
// @Summary Register a new user account
// @Description Register a new user account with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param userData body models.User true "User details"
// @Success 201 {object} gin.H
// @Router /api/signup [post]
func (h *Handler) SignUp(c *gin.Context) {
	var userData models.User
	if c.BindJSON(&userData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide necessary details"})
		c.Abort()
		return
	}

	err := h.Service.SignUp(userData)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})

}
