package instantproxy

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ipac "github.com/dthung1602/instant-proxy-api-client"
)

func dataSourceProxies() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceProxiesRead,
		Schema: map[string]*schema.Schema{
			"proxies": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceProxiesRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ipac.Client)
	proxies, err := client.GetProxies()
	if err != nil {
		return diag.FromErr(err)
	}

	result := make([]*map[string]interface{}, len(proxies), len(proxies))
	for i, proxy := range proxies {
		result[i] = &map[string]interface{}{
			"address": proxy.String(),
			"ip":      proxy.IP.String(),
			"port":    proxy.Port,
		}
	}

	if err := data.Set("proxies", result); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diag.Diagnostics{}
}
