package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var AccountPayload = Type("AccountPayload", func() {
	Member("email", String, "", func(){
		Format("email")
	})
	Member("password", String, "password", func(){
		MinLength(8)
		MaxLength(20)
		Pattern(`^[a-zA-Z0-9!-/:-@¥{}±§|[-{-~]+$`)
	})
	Required("email", "password")
})

var AccountListPayload = Type("AccountListPayload", func() {
	Attribute("ids", ArrayOf(String), "", func(){})
	Attribute("offset", Integer, "", func(){})
	Required("ids", "offset")
})
