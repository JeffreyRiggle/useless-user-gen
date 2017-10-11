package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.Handler)
	}

	return router
}

var routes = Routes{
	Route{
		"Create",
		"POST",
		"/uug/v1/create",
		HandleCreate,
	},
	Route{
		"Update",
		"PUT",
		"/uug/v1/updaterandom",
		HandleUpdate,
	},
	Route{
		"GetUsers",
		"GET",
		"/uug/v1/users",
		HandleGet,
	},
	Route{
		"Clear",
		"DELETE",
		"/uug/v1/clear",
		HandleClear,
	},
}
