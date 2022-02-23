package root

import (
	"mm/pkg/src/controllers"
	"net/http"
)

var Router http.HandlerFunc = controllers.ControlSwitch(
	controllers.Methods{
		Get:    get,
		Post:   post,
		Put:    put,
		Patch:  put,
		Delete: delete,
	})
