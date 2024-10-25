package mozerellaController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MozerellaCheese(ctx *gin.Context) {
	response := map[string]interface{}{
		"code": 200,
		"msg":  "Cheese is tasty",
	}
	ctx.JSON(http.StatusOK, response)
}