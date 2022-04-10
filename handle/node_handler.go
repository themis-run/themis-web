package handle

import (
	"time"

	"go.themis.run/themis-web/core"
	"go.themis.run/themis-web/model"
	"go.themis.run/themisclient"
)

type NodeHandler struct {
	client *themisclient.Client
}

func NewNodeHandler() *NodeHandler {
	return &NodeHandler{
		client: core.GetThemisClient(),
	}
}

func (n *NodeHandler) NodeInfoList() (string, []model.Node, error) {
	servers := n.client.Info().Servers
	leader := n.client.Info().LeaderName
	nodeList := make([]model.Node, 0)

	for s := range servers {
		info, err := n.client.NodeInfo(s)
		if err != nil {
			return "", nil, err
		}

		nodeList = append(nodeList, model.Node{
			Name:        info.Name,
			Address:     info.Address,
			RaftAddress: info.RaftAddress,
			Term:        info.Term,
			Role:        info.Role,
			LogTerm:     info.LogTerm,
			LogIndex:    info.LogIndex,
		})
	}

	return leader, nodeList, nil
}

func (n *NodeHandler) GetKeyValueByNodeName(name string) ([]model.KeyValue, error) {
	kvList, err := n.client.SearchKVListFromNodeName(name)
	if err != nil {
		return nil, err
	}

	kvs := make([]model.KeyValue, 0)
	for _, kv := range kvList {
		kvs = append(kvs, model.KeyValue{
			Key:        kv.Key,
			Value:      string(kv.Value),
			CreateTime: transforTimestamp(kv.CreateTime),
			TTL:        time.Duration(kv.Ttl),
		})
	}

	return kvs, nil
}
