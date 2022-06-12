package instantproxy

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ipac "github.com/dthung1602/instant-proxy-api-client"
)

func dataSourceAuthorizedIPs() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceAuthorizedIPsRead,
		Schema:      dataSourceAuthorizedIPSchema(),
	}
}

func dataSourceAuthorizedIPSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"value": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func datasourceAuthorizedIPsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ipac.Client)
	ips, err := client.GetAuthorizedIPs()
	if err != nil {
		return diag.FromErr(err)
	}

	result := make([]string, len(ips), len(ips))
	for i, ip := range ips {
		result[i] = ip.String()
	}

	if err := data.Set("value", result); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diag.Diagnostics{}
}
