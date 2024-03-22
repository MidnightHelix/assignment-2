package router

import (
	"github.com/MidnightHelix/assignment-2/internal/handler"
	"github.com/gin-gonic/gin"
)

type OrderRouter interface {
	Mount()
}

type orderRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.OrderHandler
}

func NewOrderRouter(v *gin.RouterGroup, handler handler.OrderHandler) OrderRouter {
	return &orderRouterImpl{v: v, handler: handler}
}

func (o *orderRouterImpl) Mount() {
	// activity
	// /users/sign-up
	o.v.POST("", o.handler.CreateOrder)

	// /users
	o.v.GET("", o.handler.GetOrders)
	// /users/:id
	o.v.PUT("/:id", o.handler.UpdateOrder)

	o.v.DELETE("/:id", o.handler.DeleteOrder)
}
