package router

import (
	"github.com/gin-gonic/gin"
	"go.themis.run/themis-web/router/api"
)

func Init() *gin.Engine {
	r := gin.Default()

	service := api.NewServiceAPI()
	r.GET("/service", service.FindAllService)
	r.GET("/service/:serviceName", service.FindInstance)

	node := api.NewNodeAPI()
	r.GET("/node", node.NodeInfo)
	r.GET("/node/:node_name", node.KeyValueByNode)

	keyvalue := api.NewKeyValueAPI()
	r.GET("/kv", keyvalue.ListKV)
	r.POST("/kv", keyvalue.SetKV)

	return r
}
