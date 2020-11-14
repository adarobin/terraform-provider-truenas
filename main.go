package main

import (
	"github.com/adarobin/terraform-provider-truenas/truenas"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: truenas.Provider})
}
