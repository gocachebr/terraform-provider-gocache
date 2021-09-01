package gocache

import (
	"context"
	"fmt"
	gc "github.com/gocachebr/gocache-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
	"strings"

)

func resourceSSL() *schema.Resource {
	auxScheme := map[string]*schema.Schema{
		"domain": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"certificate": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"private_key": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		"expiration_date": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"issue_date": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_updated": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}

	return &schema.Resource{
		CreateContext: resourceSSLCreate,
		ReadContext:   resourceSSLRead,
		UpdateContext: resourceSSLCreate,
		DeleteContext: resourceSSLDelete,
		Schema:        auxScheme,
	}
}

func resourceSSLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*gc.API)

	var diags diag.Diagnostics
	domain := d.Get("domain").(string)

	res, err := c.ListCertificates(domain)

	if err != nil {
		return diag.FromErr(err)
	}
	
	certData := res.Response.(gc.CertificateResult)

	hasCustom := false
	for _, cert := range certData.Certificates {
		if (cert["type"] == "custom"){
			hasCustom = true
			d.Set("enabled", cert["enabled"])
			d.Set("domain", domain)
			d.Set("expiration_date", cert["expirationDate"])
			d.Set("issue_date", cert["issueDate"])

			d.Set("last_updated", time.Now().Format(time.RFC850))
			break
		}
	}
	
	if !hasCustom {
		return diag.FromErr(fmt.Errorf("No certificates installed"))
	}
	
	
	return diags
}


func resourceSSLCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*gc.API)

	var diags diag.Diagnostics
	domain := d.Get("domain").(string)
	certificate := d.Get("certificate").(string)
	private_key := d.Get("private_key").(string)
	enabled := d.Get("enabled").(bool)

	res, err := c.CreateCertificate(domain, certificate, private_key, enabled)

	if err != nil {
		return diag.FromErr(err)
	}
	
	certData := res.Response.(gc.CertificateResult)

	hasCustom := false
	for _, cert := range certData.Certificates {
		if (cert["type"] == "custom"){
			hasCustom = true
			d.SetId(fmt.Sprintf("%v/%v", domain, cert["cn"]))
			if (enabled){
				c.ActivateCertificate(domain, "custom", true)
			}else{
				c.ActivateCertificate(domain, "auto", true)
			}
			d.Set("enabled", enabled)
			d.Set("private_key", private_key)
			d.Set("certificate", certificate)
			d.Set("domain", domain)
			d.Set("expiration_date", cert["expirationDate"])
			d.Set("issue_date", cert["issueDate"])

			d.Set("last_updated", time.Now().Format(time.RFC850))
			break
		}
	}

	if !hasCustom {
		return diag.FromErr(fmt.Errorf("Failed to install custom certificate. This might be a internal error, please contact support@gocache.com.br about this error"))
	}

	return diags
}

func resourceSSLDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*gc.API)

	s := strings.Split(d.Id(), "/")

	_, err := c.DeleteCertificate(s[0])

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
