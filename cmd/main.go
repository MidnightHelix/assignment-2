package main

import (
	"github.com/MidnightHelix/assignment-2/internal/handler"
	"github.com/MidnightHelix/assignment-2/internal/infrastructure"
	"github.com/MidnightHelix/assignment-2/internal/repository"
	"github.com/MidnightHelix/assignment-2/internal/router"
	"github.com/MidnightHelix/assignment-2/internal/service"

	"github.com/gin-gonic/gin"

	_ "github.com/MidnightHelix/assignment-2/cmd/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO DTS USER API DOCUMENTATION
// @version		1.0
// @description	golang kominfo assignment 2
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/api/v1
// @schemes		http
func main() {
	g := gin.Default()

	v := g.Group("/api/v1")
	usersGroup := v.Group("/orders")

	gorm := infrastructure.NewGormPostgres()
	orderRepo := repository.NewOrderQuery(gorm)
	orderSvc := service.NewOrderService(orderRepo)
	orderHdl := handler.NewOrderHandler(orderSvc)
	orderRouter := router.NewOrderRouter(usersGroup, orderHdl)

	// mount
	orderRouter.Mount()
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}
