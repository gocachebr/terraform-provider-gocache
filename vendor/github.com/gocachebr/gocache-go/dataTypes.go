package gocache

import (
	"fmt"
	"strconv"
	"strings"
)

type API_Response struct {
	Status_code    int         `json:"status_code,omitempty"`
	Response       interface{} `json:"response,omitempty"`
	HTTPStatusCode int         `json:"-"`
}

type DomainList struct {
	Domains     []string    `json:"domains"`
	Nameservers interface{} `json:"nameservers,omitempty"`
}

type DNSResult struct {
	Records []map[string]interface{} `json:"records"`
}

type CertificateResult struct {
	Msg          string                   `json:"msg,omitempty"`
	Certificates []map[string]interface{} `json:"certificates,omitempty"`
}

type CertificateInput struct {
	Certificate string `json:"certificate,omitempty"`
	Privatekey  string `json:"privatekey,omitempty"`
	Type        string `json:"type,omitempty"`
	Enable      bool   `json:"enable"`
}

type fieldType struct {
	Mode    string
	Allowed []string
}

var domainFields = map[string]fieldType{
	"cache_ttl":                        fieldType{Mode: "int", Allowed: []string{"300", "600", "900", "1800", "3600", "7200", "14400", "86400", "172800", "604800"}},
	"deploy_mode":                      fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"smart_status":                     fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"smart_tpl":                        fieldType{Mode: "string", Allowed: []string{"custom", "wordpress", "magento", "joomla"}},
	"smart_ttl":                        fieldType{Mode: "int", Allowed: []string{"300", "600", "900", "1800", "3600", "7200", "14400", "86400", "172800", "604800"}},
	"cdn_mode":                         fieldType{Mode: "string", Allowed: []string{"ns", "cname"}},
	"gzip_status":                      fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"expires_ttl":                      fieldType{Mode: "int", Allowed: []string{"-1", "3600", "7200", "14400", "43200", "86400", "172800", "345600", "604800", "1296000", "2592000", "15552000", "31536000"}},
	"ignore_vary":                      fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"ignore_cache_control":             fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"ignore_expires":                   fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"ssl_mode":                         fieldType{Mode: "string", Allowed: []string{"full", "partial"}},
	"cache_301":                        fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"cache_302":                        fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"cache_404":                        fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"header_device_type":               fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"header_geoip_continent":           fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"header_geoip_country":             fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"header_geoip_org":                 fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"caching_behavior":                 fieldType{Mode: "string", Allowed: []string{"default", "ignore_query_string"}},
	"waf_status":                       fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"waf_level":                        fieldType{Mode: "string", Allowed: []string{"low", "medium", "high"}},
	"waf_mode":                         fieldType{Mode: "string", Allowed: []string{"simulate", "challenge", "block"}},
	"expire_bypass_sec":                fieldType{Mode: "int", Allowed: []string{"-1", "300", "600", "900", "1800", "3600", "7200", "14400", "43200", "86400", "172800", "345600", "604800", "1296000", "2592000", "15552000", "31536000"}},
	"tls10":                            fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"tls11":                            fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"tls12":                            fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize":                   fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_webp":              fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_progressive":       fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_metadata":          fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"image_optimize_level":             fieldType{Mode: "int", Allowed: []string{"0-100"}},
	"rate_limit_status":                fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"rate_limit_ignore_known_bots":     fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"rate_limit_ignore_static_content": fieldType{Mode: "boolean", Allowed: []string{"true", "false"}},
	"public_log":                       fieldType{Mode: "string", Allowed: []string{"off", "v1", "v2"}},
	"log_freq":                         fieldType{Mode: "int", Allowed: []string{"1", "5", "10", "15", "30", "60"}},
}

var domainConvert = map[string]string{
	"smart_cache_status":          "smart_status",
	"smart_cache_template":        "smart_tpl",
	"smart_cache_ttl":             "smart_ttl",
	"dns_mode":                    "cdn_mode",
	"rate_limit_ignore_good_bots": "rate_limit_ignore_known_bots",
}
var domainConvertReversed = map[string]string{
	"smart_status":                 "smart_cache_status",
	"smart_tpl":                    "smart_cache_template",
	"smart_ttl":                    "smart_cache_ttl",
	"cdn_mode":                     "dns_mode",
	"rate_limit_ignore_known_bots": "rate_limit_ignore_good_bots",
}

var recordConvert = map[string]string{
	"proxied": "cloud",
}

// Convert api response map keys to terraform variable names
func responseConvert(data map[string]interface{}, fields map[string]string) map[string]interface{} {
	output := make(map[string]interface{})
	for key, field := range data {
		modes, found := domainFields[key]
		//If the key is changed~
		newKey := key
		value, ok := fields[key]
		if ok {
			newKey = value
		}

		if found {
			if modes.Mode == "boolean" {
				if field == "true" {
					output[newKey] = true
				} else {
					output[newKey] = false
				}
			} else if modes.Mode == "int" {
				i, err := strconv.Atoi(field.(string))
				if err == nil {
					output[newKey] = i
				} else {
					output[newKey] = 0
				}
			} else {
				output[newKey] = field
			}
		} else {
			output[newKey] = field
		}
	}
	return output
}

// Convert map to formData for request body
func formData(data map[string]interface{}, fields map[string]string) string {
	result := ""

	for key, value := range data {
		key = strings.ToLower(key)
		field, ok := fields[key]
		name := key
		if ok {
			name = field
		}
		result += fmt.Sprintf("%v=%v&", name, value)
	}

	if result[len(result)-1] == '&' {
		result = result[:len(result)-1]
	}
	return result
}

func GetAllDomainSettings() map[string]fieldType {
	return domainFields
}

func GetAllDomainSettingsAdjusted() map[string]fieldType {
	aux := GetAllDomainSettings()
	domainFields := make(map[string]fieldType)
	for key, value := range aux {
		alias, ok := domainConvertReversed[key]
		if ok {
			domainFields[alias] = value
		} else {
			domainFields[key] = value
		}
	}
	return domainFields
}

func inList(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func CheckDomainValue(field string, value string) bool {
	allowed, ok := GetAllDomainSettingsAdjusted()[field]
	if ok {
		if allowed.Allowed[0] == "0-100" {
			i, err := strconv.Atoi(field)
			if err != nil {
				return i >= 0 && i <= 100
			}
		} else if inList(value, allowed.Allowed) {
			return true
		}
	}
	return false
}

func IsDomainField(field string) bool {
	for k := range GetAllDomainSettingsAdjusted() {
		if field == k {
			return true
		}
	}
	return false
}
