package design

import (
. "github.com/goadesign/goa/design"
. "github.com/goadesign/goa/design/apidsl"
)

var accountMediaType = MediaType("application/vnd.goa.example.account+json", func() {
	Description("")

	Attributes(func() {
		Attribute("userId", Integer, "Unique User ID")
		Attribute("token", String, "Unique User Token")
		Attribute("name", String, "Name of user")
		Required("userId")
	})

	View("default", func() {
		Attribute("userId")
		Attribute("token")
		Attribute("name")
	})

	View("tiny", func() {
		Attribute("userId")
	})

})
