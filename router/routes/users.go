package routes

import (
	controllers "barber/controller/users"
	"net/http"
)

var userRouters = []Route{
	{
		URI:                "/login",
		Method:             http.MethodPost,
		Func:               controllers.Login,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Func:               controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/me",
		Method:             http.MethodGet,
		Func:               controllers.GetUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users",
		Method:             http.MethodPatch,
		Func:               controllers.UpdateUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/subscription",
		Method:             http.MethodGet,
		Func:               controllers.CheckSubscription,
		NeedAuthentication: true,
	},
}
