package main

import (
	"etcd-visual/router/clustermgnt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	InitRouter(r)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}

func InitRouter(engine *gin.Engine) {
	clustermgnt.Routes.RegisterToEngine(engine)
}
