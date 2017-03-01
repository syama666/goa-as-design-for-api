package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var baseKing = Type("baseKing", func() {
	Attribute("kingId", Integer, func() {
		Example(1192)
	})
	Attribute("name", String, func() {
		Example("Louis XIV")
	})
	Attribute("age", Integer, func() {
		Example(76)
	})
	Attribute("dead", Boolean)
	Attribute("birthDay", DateTime, func() {
		Example("1638-09-05T00:00:10+09:00")
	})
	Attribute("deadDay", DateTime, func() {
		Example("1715-09-01T00:00:10+09:00")
	})
	Required(
		"name",
		"age",
		"dead",
		"birthDay",
	)
})

var responseKing = Type("responseKing", func() {
	Reference(baseKing)
	Attribute("kingId")
	Attribute("name")
	Attribute("age")
	Attribute("dead")
	Attribute("birthDay")
	Attribute("deadDay")
	Required(
		"kingId",
		"name",
		"age",
		"dead",
		"birthDay",
	)
})

var updateKing = Type("updateKing", func() {
	Reference(baseKing)
	Attribute("kingId")
	Attribute("name")
	Attribute("age")
	Attribute("dead")
	Attribute("birthDay")
	Attribute("deadDay")
	Required(
		"kingId",
		"name",
		"age",
		"dead",
		"birthDay",
	)
})

var resKingList = MediaType("application/example.king.list+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("kings", ArrayOf(responseKing))
	})
	View("default", func() {
		Attribute("kings")
	})
})

var resUpdateKing = MediaType("application/example.king.update+json", func() {
	ContentType("application/json")
	Attributes(func() {
		Attribute("king", responseKing)
	})
	View("default", func() {
		Attribute("king")
	})

})
