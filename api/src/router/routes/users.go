package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Function:           controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodGet,
		Function:           controllers.SearchUsers,
		NeedAuthentication: false,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Function:           controllers.SearchUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodPut,
		Function:           controllers.UpdateUser,
		NeedAuthentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodDelete,
		Function:           controllers.DeleteUser,
		NeedAuthentication: true,
	},
}
