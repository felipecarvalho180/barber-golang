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
}
