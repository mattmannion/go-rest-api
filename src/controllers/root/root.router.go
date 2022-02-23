package controller_root

import (
	"mm/pkg/src/controllers"
	"net/http"
)

var RootRouter http.HandlerFunc = controllers.ControlSwitch(
	controllers.Methods{
		Get:  RootGet,
		Post: RootPost,
	})
