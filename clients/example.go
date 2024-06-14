package clients

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
	return
}
