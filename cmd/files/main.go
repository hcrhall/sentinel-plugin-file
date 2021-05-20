package main

import (
	"github.com/hashicorp/sentinel-sdk/rpc"
	"github.com/sentinel-plugin-file/files"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		ImportFunc: files.New,
	})
}
