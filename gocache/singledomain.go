package gocache

// import (
// 	"context"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// 	gc "github.com/gocache/terraform-provider-gocache/gocache-go"

// 	_ "encoding/json"
// 	"strconv"
// 	"time"
// )

// func dataSourceSingleDomain() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: dataSourceSingleDomainRead,
// 		Schema: map[string]*schema.Schema{
// 			"response": {
// 				Type:     schema.TypeMap,
// 				Computed: true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 		},
// 	}
// }

// func dataSourceSingleDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	c := m.(*gc.API)

// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	resp, err := c.GetDomain("gocache.com.br")

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	if err := d.Set("response", resp.Response); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

// 	return diags
// }

// func dataSourceDomains() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: dataSourceCoffeesRead,
// 		Schema: map[string]*schema.Schema{

// 			"status_code": &schema.Schema{
// 				Type:     schema.TypeInt,
// 				Computed: true,
// 			},

// 			"response": &schema.Schema{
// 				Type:     schema.TypeSet,
// 				Computed: true,
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"domains": {
// 							Type:     schema.TypeList,
// 							Computed: true,
// 							Elem: &schema.Schema{
// 								Type: schema.TypeString,
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func dataSourceCoffeesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	c := m.(*gc.API)

// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	resp, err := c.ListDomains()

// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	if err := d.Set("status_code", resp.Status_code); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	list := resp.Response.(gc.DomainList)

// 	ois := make([]interface{}, 1)

// 	oi := make(map[string]interface{})

// 	is := make([]interface{}, len(list.Domains))
// 	for i, b := range list.Domains {
// 		is[i] = b
// 	}

// 	oi["domains"] = is

// 	ois[0] = oi

// 	if err := d.Set("response", ois); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

// 	return diags
// }
