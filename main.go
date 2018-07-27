//go:generate goagen bootstrap -d github.com/tkc/goa-sandbox/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tkc/goa-sandbox/app"
)

func main() {
	// Create service
	service := goa.New("goa sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "jwt" controller
	c2 := NewJWTController(service)
	app.MountJWTController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
