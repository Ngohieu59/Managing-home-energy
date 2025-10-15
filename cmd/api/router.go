package api

import (
	"Managing-home-energy/cmd/api/controller"
	"Managing-home-energy/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func InitRouter(di *do.Injector) (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(middlewares.GenRequestId()) // Register the GenRequestId() middleware to run before any other handlers.
	r.Use(middlewares.GinZap())       // Register the GinZap() middleware to log HTTP requests for all routes.
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userController := controller.NewUserController(di)
	authController := controller.NewAuthController(di)
	ebillsController := controller.NewEBillsController(di)

	v1 := r.Group("/api/v1")

	//passwordlogin
	authGroup := v1.Group("/auth")
	authGroup.POST("/login", authController.PasswordLogin)

	// CRUD
	userGroup := v1.Group("/user")
	userGroup.POST("/create", userController.Create)
	userGroup.Use(middlewares.Auth(di))
	userGroup.PUT("/update/:id", userController.Update)
	userGroup.DELETE("/delete/:id", userController.Delete)
	userGroup.GET("/list", userController.List)

	// Electricity bill API
	eBillsGroup := v1.Group("/eBills")
	eBillsGroup.Use(middlewares.Auth(di))
	eBillsGroup.GET("/eMoney", ebillsController.EMoney)
	eBillsGroup.GET("/Report", ebillsController.ReportMonthly)
	return r, nil
}
