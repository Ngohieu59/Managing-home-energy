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
	eBillsController := controller.NewEBillsController(di)
	staffController := controller.NewStaffController(di)

	v1 := r.Group("/api/v1")

	//password login
	authGroup := v1.Group("/auth")
	authGroup.POST("/login", authController.PasswordLogin)
	authGroup.POST("/loginStaff", authController.StaffPasswordLogin)

	// CRUD
	userGroup := v1.Group("/user")
	userGroup.POST("/create", userController.Create)
	userGroup.Use(middlewares.Auth(di))
	userGroup.PUT("/update/:id", userController.Update)
	userGroup.DELETE("/delete/:id", userController.Delete)
	userGroup.GET("/list", userController.List)

	// Electricity bill API
	eBillsGroup := v1.Group("/eBills")
	eBillsGroup.GET("/EstimateEBill", eBillsController.EstimateEBill)
	eBillsGroup.Use(middlewares.Auth(di))
	eBillsGroup.GET("/EAmount", eBillsController.EAmount)
	eBillsGroup.GET("/Report", eBillsController.ReportMonthly)

	// api for staff
	staffGroup := v1.Group("/staff")
	staffGroup.Use(middlewares.Auth(di))
	staffGroup.GET("/ReportList", staffController.GetReportList)
	staffGroup.GET("/ReportEUsed", staffController.GetReportEUsed)
	return r, nil
}
