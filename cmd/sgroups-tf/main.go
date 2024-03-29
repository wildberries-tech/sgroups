package main

import (
	"os"
	"sync"

	details "github.com/H-BF/sgroups/cmd/sgroups-tf/internal"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		//Level:      hclog.Debug,
		Level:      hclog.Error,
		JSONFormat: true,
		Output:     os.Stdout,
		Mutex:      new(sync.Mutex),
	})
	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return &schema.Provider{
				Schema:               details.SGroupsConfigSchema(),
				ConfigureContextFunc: details.SGroupsConfigure,
				ResourcesMap: map[string]*schema.Resource{
					details.RcNetwork:  details.SGroupsRcNetwork(),
					details.RcSG:       details.SGroupsRcSG(),
					details.RcRule:     details.SGroupsRcRule(),
					details.RcFqdnRule: details.SGroupsRcFqdnRule(),

					details.RcNetworks:  details.SGroupsRcNetworks(),
					details.RcSGs:       details.SGroupsRcSGs(),
					details.RcRules:     details.SGroupsRcRules(),
					details.RcFqdnRules: details.SGroupsRcFqdnRules(),
				},
			}
		},
		Logger: logger,
		//NoLogOutputOverride: true,
		//Debug:               true,
		//ProviderAddr:        "registry.terraform.io/h-bf/sgroups",
	}
	plugin.Serve(opts)
}
