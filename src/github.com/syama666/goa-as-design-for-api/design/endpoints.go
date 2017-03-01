package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("go-as-design-for-api", func() {
	Title("goa :: Design-first API Generation")
	Description("example design")
	Host("localhost")
	Scheme("http")
})

var _ = Resource("v1_kings", func() {
	BasePath("/v1/kings")
	Action("list", func() {
		Routing(GET(""))
		Description("王様リスト取得")
		Params(func() {
			Param("page", Integer, "e.g.) page=1", func() {
				Default(1)
			})
			Param("perPage", Integer, "e.g.) perPage=10", func() {
				Default(10)
			})
			Param("sort", String, "e.g.) sort=kingId or sort=-kingId", func() {
				Pattern("^-?[0-9a-zA-Z]+$")
				Default("-kingId")
			})
		})
		Response(OK, func() {
			Media(resKingList)
		})
	})
	Action("update", func() {
		Routing(PUT("/:kingId"))
		Description("王様1件更新")
		Params(func() {
			Param("kingId", Integer)
		})
		Payload(updateKing)
		Response(OK, func() {
			Media(resUpdateKing)
		})
	})
})

var _ = Resource("v1_swagger", func() {
	BasePath("/v1")
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Action("viewJSON", func() {
		Routing(GET("/swagger.json"))
		Description("view swagger.json")
		Response(OK)
		Response(NotFound)
		Response(InternalServerError)
	})
	Metadata("swagger:generate", "false")
})
