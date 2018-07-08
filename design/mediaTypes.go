package design

import (
. "github.com/goadesign/goa/design"
. "github.com/goadesign/goa/design/apidsl"
)

var accountRegister = MediaType("application/vnd.goa.example.account.register+json", func() {
	Description("")
	Attributes(func() {
		Attribute("userId", Integer, "Unique User ID")
		Attribute("token", String, "Unique User Token")
		Attribute("name", String, "Name of user")
		Required("userId", "token", "name")
	})
	View("default", func() {
		Attribute("userId")
		Attribute("token")
		Attribute("name")
	})
})

var accountLogIn = MediaType("application/vnd.goa.example.account.login+json", func() {
	Description("")
	Attributes(func() {
		Attribute("userId", Integer, "Unique User ID")
		Attribute("token", String, "Unique User Token")
		Attribute("name", String, "Name of user")
		Required("userId", "token", "name")
	})
	View("default", func() {
		Attribute("userId")
		Attribute("token")
		Attribute("name")
	})
})

var accountLogOut = MediaType("application/vnd.goa.example.account.logout+json", func() {
	Description("")
	Attributes(func() {
		Attribute("status", Boolean, "action status")
	})
	View("default", func() {
		Attribute("status")
	})
})

var accountCurrentUser = MediaType("application/vnd.goa.example.account.currentUser+json", func() {
	Description("")
	Attributes(func() {
		Attribute("userId", Integer, "Unique User ID")
		Attribute("name", String, "Name of user")
		Required("userId", "name")
	})
	View("default", func() {
		Attribute("userId")
		Attribute("name")
	})
})

var accountOK = MediaType("application/vnd.goa.example.account.ok+json", func() {
	Description("")
	Attributes(func() {
		Attribute("status", Boolean, "action status")
	})
	View("default", func() {
		Attribute("status")
	})
})