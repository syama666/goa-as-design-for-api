package controller

import (
	"github.com/goadesign/goa"
	"github.com/syama666/goa-as-design-for-api/gen/app"
)

// V1BugsController implements the v1_bugs resource.
type V1BugsController struct {
	*goa.Controller
}

// NewV1BugsController creates a v1_bugs controller.
func NewV1BugsController(service *goa.Service) *V1BugsController {
	return &V1BugsController{Controller: service.NewController("V1BugsController")}
}

// List runs the list action.
func (c *V1BugsController) List(ctx *app.ListV1BugsContext) error {
	// V1BugsController_List: start_implement

	// Put your logic here

	// V1BugsController_List: end_implement
	res := &app.TestList{}
	return ctx.OK(res)
}
