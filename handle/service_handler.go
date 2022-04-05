package handle

import (
	"go.themis.run/themis-web/core"
	"go.themis.run/themis-web/model"
	"go.themis.run/themisclient"
)

type ServiceHandler struct {
	client *themisclient.Client
}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{
		client: core.GetThemisClient(),
	}
}

func (s *ServiceHandler) FindAllService() ([]model.Service, error) {
	instances, err := s.client.DiscoveryService("")
	if err != nil {
		return nil, err
	}

	serviceMap := make(map[string]map[string]int)
	for _, v := range instances {
		if _, ok := serviceMap[v.ServiceName]; !ok {
			serviceMap[v.ServiceName] = make(map[string]int)
		}

		count := serviceMap[v.ServiceName][v.ClusterName]
		serviceMap[v.ServiceName][v.ClusterName] = count + 1
	}

	services := make([]model.Service, 0)
	for k, v := range serviceMap {
		count := 0
		for _, clusterCount := range v {
			count += clusterCount
		}

		services = append(services, model.Service{
			Name:          k,
			ClusterCount:  len(v),
			InstanceCount: count,
		})
	}

	return services, nil
}
