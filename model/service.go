package model

type Service struct {
	Name          string `json:"name"`
	ClusterCount  int    `json:"cluster_count"`
	InstanceCount int    `json:"instance_count"`
}

type ServiceResponse struct {
	ServiceList  []Service `json:"services"`
	ServiceCount int       `json:"service_count"`
}
