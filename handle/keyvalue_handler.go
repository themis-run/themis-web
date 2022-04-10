package handle

import (
	"time"

	"go.themis.run/themis-web/core"
	"go.themis.run/themis-web/model"
	"go.themis.run/themisclient"
)

var timeLayout = "2006-01-02 15:04:05"

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
			CreateTime: transforTimestamp(kv.CreateTime),
			TTL:        time.Duration(kv.Ttl),
		})
	}

	return res, nil
}

func (k *KeyValueHandler) SetKV(kv model.KeyValue) error {
	return k.client.Set(kv.Key, kv.Value)
}

func transforTimestamp(timestamp int64) string {
	t := time.UnixMilli(timestamp)
	return t.Format(timeLayout)
}
