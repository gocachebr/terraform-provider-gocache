package gocache

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSmartRulesSettings() *schema.Resource {
	matchSchema := getSchema("smart_rule_match")
	actionSchema := getSchema("smart_rule_settings_action")
	metadataSchema := getSchema("smart_rule_metadata")

	auxScheme := map[string]*schema.Schema{
		"domain": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"bulk_id": &schema.Schema{
			Type:     schema.TypeInt,
			Computed: true,
		},
		"smart_rule": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"match": &schema.Schema{
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: matchSchema,
						},
					},
					"action": &schema.Schema{
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Resource{
							Schema: actionSchema,
						},
					},
					"metadata": &schema.Schema{
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: metadataSchema,
						},
					},
				},
			},
		},
	}

	return &schema.Resource{
		CreateContext: resourceSmartRulesSettingsCreate,
		ReadContext:   resourceSmartRulesSettingsRead,
		UpdateContext: resourceSmartRulesSettingsUpdate,
		DeleteContext: resourceSmartRulesSettingsDelete,
		Schema:        auxScheme,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceSmartRulesSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return createSmartRulesBulk(ctx,d,m,"settings")
}

func resourceSmartRulesSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return readSmartRulesBulk(ctx,d,m,"settings")
}

func resourceSmartRulesSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return updateSmartRulesBulk(ctx,d,m,"settings")
}

func resourceSmartRulesSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return deleteSmartRulesBulk(ctx,d,m,"settings")
}
