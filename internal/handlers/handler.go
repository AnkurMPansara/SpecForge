package handler

import (
	"github.com/gin-gonic/gin"
)

func HandleRoutes(router *gin.Engine) {
	for _, routeGroup := range RouteMappings {
		RequestHandlerGroup := router.Group(routeGroup.apiGroup)
		registerRoute(RequestHandlerGroup, routeGroup.Module)

	}
}

func registerRoute(apiGroup *gin.RouterGroup, endpoint []RouteHandler) {
	for _, handler := range endpoint {
		switch handler.Method {
		case "GET":
			apiGroup.GET(handler.Path, handler.HandlerFunc)
		case "POST":
			apiGroup.POST(handler.Path, handler.HandlerFunc)
		}
	}
}