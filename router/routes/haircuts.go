package routes

import (
	controllers "barber/controller/haircuts"
	"net/http"
)

var haircutRouters = []Route{
	{
		URI:                "/haircuts",
		Method:             http.MethodPost,
		Func:               controllers.CreateHaircut,
		NeedAuthentication: true,
	},
	{
		URI:                "/haircuts",
		Method:             http.MethodGet,
		Func:               controllers.GetHaircuts,
		NeedAuthentication: true,
	},
	{
		URI:                "/haircuts/{ID}",
		Method:             http.MethodPatch,
		Func:               controllers.UpdateHaircut,
		NeedAuthentication: true,
	},
}
