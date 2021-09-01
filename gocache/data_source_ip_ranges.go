package gocache

import (
	"strconv"
	"context"
	gc "github.com/gocachebr/gocache-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func dataSourceIpRanges() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpRangesRead,
		Schema: map[string]*schema.Schema{
			"ipv4_cidr_blocks":{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cidr_blocks":{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIpRangesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.API)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resp, err := c.ListIpRanges()

	if err != nil {
		return diag.FromErr(err)
	}


	d.Set("ipv4_cidr_blocks", resp)
	d.Set("cidr_blocks", resp)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}