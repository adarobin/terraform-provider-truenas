package truenas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		TrueNASHost: d.Get("truenas_host").(string),
	}

	return config, nil
}

// Provider returns a terraform resource provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"truenas_host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRUENAS_HOST", nil),
				Description: "The hostname or IP of the TrueNAS management interface",
			},
		},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}
