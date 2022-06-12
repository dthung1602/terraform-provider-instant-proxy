package instantproxy

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ipac "github.com/dthung1602/instant-proxy-api-client"
)

const uniqueTerraformId = "UNIQUE_TERRAFORM_ID_FOR_AUTHORIZED_IPS"

func resourceAuthorizedIPs() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceAuthorizedIPsRead,
		CreateContext: resourceAuthorizedIPsCreate,
		UpdateContext: resourceAuthorizedIPsUpdate,
		DeleteContext: resourceAuthorizedIPsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceAuthorizedIPSchema(),
	}
}

func resourceAuthorizedIPSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"value": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourceAuthorizedIPsRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return datasourceAuthorizedIPsRead(ctx, data, meta)
}

func resourceAuthorizedIPsCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ipac.Client)

	ips, err := parseIPs(data.Get("value"))
	if err != nil {
		return diag.FromErr(err)
	}

	if err := client.AddAuthorizedIPs(ips); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(uniqueTerraformId)
	return resourceAuthorizedIPsRead(ctx, data, meta)
}

func resourceAuthorizedIPsUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if data.HasChange("value") {
		client := meta.(*ipac.Client)

		ips, err := parseIPs(data.Get("value"))
		if err != nil {
			return diag.FromErr(err)
		}

		if err := client.SetAuthorizedIPs(ips); err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceAuthorizedIPsRead(ctx, data, meta)
}

func resourceAuthorizedIPsDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*ipac.Client)
	var diags diag.Diagnostics

	ips, err := parseIPs(data.Get("value"))

	if err != nil {
		return diag.FromErr(err)
	}

	if err := client.RemoveAuthorizedIPs(ips); err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")
	return diags
}

func parseIPs(input any) ([]net.IP, error) {
	ipstrs := input.([]interface{})
	ips := make([]net.IP, len(ipstrs))

	for i, ip := range ipstrs {
		ips[i] = net.ParseIP(ip.(string))
		if ips[i] == nil {
			return nil, fmt.Errorf("cannot parse IP: %s", ip)
		}
	}
	return ips, nil
}
