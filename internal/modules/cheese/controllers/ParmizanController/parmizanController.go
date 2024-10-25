package parmizanController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParmizanCheese(ctx *gin.Context) {
	response := map[string]interface{}{
		"code": 200,
		"msg":  "Cheese is divine",
	}
	ctx.JSON(http.StatusOK, response)
}