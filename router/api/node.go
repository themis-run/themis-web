package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.themis.run/themis-web/handle"
	"go.themis.run/themis-web/model"
)

type NodeAPI struct {
	handler *handle.NodeHandler
}

func NewNodeAPI() *NodeAPI {
	return &NodeAPI{
		handler: handle.NewNodeHandler(),
	}
}

func (n *NodeAPI) NodeInfo(c *gin.Context) {
	leader, nodes, err := n.handler.NodeInfoList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, model.Success(model.NodeInfoResponse{
		LeaderName: leader,
		NodeList:   nodes,
	}))
}

func (n *NodeAPI) KeyValueByNode(c *gin.Context) {
	nodeName := c.Param("node_name")
	kvs, err := n.handler.GetKeyValueByNodeName(nodeName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, model.Success(model.NodeKeyValueResponse{
		Name:         nodeName,
		KeyValueList: kvs,
	}))
}
