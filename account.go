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

	res := &app.GoaExampleAccount{}
	return ctx.OK(res)
	// AccountController_CurrentUser: end_implement
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	// AccountController_List: start_implement

	// Put your logic here

	res := app.GoaExampleAccountCollection{}
	return ctx.OK(res)
	// AccountController_List: end_implement
}

// Login runs the login action.
func (c *AccountController) Login(ctx *app.LoginAccountContext) error {
	// AccountController_Login: start_implement

	// Put your logic here

	return nil
	// AccountController_Login: end_implement
}
