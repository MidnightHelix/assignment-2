package handler

import (
	"net/http"
	"strconv"

	"github.com/MidnightHelix/assignment-2/internal/model"
	"github.com/MidnightHelix/assignment-2/internal/service"
	"github.com/MidnightHelix/assignment-2/pkg"
	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	GetOrders(ctx *gin.Context)
	GetOrdersByID(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderHandlerImpl struct {
	svc service.OrderService
}

func NewOrderHandler(svc service.OrderService) OrderHandler {
	return &orderHandlerImpl{
		svc: svc,
	}
}

// ShowOrders godoc
//
//	@Summary		Show orders list
//	@Description	Get all orders datat
//	@Tags			orders
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Order
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/orders [get]
func (u *orderHandlerImpl) GetOrders(ctx *gin.Context) {
	orders, err := u.svc.GetOrders(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (u *orderHandlerImpl) GetOrdersByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	order, err := u.svc.GetOrdersById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

//	 CreateOrder godoc
//
//		@Summary		Create an order
//		@Description	Create order with input payload
//		@Tags			orders
//		@Accept			json
//		@Produce		json
//		@Param order body Order true "Create Order"
//		@Success		201	{object}	[]model.Order
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
//		@Router			/orders [post]
func (u *orderHandlerImpl) CreateOrder(ctx *gin.Context) {

	order := model.Order{}
	if err := ctx.Bind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	order, err := u.svc.CreateOrder(ctx, order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

//	 UpdateOrder godoc
//
//		@Summary		Update an order
//		@Description	Update order with input payload
//		@Tags			orders
//		@Accept			json
//		@Produce		json
//		@Param order body Order true "Update Order"
//		@Param        id   path      int  true  "Order ID"
//		@Success		200	{object}	[]model.Order
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
//		@Router			/orders/{id} [put]
func (u *orderHandlerImpl) UpdateOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	order, err := u.svc.GetOrdersById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if order.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Order Not Found"})
		return
	}

	req := model.Order{ID: uint64(id)}
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	order, err = u.svc.UpdateOrder(ctx, req, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// DeleteOrder godoc
//
// @Summary		Delete an order
// @Description	Delete order with id param
// @Tags			orders
// @Accept			json
// @Produce		json
// @Param        id   path      int  true  "Order ID"
// @Success		200	{object}	[]model.Order
// @Failure		400	{object}	pkg.ErrorResponse
// @Failure		404	{object}	pkg.ErrorResponse
// @Failure		500	{object}	pkg.ErrorResponse
// @Router			/orders/{id} [delete]
func (u *orderHandlerImpl) DeleteOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	order, err := u.svc.GetOrdersById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if order.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "Order Not Found"})
		return
	}

	err = u.svc.DeleteOrder(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"message": "Order deleted",
	})
}
