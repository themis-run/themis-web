package model

type Instance struct {
	ServiceName string            `json:"service_name"`
	ClusterName string            `json:"cluster_name"`
	IP          string            `json:"ip"`
	Port        int               `json:"port"`
	CreateTime  string            `json:"create_time"`
	IsHealthy   bool              `json:"is_healthy"`
	Meta        map[string]string `json:"meta"`
}

type InstancesResponse struct {
	InstanceList []Instance `json:"instances"`
	Count        int        `json:"instance_count"`
}
