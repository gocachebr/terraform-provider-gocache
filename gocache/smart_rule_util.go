package gocache

import (
	gc "github.com/gocachebr/gocache-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"context"
	"fmt"
	"strings"
	"encoding/json"
)

func processRuleRequest(d *schema.ResourceData, i int, fieldType string, ruleType string) map[string]interface{} {
	changes := make(map[string]interface{})
	resource := fieldType
	if ruleType != "" {
		resource = fmt.Sprintf("%s_%s",ruleType,fieldType)
	}
	for key := range gc.GetAllFieldsAdjusted(fmt.Sprintf("smart_rule_%s",resource)) {
		var val interface{}
		var ok bool
		val,ok = d.GetOk(fmt.Sprintf("smart_rule.%d.%s.0.%s",i,fieldType,key))


		if ok {
			switch key {
			case "set_response_headers":
				json, err := json.Marshal(val)
				if err == nil {
					val = string(json)
				}
			case "set_request_headers":
				headerList := make([]string,0)
				headers := val.(map[string]interface{})
				for k,v := range headers {
	
					k = strings.Replace(strings.Replace(k,"\\","\\\\",-1),"\"","\\\"",-1)
					v = strings.Replace(strings.Replace(v.(string),"\\","\\\\",-1),"\"","\\\"",-1)
	
					headerList = append(headerList,fmt.Sprintf("%s:%s",k,v))
	
				} 
				val = headerList
			}

			newKey := gc.GetFieldName(key,fmt.Sprintf("smart_rule_%s",resource))

			if val != nil {
				changes[newKey] = val
			}
		}
	}
	return changes
}

func convertRulesToMapArray(rules []gc.SmartRule, ruleType string, d *schema.ResourceData) []map[string][]map[string]interface{}{
	var result []map[string][]map[string]interface{} 
	for i, r := range rules {
		rule := make(map[string][]map[string]interface{}) 
		rule["match"] = processRuleResponse(r.Match,"match","",i,d)
		rule["action"] = processRuleResponse(r.Action,"action",ruleType,i,d)
		metadata := processRuleResponse(r.Metadata,"metadata","",i,d)
		if len(metadata[0]) > 0 {
			rule["metadata"] = metadata
		}
		result = append(result,rule)
	}
	return result
}

func processRuleResponse(values map[string]interface{}, fieldType string, ruleType string, ruleIndex int, d *schema.ResourceData) []map[string]interface{} {
	changes := make([]map[string]interface{},1)
	changes[0] = make(map[string]interface{})
	resource := fieldType
	if ruleType != "" {
		resource = fmt.Sprintf("%s_%s",ruleType,fieldType)
	}
	resource = fmt.Sprintf("smart_rule_%s",resource)
	for key,val := range values {
		// Avoid reading status if is true and not defined in config
		if key == "status" && fieldType == "metadata"{
			_,ok := d.GetOk(fmt.Sprintf("smart_rule.%d.%s.0.%s",ruleIndex,fieldType,key))
			if !ok && val == "true" {
				continue
			} 
		}

		newKey := gc.GetFieldName(key,fmt.Sprintf("%v_reversed",resource))

		switch newKey {
		case "set_response_headers":
			var valmap map[string]interface{}
			err := json.Unmarshal([]byte(val.(string)),&valmap)
			if err == nil {
				val = valmap
			}
		case "set_request_headers":
			valmap := make(map[string]interface{})
			for _,h := range val.([]interface{}) {

				split := strings.Split(h.(string), ":")

				valmap[split[0]] = split[1]

			} 
			val = valmap
		}

		if gc.FieldExists(newKey,resource){
			changes[0][newKey] = val
		}
	}
	return changes
}

func readRulesConfig(d *schema.ResourceData, rules []interface{}, ruleType string) ([]gc.SmartRule, error){
	newRules := make([]gc.SmartRule,0)

	for i, r := range rules {
		rule := r.(map[string]interface{})
		match := rule["match"].([]interface{})
		action := rule["action"].([]interface{})
		metadata := rule["metadata"].([]interface{})

		if len(match) > 1 {
			return nil,fmt.Errorf("must have only one match resource in rule %d",i)
		}

		if len(action) > 1 {
			return nil,fmt.Errorf("must have only one action resource in rule %d",i)
		}

		if len(metadata) > 1 {
			return nil,fmt.Errorf("must have only one action resource in rule %d",i)
		}

		newRule := gc.SmartRule{}

		newRule.Match = processRuleRequest(d,i,"match","")
		newRule.Action = processRuleRequest(d,i,"action",ruleType)
		newRule.Metadata = processRuleRequest(d,i,"metadata","")

		newRules = append(newRules,newRule)
	}
	return newRules, nil
}

