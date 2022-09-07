package clustermgnt

import (
	"etcd-visual/model"
	"etcd-visual/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	cluster, err := model.GetClusterWithAssociation(id)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, router.GenerateSuccessResponse(cluster))
}
