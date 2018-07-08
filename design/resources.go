package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("account", func() {

	Action("register", func() {
		Routing(POST("register"))
		Description("")
		Params(func() {
			Param("name", String, "")
			Param("email", String, "")
			Param("passWord", String, "")
		})
		//Required("name")
		//Required("email")
		//Required("passWord")
		Response(OK, accountRegister)
		Response(BadRequest, ErrorMedia)
	})

	Action("login", func() {
		Routing(POST("login"))
		Description("")
		Params(func() {
			Param("email", String, "")
			Param("passWord", String, "")
		})
		//Required("email")
		//Required("passWord")
		Response(OK, accountLogIn)
		Response(BadRequest, ErrorMedia)
	})

	Action("logout", func() {
		Routing(POST("logout"))
		Description("")
		Response(OK, accountLogOut)
		Response(BadRequest, ErrorMedia)
	})

	Action("currentUser", func() {
		Routing(GET("currentuser"))
		Description("")
		Params(func() {
			Param("token", UUID, "")
		})
		//Required("token")
		Response(OK, accountCurrentUser)
		Response(BadRequest, ErrorMedia)
	})
})