func getSchema(resource string) map[string]*schema.Schema {	
	newSchema := map[string]*schema.Schema{}

	for ind, content := range gc.GetAllFieldsAdjusted(resource) {
		newSchema[ind] = &schema.Schema{
			Type: schema.TypeString,
		}

		if resource == "smart_rule_rewrite_action" || resource == "smart_rule_firewall_action" {
			newSchema[ind].Required = true
		}else {
			newSchema[ind].Optional = true
		}

		if content.Mode == "list" {
			newSchema[ind].Type = schema.TypeList
			newSchema[ind].Elem = &schema.Schema{
				Type: schema.TypeString,
			}
		}else if content.Mode == "map"{
			newSchema[ind].Type = schema.TypeMap
			newSchema[ind].Elem = &schema.Schema{
				Type: schema.TypeString,
			}
		}else{
			newSchema[ind].StateFunc = func(val interface{}) string {
				return fmt.Sprintf("%v", val)
			}
		}
	}
	return newSchema
}

func createSmartRulesBulk(ctx context.Context, d *schema.ResourceData, m interface{}, ruleType string) diag.Diagnostics {
	c := m.(*gc.API)

	var diags diag.Diagnostics

	domain := d.Get("domain").(string)

	d.SetId(domain)

	rules := d.Get("smart_rule").([]interface{})

	newRules, err := readRulesConfig(d, rules, ruleType)

	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := c.CreateRules(domain, newRules, ruleType)

	if err != nil {
		return diag.FromErr(err)
	}

	var id int

	if resp.HTTPStatusCode == 200 {
		id = resp.Response.(gc.SmartRuleResult).Bulk_id
	} else {
		return diag.FromErr(fmt.Errorf("Status code was %d, msg: %s", resp.HTTPStatusCode, resp.Msg))
	}

	resp, err = c.ApplyBulk(domain, id, ruleType)

	if err != nil {
		return diag.FromErr(fmt.Errorf("%v %v", err, resp))
	}

	d.Set("bulk_id", id)

	return diags
}

func readSmartRulesBulk(ctx context.Context, d *schema.ResourceData, m interface{}, ruleType string) diag.Diagnostics {
	c := m.(*gc.API)

	var diags diag.Diagnostics

	domain := d.Id()
	// id := d.Get("bulk_id").(int)

	resp, err := c.GetRules(domain, ruleType)

	if err != nil {
		return diag.FromErr(err)
	}

	if resp.HTTPStatusCode == 200 {

		rules := resp.Response.(gc.SmartRuleResult).Rules

		convert := convertRulesToMapArray(rules, ruleType, d)

		err := d.Set("smart_rule", convert)

		// return diag.FromErr(fmt.Errorf("Convert:%v\n Rules: %v \n D: %v\n Resp: %v",convert, rules,d.Get("smart_rule"), resp))

		if err != nil {
			return diag.FromErr(err)
		}

	}

	return diags
}

func updateSmartRulesBulk(ctx context.Context, d *schema.ResourceData, m interface{}, ruleType string) diag.Diagnostics {
	c := m.(*gc.API)

	var diags diag.Diagnostics

	id, ok := d.GetOk("bulk_id")

	if !ok {
		return createSmartRulesBulk(ctx,d,m,ruleType)
	}

	domain := d.Get("domain").(string)

	d.SetId(domain)
	rules := d.Get("smart_rule").([]interface{})

	newRules, err := readRulesConfig(d, rules, ruleType)

	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := c.UpdateRules(domain, newRules, id.(int), ruleType)

	if err != nil {
		if resp.HTTPStatusCode == 404 {
			resp, err = c.CreateRules(domain, newRules, ruleType)
		} else {
			return diag.FromErr(err)
		}
	}

	if resp.HTTPStatusCode == 200 {
		id = resp.Response.(gc.SmartRuleResult).Bulk_id
	} else {
		return diag.FromErr(fmt.Errorf("Status code was %d, msg: %s, err: %v", resp.HTTPStatusCode, resp.Msg, err))
	}

	resp, err = c.ApplyBulk(domain, id.(int), ruleType)

	if err != nil {
		return diag.FromErr(fmt.Errorf("%v %v", err, resp))
	}

	d.Set("bulk_id", id)

	return diags
}

func deleteSmartRulesBulk(ctx context.Context, d *schema.ResourceData, m interface{}, ruleType string) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*gc.API)

	domain := d.Get("domain").(string)

	id := d.Get("bulk_id").(int)

	_, err := c.DeleteBulk(domain, id, ruleType)

	if err != nil {
		return diag.FromErr(err)
	}

	_, err = c.DeleteBulk(domain, -1, ruleType)

	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}