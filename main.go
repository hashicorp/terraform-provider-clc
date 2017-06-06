package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-clc/clc"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: clc.Provider})
}
