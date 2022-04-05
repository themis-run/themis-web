package model

type Service struct {
	Name          string `json:"name"`
	ClusterCount  string `json:"cluster_count"`
	InstanceCount string `json:"instance_count"`
}
