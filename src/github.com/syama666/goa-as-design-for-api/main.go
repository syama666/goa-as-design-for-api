//go:generate goagen bootstrap -d github.com/syama666/goa-as-design-for-api/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/syama666/goa-as-design-for-api/controller"
	"github.com/syama666/goa-as-design-for-api/gen/app"
)

func main() {
	// Create service
	service := goa.New("go-as-design-for-api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "v1_kings" controller
	c := controller.NewV1KingsController(service)
	app.MountV1KingsController(service, c)
	// Mount "v1_swagger" controller
	c2 := controller.NewV1SwaggerController(service)
	app.MountV1SwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":10080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
