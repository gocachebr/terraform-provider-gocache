package gocache

import (
	"context"
	"errors"
	"fmt"
	gc "github.com/gocachebr/gocache-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
	"strings"
)

func resourceRecord() *schema.Resource {
	auxScheme := map[string]*schema.Schema{
		"domain": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"record_id": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"proxied": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		"ttl": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"content": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"priority": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"weight": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"last_updated": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}

	return &schema.Resource{
		CreateContext: resourceRecordCreate,
		ReadContext:   resourceRecordRead,
		UpdateContext: resourceRecordUpdate,
		DeleteContext: resourceRecordDelete,
		Schema:        auxScheme,
		Importer: &schema.ResourceImporter{
	      StateContext: schema.ImportStatePassthroughContext,
	    },
	}
}

func resourceRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*gc.API)

	var diags diag.Diagnostics
	domain := d.Get("domain").(string)


	s := strings.Split(d.Id(), "/")

	id := s[1]
	if (domain == ""){
		domain = s[0]
	}
	if (domain != s[0]){
		return diag.FromErr(errors.New("Mismatched domain/record id"))
	}

	res, err := c.GetRecord(domain, id)

	if err != nil {
		return diag.FromErr(err)
	}

	rc := res.Response.(gc.DNSResult).Records
	for _, rec := range rc {
		if rec["record_id"] == id {
			d.Set("name", rec["name"])
			d.Set("type", rec["type"])
			d.Set("proxied", rec["cloud"])
			d.Set("ttl", rec["ttl"])
			d.Set("content", rec["content"])
			d.Set("type", rec["type"])
			d.Set("record_id", id)
			d.Set("domain", domain)
			if rec["port"] != nil {
				d.Set("port", rec["port"])
			}
			if rec["weight"] != nil {
				d.Set("weight", rec["weight"])
			}
			if rec["priority"] != nil {
				d.Set("priority", rec["priority"])
			}
			return diags
		}
	}
	return diag.FromErr(errors.New(fmt.Sprintf("No `record_id` %s found", id)))	
}

func resourceRecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	c := m.(*gc.API)

	var diags diag.Diagnostics
	domain := d.Get("domain").(string)

	recordInfo := make(map[string]interface{})
	recordInfo["name"] = d.Get("name").(string)
	recordInfo["type"] = d.Get("type").(string)

	fmt.Printf("Gote: %v\n", d.Get("proxied"))

	if d.Get("proxied").(bool) {
		recordInfo["cloud"] = "1"
	} else {
		recordInfo["cloud"] = "0"
	}
	recordInfo["content"] = d.Get("content").(string)
	recordInfo["ttl"] = d.Get("ttl").(int)

	if d.Get("port") != nil {
		recordInfo["port"] = d.Get("port").(int)
	}
	if d.Get("weight") != nil {
		recordInfo["weight"] = d.Get("weight").(int)
	}
	if d.Get("priority") != nil {
		recordInfo["priority"] = d.Get("priority").(int)
	}

	resp, err := c.CreateRecord(domain, recordInfo)

	if err != nil {
		return diag.FromErr(err)
	}

	rciq := resp.Response.(gc.DNSResult).Records[0]



	d.SetId(fmt.Sprintf("%v/%v", domain, rciq["record_id"]))
	d.Set("record_id", fmt.Sprintf("%v", rciq["record_id"]))

	d.Set("name", rciq["name"])
	if rciq["cloud"].(string) == "1" {
		d.Set("proxied", true)
	} else {
		d.Set("proxied", false)
	}
	d.Set("content", rciq["content"])
	d.Set("ttl", rciq["ttl"])
	d.Set("type", rciq["type"])
	d.Set("last_updated", time.Now().Format(time.RFC850))

	//if err := ; err != nil {
	return diags
}

func resourceRecordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*gc.API)

	if d.HasChange("name") {
		return diag.FromErr(errors.New("Cannot update field 'name', you need to delete and create in different steps"))
	}

	if d.HasChange("type") {
		return diag.FromErr(errors.New("Cannot update field 'type', you need to delete and create in different steps"))
	}

	domain := d.Get("domain").(string)
	recordInfo := make(map[string]interface{})
	recordInfo["name"] = d.Get("name").(string)
	recordInfo["type"] = d.Get("type").(string)
	recordInfo["weight"] = d.Get("weight").(int)
	recordInfo["priority"] = d.Get("priority").(int)
	recordInfo["port"] = d.Get("port").(int)
	recordInfo["record_id"] = d.Get("record_id").(string)
	if d.Get("proxied").(bool) {
		recordInfo["cloud"] = "1"
	} else {
		recordInfo["cloud"] = "0"
	}
	recordInfo["content"] = d.Get("content").(string)
	recordInfo["ttl"] = d.Get("ttl").(int)

	_, err := c.UpdateRecord(domain, recordInfo)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))
	return diags
}

func resourceRecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*gc.API)

	s := strings.Split(d.Id(), "/")

	_, err := c.DeleteRecord(s[0], s[1])

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
