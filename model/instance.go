package model

import "time"

type Instance struct {
	ServiceName string            `json:"service_name"`
	ClusterName string            `json:"cluster_name"`
	IP          string            `json:"ip"`
	Port        int               `json:"port"`
	CreateTime  time.Time         `json:"create_time"`
	IsHealthy   bool              `json:"is_healthy"`
	Meta        map[string]string `json:"meta"`
}
