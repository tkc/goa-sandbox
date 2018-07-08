package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access")
})

var SigninBasicAuth = BasicAuthSecurity("SigninBasicAuth")

var _ = Resource("jwt", func() {

	Description("This resource uses JWT to secure its endpoints")
	DefaultMedia(OK)

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("signin", func() {
		Description("Creates a valid JWT")
		Security(SigninBasicAuth)
		Routing(POST("/jwt/signin"))
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(OK, accountRegister)
		Response(Unauthorized, ErrorMedia)
	})

	Action("secure", func() {
		Description("This action is secured with the jwt scheme")
		Routing(GET("/jwt"))
		Params(func() {
			Param("fail", Boolean, "Force auth failure via JWT validation middleware")
		})
		Response(OK, accountOK)
		Response(Unauthorized, ErrorMedia)
	})

	Action("unsecure", func() {
		Description("This action does not require auth")
		Routing(GET("/jwt/unsecure"))
		NoSecurity()
		Response(OK, accountOK)
		Response(Unauthorized, ErrorMedia)
	})
})