package core

import (
	"sync"

	"go.themis.run/themis-web/conf"
	"go.themis.run/themisclient"
)

var client *themisclient.Client
var one sync.Once

func GetThemisClient() (*themisclient.Client, error) {
	var err error
	one.Do(func() {
		serverName, address := conf.GetThemisServer()
		opt := []themisclient.Option{
			themisclient.WithServerName(serverName),
			themisclient.WithServerAddress(address),
		}

		config := themisclient.NewConfigration(opt...)

		client, err = themisclient.NewClient(config)
	})
	return client, err
}
