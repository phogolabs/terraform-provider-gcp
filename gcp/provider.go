package gcp

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// NewProvider creates a new provider
func NewProvider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"gcp_app_engine_ssl_certificate": NewAppEngineSSLCertificate(),
		},
	}
}
