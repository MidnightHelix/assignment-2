package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MidnightHelix/assignment-2/internal/handler"
	"github.com/MidnightHelix/assignment-2/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/MidnightHelix/assignment-2/internal/service/mocks"
)

func TestGetOrders(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &mocks.OrderService{}

	mockSvc.On("GetOrders", mock.Anything).Return([]model.Order{}, nil)

	handler := handler.NewOrderHandler(mockSvc)

	router := gin.New()
	router.GET("/orders", handler.GetOrders)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var orders []model.Order
	err := json.Unmarshal(w.Body.Bytes(), &orders)
	assert.NoError(t, err)

	mockSvc.AssertCalled(t, "GetOrders", mock.Anything)
}

func TestCreateOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &mocks.OrderService{}

	mockOrder := model.Order{ID: 1, CustomerName: "Test"}

	mockSvc.On("CreateOrder", mock.Anything, mock.Anything).Return(mockOrder, nil)

	handler := handler.NewOrderHandler(mockSvc)

	router := gin.New()
	router.POST("/orders", handler.CreateOrder)

	requestBody, err := json.Marshal(mockOrder)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	mockSvc.AssertCalled(t, "CreateOrder", mock.Anything, mock.Anything)
}

func TestUpdateOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &mocks.OrderService{}

	mockID := uint64(1)
	mockOrder := model.Order{ID: mockID, CustomerName: "customer"}

	mockSvc.On("GetOrdersById", mock.Anything, mockID).Return(mockOrder, nil)

	mockSvc.On("UpdateOrder", mock.Anything, mock.Anything, mockID).Return(mockOrder, nil)

	handler := handler.NewOrderHandler(mockSvc)

	router := gin.New()
	router.PUT("/orders/:id", handler.UpdateOrder)

	requestBody, err := json.Marshal(mockOrder)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/orders/1", bytes.NewBuffer(requestBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockSvc.AssertCalled(t, "GetOrdersById", mock.Anything, mockID)
	mockSvc.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mockID)
}

func TestDeleteOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := &mocks.OrderService{}

	mockID := uint64(1)
	mockOrder := model.Order{ID: mockID, CustomerName: "customer"}

	mockSvc.On("GetOrdersById", mock.Anything, mockID).Return(mockOrder, nil)

	mockSvc.On("DeleteOrder", mock.Anything, mockID).Return(nil)

	handler := handler.NewOrderHandler(mockSvc)

	router := gin.New()
	router.DELETE("/orders/:id", handler.DeleteOrder)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/orders/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockSvc.AssertCalled(t, "GetOrdersById", mock.Anything, mockID)
	mockSvc.AssertCalled(t, "DeleteOrder", mock.Anything, mockID)
}
