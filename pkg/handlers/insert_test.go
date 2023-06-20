package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DaZZler12/MyRestServer/mocks"
	"github.com/DaZZler12/MyRestServer/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_InsertItem(t *testing.T) {
	testCases := []struct {
		desc          string
		service       *mocks.UserService
		requestBody   interface{}
		expStatusCode int
	}{
		{
			desc: "success",
			service: func() *mocks.UserService {
				mockService := new(mocks.UserService)
				mockService.On("InsertItem", mock.Anything).Return(nil)
				return mockService
			}(),
			requestBody: models.Item{
				Brand:     "BrandName",
				Model:     "ModelName",
				Item_Name: "ItemName",
				Year:      2023,
				Price:     9.99,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expStatusCode: http.StatusCreated,
		},
		{
			desc: "failure - service error",
			service: func() *mocks.UserService {
				mockService := new(mocks.UserService)
				mockService.On("InsertItem", mock.Anything).Return(errors.New("insert error"))
				return mockService
			}(),
			requestBody: models.Item{
				Brand:     "BrandName",
				Model:     "ModelName",
				Item_Name: "ItemName",
				Year:      2023,
				Price:     9.99,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			handler := &Handler{
				Service: tC.service,
			}
			server := gin.Default()
			server.POST("/api/items", handler.InsertItem)

			requestBody, err := json.Marshal(tC.requestBody)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/api/items", bytes.NewBuffer(requestBody))
			server.ServeHTTP(recorder, request)

			assert.Equal(t, tC.expStatusCode, recorder.Code)
			tC.service.AssertExpectations(t)
		})
	}
}
