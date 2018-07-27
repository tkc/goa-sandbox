package main

import (
	"github.com/goadesign/goa"
	"github.com/tkc/goa-sandbox/app"
)

// JWTController implements the jwt resource.
type JWTController struct {
	*goa.Controller
}

// NewJWTController creates a jwt controller.
func NewJWTController(service *goa.Service) *JWTController {
	return &JWTController{Controller: service.NewController("JWTController")}
}

// Secure runs the secure action.
func (c *JWTController) Secure(ctx *app.SecureJWTContext) error {
	// JWTController_Secure: start_implement

	// Put your logic here

	res := &app.GoaExampleAccount{}
	return ctx.OK(res)
	// JWTController_Secure: end_implement
}

// SignIn runs the signIn action.
func (c *JWTController) SignIn(ctx *app.SignInJWTContext) error {
	// JWTController_SignIn: start_implement

	// Put your logic here

	res := &app.GoaExampleAccount{}
	return ctx.OK(res)
	// JWTController_SignIn: end_implement
}

// Unsecure runs the unsecure action.
func (c *JWTController) Unsecure(ctx *app.UnsecureJWTContext) error {
	// JWTController_Unsecure: start_implement

	// Put your logic here

	res := &app.GoaExampleAccount{}
	return ctx.OK(res)
	// JWTController_Unsecure: end_implement
}
