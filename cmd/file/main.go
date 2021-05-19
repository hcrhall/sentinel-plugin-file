package main

import (
	"github.com/hashicorp/sentinel-sdk/rpc"
	"github.com/hcrhall/sentinel-plugin-file/file"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		ImportFunc: file.New,
	})
}
