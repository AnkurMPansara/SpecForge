package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(router *gin.Engine) {
	router.GET("/ping", handlePing)
}

//handles the /ping route and responds with a JSON object
func handlePing(ctx *gin.Context) {
    response := map[string]interface{}{
        "code": 200,
        "msg":  "Success",
    }
    ctx.JSON(http.StatusOK, response)
}