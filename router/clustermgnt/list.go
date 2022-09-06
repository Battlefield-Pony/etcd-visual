package clustermgnt

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func list(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
