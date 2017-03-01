package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var responseBaseType = Type("responseBaseType", func() {
	Attribute("id", Integer, func() {
		Example(6631)
	})
	Attribute("name", String, func() {
		Example("テスト名前")
		MinLength(1)
	})
	Attribute("activeFlag", Boolean)
	Attribute("createdAt", DateTime, func() {
		Example("2017-02-14T00:00:10+09:00")
	})
	Attribute("children", ArrayOf(base2Type))
})

var baseType = Type("baseType", func() {
	Attribute("field1", Integer)
	Attribute("field2", Boolean)
	Required("field1", "field2")
})

var base2Type = Type("base2Type", func() {
	//var base2Type = Type("mixType", func() {
	Reference(baseType)
	Attribute("field1")
	Attribute("field2")
	Required("field1", "field2")
})

var resTestList = MediaType("application/test.list+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("res", ArrayOf(responseBaseType))
	})
	View("default", func() {
		Attribute("res")
	})
})

var _ = Resource("v1_bugs", func() {
	BasePath("/v1/bugs")
	Action("list", func() {
		Routing(GET(""))
		Description("bug list response")
		Params(func() {
			Param("page", Integer, "e.g.) page=1", func() {
				Default(1)
			})
			Param("perPage", Integer, "e.g.) perPage=10", func() {
				Default(10)
			})
		})
		Response(OK, func() {
			Media(resTestList)
		})
	})
})
