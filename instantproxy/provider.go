package instantproxy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ipac "github.com/dthung1602/instant-proxy-api-client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:               providerSchema(),
		ConfigureContextFunc: configureContext,
		ResourcesMap:         map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"instantproxy_proxies":        dataSourceProxies(),
			"instantproxy_authorized_ips": dataSourceAuthorizedIPs(),
		},
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("INSTANT_PROXY_USERNAME", ""),
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Sensitive:   true,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("INSTANT_PROXY_PASSWORD", ""),
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			DefaultFunc: schema.EnvDefaultFunc("INSTANT_PROXY_ENDPOINT", ""),
		},
	}
}

func configureContext(_ context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	client := ipac.NewClient(
		data.Get("username").(string),
		data.Get("password").(string),
		data.Get("endpoint").(string),
	)
	err := client.Authenticate()
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return client, nil
}
