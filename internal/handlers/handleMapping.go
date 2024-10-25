package handler

import (
	parmizanController "SpecForge_api_backend/internal/modules/cheese/controllers/ParmizanController"
	mozerellaController "SpecForge_api_backend/internal/modules/cheese/controllers/mozerellaController"

	"github.com/gin-gonic/gin"
)

type RouteGroupDefintion struct {
	apiGroup  string
	Module    []RouteHandler
}

type RouteHandler struct {
	Method  string
	Path string
	HandlerFunc gin.HandlerFunc
}

var RouteMappings = []RouteGroupDefintion{
	{"/cheese", cheese},
}

var cheese = []RouteHandler{
	{"GET","/mozerella", mozerellaController.MozerellaCheese},
	{"GET","/parmizan", parmizanController.ParmizanCheese},
}