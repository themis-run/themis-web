package handle

import (
	"time"

	"go.themis.run/themis-web/core"
	"go.themis.run/themis-web/model"
	"go.themis.run/themisclient"
)

type KeyValueHandler struct {
	client *themisclient.Client
}

func NewKeyValueHandler() *KeyValueHandler {
	return &KeyValueHandler{
		client: core.GetThemisClient(),
	}
}

func (k *KeyValueHandler) ListKV() ([]model.KeyValue, error) {
	kvs, err := k.client.ListAllKV()
	if err != nil {
		return nil, err
	}

	res := make([]model.KeyValue, 0)
	for _, kv := range kvs {
		res = append(res, model.KeyValue{
			Key:        kv.Key,
			Value:      string(kv.Value),
			CreateTime: time.UnixMilli(kv.CreateTime),
			TTL:        time.Duration(kv.Ttl),
		})
	}

	return res, nil
}

func (k *KeyValueHandler) SetKV(kv model.KeyValue) error {
	return k.client.Set(kv.Key, kv.Value)
}
