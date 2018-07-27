package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("account", func() {

	Action("login", func() {
		Routing(POST("login"))
		Description("")
		Payload(AccountPayload)
		Response(OK, AccountPayload)
		Response(BadRequest, ErrorMedia)
	})

	Action("currentUser", func() {
		Routing(GET("current_user"))
		Description("")
		Response(OK, accountMediaType)
		Response(BadRequest, ErrorMedia)
	})

	Action("list", func() {
		Routing(GET("list"))
		Description("")
		Params(func() {
			Param("ids", ArrayOf(Integer))
		})
		Payload(AccountListPayload)
		Response(OK, func() { Media(CollectionOf(accountMediaType), "default") })
		Response(BadRequest, ErrorMedia)
	})

})
