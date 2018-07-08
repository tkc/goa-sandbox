//go:generate goagen bootstrap -d github.com/tkc/goa-sandbox/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/tkc/goa-sandbox/app"
	"github.com/goadesign/goa/middleware"
)

func main() {

	service := goa.New("goa sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// TODO  : want to remove
	app.UseSigninBasicAuthMiddleware(service, NewBasicAuthMiddleware())

	jwtMiddleware, _ := NewJWTMiddleware()
	app.UseJWTMiddleware(service, jwtMiddleware)

	c := NewAccountController(service)
	app.MountAccountController(service, c)

	c2,_ := NewJWTController(service)
	app.MountJWTController(service, c2)

	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
