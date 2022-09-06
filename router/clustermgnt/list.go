package clustermgnt

import (
	"etcd-visual/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func list(ctx *gin.Context) {

	var filter model.Cluster
	err := ctx.Bind(&filter)
	if err != nil {
		return
	}
	cluster, err := model.ListCluster(filter)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, cluster)
}
