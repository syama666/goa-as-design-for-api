package controller

import (
	"github.com/goadesign/goa"
	"github.com/syama666/goa-as-design-for-api/gen/app"
	"github.com/syama666/goa-as-design-for-api/gen/swagger"
)

// V1SwaggerController implements the v1_swagger resource.
type V1SwaggerController struct {
	*goa.Controller
}

// NewV1SwaggerController creates a v1_swagger controller.
func NewV1SwaggerController(service *goa.Service) *V1SwaggerController {
	return &V1SwaggerController{Controller: service.NewController("V1SwaggerController")}
}

// ViewJSON runs the viewJSON action.
func (c *V1SwaggerController) ViewJSON(ctx *app.ViewJSONV1SwaggerContext) error {
	swaggerBytes, err := swagger.Asset("gen/swagger/swagger.json")
	if err != nil {
		c.Service.LogError(err.Error())
		ctx.InternalServerError()
		return nil
	}
	ctx.OK(swaggerBytes)

	return nil
}
