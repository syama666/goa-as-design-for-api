package controller

import (
	"github.com/goadesign/goa"
	"github.com/syama666/goa-as-design-for-api/gen/app"
)

// V1KingsController implements the v1_kings resource.
type V1KingsController struct {
	*goa.Controller
}

// NewV1KingsController creates a v1_kings controller.
func NewV1KingsController(service *goa.Service) *V1KingsController {
	return &V1KingsController{Controller: service.NewController("V1KingsController")}
}

// List runs the list action.
func (c *V1KingsController) List(ctx *app.ListV1KingsContext) error {
	// V1KingsController_List: start_implement

	// Put your logic here

	// V1KingsController_List: end_implement
	res := &app.ExampleKingList{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *V1KingsController) Update(ctx *app.UpdateV1KingsContext) error {
	// V1KingsController_Update: start_implement

	// Put your logic here

	// V1KingsController_Update: end_implement
	res := &app.ExampleKingUpdate{}
	return ctx.OK(res)
}
