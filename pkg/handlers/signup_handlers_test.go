// package handlers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/DaZZler12/MyRestServer/mocks/pkg/service"
// 	"github.com/DaZZler12/MyRestServer/pkg/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestSignUp(t *testing.T) {
// 	t.Run("success", func(t *testing.T) {
// 		mockSignUpService := new(service.UserService)

// 		mockSignUpService.On("SignUp", mock.Anything).Return(nil)

// 		gin := gin.Default()

// 		rec := httptest.NewRecorder()

// 		pc := &Handler{
// 			Service: mockSignUpService,
// 		}

// 		gin.POST("/api/signup", pc.SignUp)

// 		body := models.User{
// 			Name:  "test",
// 			Email: "test@gmail.com",
// 			Age:   "1000",
// 		}

// 		var buf bytes.Buffer
// 		err := json.NewEncoder(&buf).Encode(body)
// 		assert.NoError(t, err)

// 		req := httptest.NewRequest(http.MethodPost, "/api/signup", &buf)
// 		gin.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusCreated, rec.Code)

// 		mockSignUpService.AssertExpectations(t)
// 	})

// 	t.Run("failure", func(t *testing.T) {
// 		mockSignUpService := new(service.UserService)

// 		mockSignUpService.On("SignUp", mock.Anything).Return(errors.New("some error"))

// 		gin := gin.Default()

// 		rec := httptest.NewRecorder()

// 		pc := &Handler{
// 			Service: mockSignUpService,
// 		}

// 		gin.POST("/api/signup", pc.SignUp)

// 		body := models.User{
// 			Name:  "test",
// 			Email: "test@gmail.com",
// 			Age:   "1000",
// 		}

// 		var buf bytes.Buffer
// 		err := json.NewEncoder(&buf).Encode(body)
// 		assert.NoError(t, err)

// 		req := httptest.NewRequest(http.MethodPost, "/api/signup", &buf)
// 		gin.ServeHTTP(rec, req)

// 		assert.Equal(t, http.StatusInternalServerError, rec.Code)

//			mockSignUpService.AssertExpectations(t)
//		})
//	}
package handlers
