package gcp

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// NewProvider creates a new provider
func NewProvider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"google_app_engine_ssl_certificate": NewAppEngineSSLCertificate(),
		},
	}
}
