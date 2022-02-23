package users

import (
	"mm/pkg/src/controllers"
	"net/http"
)

var Router http.HandlerFunc = controllers.ControlSwitch(
	controllers.Methods{})
