package main

import (
	"github.com/goadesign/goa"
	"github.com/tkc/goa-sandbox/app"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

// CurrentUser runs the currentUser action.
func (c *AccountController) CurrentUser(ctx *app.CurrentUserAccountContext) error {
	// AccountController_CurrentUser: start_implement

	// Put your logic here

	res := &app.GoaExampleAccountCurrentuser{}
	return ctx.OK(res)
	// AccountController_CurrentUser: end_implement
}

// Login runs the login action.
func (c *AccountController) Login(ctx *app.LoginAccountContext) error {
	// AccountController_Login: start_implement

	// Put your logic here

	res := &app.GoaExampleAccountLogin{}
	return ctx.OK(res)
	// AccountController_Login: end_implement
}

// Logout runs the logout action.
func (c *AccountController) Logout(ctx *app.LogoutAccountContext) error {
	// AccountController_Logout: start_implement

	// Put your logic here

	res := &app.GoaExampleAccountLogout{}
	return ctx.OK(res)
	// AccountController_Logout: end_implement
}

// Register runs the register action.
func (c *AccountController) Register(ctx *app.RegisterAccountContext) error {
	// AccountController_Register: start_implement

	// Put your logic here

	res := &app.GoaExampleAccountRegister{}
	return ctx.OK(res)
	// AccountController_Register: end_implement
}
