package gocache

import (
	"context"
	"errors"
	"fmt"
	gc "github.com/gocachebr/gocache-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceDomain() *schema.Resource {
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
	}
	for ind, content := range gc.GetAllFieldsAdjusted("domain") {
		schemeMode := schema.TypeString
		if content.Mode == "boolean" {
			schemeMode = schema.TypeBool
		} else if content.Mode == "int" {
			schemeMode = schema.TypeInt
		}
		auxScheme[ind] = &schema.Schema{
			Type:     schemeMode,
			Default:     content.Default,
			Optional: true,
		}
	}

	return &schema.Resource{
		CreateContext: resourceDomainCreate,
		ReadContext:   resourceDomainRead,
		UpdateContext: resourceDomainUpdate,
		DeleteContext: resourceDomainDelete,
		Schema:        auxScheme,
		Importer: &schema.ResourceImporter{
	      StateContext: schema.ImportStatePassthroughContext,
	    },
	}
}

func resourceDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*gc.API)

	var diags diag.Diagnostics
	domain := d.Get("domain").(string)

	changes := make(map[string]interface{})
	for ind, content := range gc.GetAllFieldsAdjusted("domain") {
		val := d.Get(ind)

		indNew := gc.GetFieldName(ind,"domain")

		if val != nil {
			if content.Mode == "boolean" || content.Mode == "int" {
				changes[indNew] = val
			} else {
				if val.(string) != "" {
					changes[indNew] = val
				}
			}
		}
	}

	d.SetId(domain)

	resp, err := c.CreateDomain(domain, changes)

	if resp == nil {
		return diag.FromErr(err)
	}
	if resp.HTTPStatusCode == 200 {
		diags = resourceDomainUpdate(ctx, d, m)
	} else {
		if err != nil {
			return diag.FromErr(err)
		}else {
			return diag.FromErr(fmt.Errorf("%d %s",resp.HTTPStatusCode,resp.Msg))
		}
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return diags
}

func resourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.API)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	domain := d.Id()

	resp, err := c.GetDomain(domain)

	if err != nil {
		return diag.FromErr(err)
	}

	if resp.HTTPStatusCode == 200 {
		vals := resp.Response.(map[string]interface{})

		d.Set("domain", domain)
		for ind, val := range vals {
			if gc.FieldExists(ind,"domain_reversed") {
				if err := d.Set(ind, val); err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  fmt.Sprintf("Invalid type: %v", ind),
						Detail:   err.Error(),
					})
				}
			}
		}
	}else {
		return diag.FromErr(fmt.Errorf("Status Code: %d Msg: %s",resp.HTTPStatusCode,resp.Msg))
	}

	return diags
}

func resourceDomainUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.API)

	domain := d.Id()

	if d.HasChange("domain") && (d.Get("domain").(string) != domain) {
		return diag.FromErr(errors.New("Cannot update field 'domain', you need to delete and create in different steps"))
	}

	changes := make(map[string]interface{})

	for ind, content := range gc.GetAllFieldsAdjusted("domain") {
		if d.HasChange(ind) {
			val := d.Get(ind)

			indNew := gc.GetFieldName(ind,"domain")

			if val != nil {
				if content.Mode == "boolean" || content.Mode == "int" {
					changes[indNew] = val
				} else {
					if val.(string) != "" {
						changes[indNew] = val
					}
				}
			}
		}
	}
	if len(changes) > 0 {

		_, err := c.UpdateDomain(domain, changes)

		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceDomainRead(ctx, d, m)
}

func resourceDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*gc.API)

	domain := d.Id()

	_, err := c.DeleteDomain(domain)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
