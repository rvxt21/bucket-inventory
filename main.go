package main

import (
	"github.com/rvxt21/bucket-inventory/app"
	"github.com/rvxt21/bucket-inventory/config"
)

// main godoc
//
//	@title			Bucket Inventory API
//	@version		1.0
//	@BasePath		/
//	@description	File storage service: uploads go to S3, metadata to Postgres, downloads via presigned links.
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	app.App(cfg).Run()
}
