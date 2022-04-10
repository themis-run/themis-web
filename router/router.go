package router

import (
	"github.com/gin-gonic/gin"
	"go.themis.run/themis-web/router/api"
)

func Init() *gin.Engine {
	r := gin.Default()

	service := api.NewServiceAPI()
	r.GET("/services", service.FindAllService)
	r.GET("/services/:serviceName", service.FindInstance)

	node := api.NewNodeAPI()
	r.GET("/node", node.NodeInfo)
	r.GET("/node/kv/:node_name", node.KeyValueByNode)

	keyvalue := api.NewKeyValueAPI()
	r.GET("/kv", keyvalue.ListKV)
	r.POST("/kv", keyvalue.SetKV)

	return r
}
