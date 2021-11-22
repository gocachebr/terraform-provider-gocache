package gocache

import (
	"context"
	gc "github.com/gocachebr/gocache-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The API key for operations.",
				DefaultFunc: schema.EnvDefaultFunc("GOCACHE_API_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"gocache_domain": resourceDomain(),
			"gocache_domain_dnssec": resourceDNSSEC(),
			"gocache_record": resourceRecord(),
			"gocache_ssl_certificate": resourceSSL(),
			"gocache_smart_rules_settings": resourceSmartRulesSettings(),
			"gocache_smart_rules_rewrite": resourceSmartRulesRewrite(),
			"gocache_smart_rules_firewall": resourceSmartRulesFirewall(),
		},
		DataSourcesMap:       map[string]*schema.Resource{
			"gocache_ip_ranges": dataSourceIpRanges(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("api_token").(string)

	var diags diag.Diagnostics

	c, err := gc.NewClient(token)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create GoCache client",
			Detail:   "Unable to validate api token",
		})
		return nil, diags
	}

	return c, diags
}
