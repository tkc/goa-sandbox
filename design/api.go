package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa sample", func() {

	Title("goa-sandbox")
	Description("sample")

	Contact(func() {
		Name("tkc")
		Email("iijima@onga.jp")
		URL("https://github.com/tkc/goa-sandbox/issues")
	})

	License(func() {
		Name("MIT")
		URL("https://github.com/tkc/goa-sandbox/blob/master/LICENSE")
	})

	Docs(func() {
		Description("wiki")
		URL("https://github.com/tkc/goa-sandbox/wiki")
	})

	Host("localhost:8081")
	//Scheme("http", "https")

	Scheme("http")
	BasePath("/")

	Origin("http://localhost:8080/swagger", func() {
		Expose("X-Time")
		Methods("GET", "POST", "PUT", "DELETE")
		Headers( "Content-Type")
		MaxAge(600)
		Credentials()
	})

})
