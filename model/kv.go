package model

import "time"

type KeyValue struct {
	Key        string        `json:"key"`
	Value      string        `json:"value"`
	CreateTime time.Time     `json:"create_time"`
	TTL        time.Duration `json:"ttl"`
}
