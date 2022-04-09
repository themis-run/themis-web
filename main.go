package main

import (
	"go.themis.run/themis-web/conf"
	"go.themis.run/themis-web/router"
)

func main() {
	conf.Setup()
	r := router.Init()

	if err := r.Run(conf.Port()); err != nil {
		panic(err)
	}
}
