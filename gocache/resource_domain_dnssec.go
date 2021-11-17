package gocache

import (
	"context"
	"fmt"
	gc "github.com/gocachebr/gocache-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"strings"
)

func resourceDNSSEC() *schema.Resource {
	auxScheme := map[string]*schema.Schema{
		"domain": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"last_updated": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"digest": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"digest_type": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"digest_type_name": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"algorithm": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"public_key": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"key_tag": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
	}

	return &schema.Resource{
		CreateContext: resourceDNSSECCreate,
		ReadContext:   resourceDNSSECRead,
		UpdateContext: resourceDNSSECUpdate,
		DeleteContext: resourceDNSSECDelete,
		Schema:        auxScheme,
		Importer: &schema.ResourceImporter{
	      StateContext: schema.ImportStatePassthroughContext,
	    },
	}
}

func resourceDNSSECCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*gc.API)

	var diags diag.Diagnostics
	domain := d.Get("domain").(string)

	d.SetId(domain)

	resp, err := c.CreateDNSSEC(domain)

	if resp == nil {
		return diag.FromErr(err)
	}

	if resp.HTTPStatusCode == 200 {
		vals := resp.Response.(map[string]interface{})

		d.Set("domain", domain)
		for ind, val := range vals {
			if ind == "digest_type" {
				digest_type := strings.Split(val.(string)," ")
				num, err := strconv.Atoi(digest_type[0])
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  fmt.Sprintf("Cannot convert string to int , param:%v val:%v",ind,val),
						Detail:   err.Error(),
					})
				}
				d.Set("digest_type",num)
				d.Set("digest_type_name",digest_type[1])
			} else if ind == "algorithm" || ind == "key_tag" {
				num, err := strconv.Atoi(val.(string))
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  fmt.Sprintf("Cannot convert string to int , param:%v val:%v",ind,val),
						Detail:   err.Error(),
					})
				}
				d.Set(ind,num)
			}else{
				d.Set(ind,val)
			}
		}
	}

	return diags
}

func resourceDNSSECRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.API)

	var diags diag.Diagnostics
	domain := d.Get("domain").(string)

	d.SetId(domain)

	resp, err := c.GetDNSSEC(domain)

	if resp == nil {
		return diag.FromErr(err)
	}


	if resp.HTTPStatusCode == 200 {
		vals := resp.Response.(map[string]interface{})

		d.Set("domain", domain)
		for ind, val := range vals {
			if ind == "digest_type" {
				digest_type := strings.Split(val.(string)," ")
				num, err := strconv.Atoi(digest_type[0])
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  fmt.Sprintf("Cannot convert string to int , param:%v val:%v",ind,val),
						Detail:   err.Error(),
					})
				}
				d.Set("digest_type",num)
				d.Set("digest_type_name",digest_type[1])
			} else if ind == "algorithm" || ind == "key_tag" {
				num, err := strconv.Atoi(val.(string))
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  fmt.Sprintf("Cannot convert string to int , param:%v val:%v",ind,val),
						Detail:   err.Error(),
					})
				}
				d.Set(ind,num)
			}else{
				d.Set(ind,val)
			}
		}
	}

	return diags
}

func resourceDNSSECUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceDNSSECCreate(ctx, d, m)
}

func resourceDNSSECDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*gc.API)

	domain := d.Id()

	resp, err := c.DeleteDNSSEC(domain)

	if err != nil {
		return diag.FromErr(err)
	}else if resp.HTTPStatusCode == 428 {
		return diag.FromErr(fmt.Errorf("%s",resp.Msg))
	}

	return diags
}
