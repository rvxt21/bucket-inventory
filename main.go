package main

import (
	"github.com/rvxt21/bucket-inventory/app"
	"github.com/rvxt21/bucket-inventory/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	app.App(cfg).Run()
}
