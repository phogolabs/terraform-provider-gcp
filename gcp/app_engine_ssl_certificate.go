package gcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"google.golang.org/api/appengine/v1"
)

// NewAppEngineSSLCertificate returns the app_engine_ssl_certificate resource
func NewAppEngineSSLCertificate() *schema.Resource {
	read := func(d *schema.ResourceData, m interface{}) error {
		var (
			ctx   = context.Background()
			appID = d.Get("app_id").(string)
		)

		service, err := appengine.NewService(ctx)
		if err != nil {
			return err
		}

		record, err := service.Apps.AuthorizedCertificates.Get(appID, d.Id()).Do()
		if err != nil {
			return err
		}

		d.SetId(record.Id)

		if err := d.Set("name", record.DisplayName); err != nil {
			return err
		}

		if err := d.Set("private_key", record.CertificateRawData.PrivateKey); err != nil {
			return err
		}

		if err := d.Set("certificate", record.CertificateRawData.PublicCertificate); err != nil {
			return err
		}

		return nil
	}

	create := func(d *schema.ResourceData, m interface{}) error {
		var (
			ctx   = context.Background()
			appID = d.Get("app_id").(string)
		)

		service, err := appengine.NewService(ctx)
		if err != nil {
			return err
		}

		input := &appengine.AuthorizedCertificate{
			DisplayName: d.Get("name").(string),
			CertificateRawData: &appengine.CertificateRawData{
				PrivateKey:        d.Get("private_key").(string),
				PublicCertificate: d.Get("certificate").(string),
			},
		}

		record, err := service.Apps.AuthorizedCertificates.Create(appID, input).Do()
		if err != nil {
			return err
		}

		d.SetId(record.Id)

		if err := d.Set("name", record.DisplayName); err != nil {
			return err
		}

		if err := d.Set("private_key", record.CertificateRawData.PrivateKey); err != nil {
			return err
		}

		if err := d.Set("certificate", record.CertificateRawData.PublicCertificate); err != nil {
			return err
		}

		return nil
	}

	update := func(d *schema.ResourceData, m interface{}) error {
		var (
			ctx   = context.Background()
			appID = d.Get("app_id").(string)
		)

		service, err := appengine.NewService(ctx)
		if err != nil {
			return err
		}

		input := &appengine.AuthorizedCertificate{
			DisplayName: d.Get("name").(string),
			CertificateRawData: &appengine.CertificateRawData{
				PrivateKey:        d.Get("private_key").(string),
				PublicCertificate: d.Get("certificate").(string),
			},
		}

		record, err := service.Apps.AuthorizedCertificates.
			Patch(appID, d.Id(), input).
			UpdateMask("display_name,certificate_raw_data").
			Do()
		if err != nil {
			return err
		}

		d.SetId(record.Id)

		if err := d.Set("name", record.DisplayName); err != nil {
			return err
		}

		if err := d.Set("private_key", record.CertificateRawData.PrivateKey); err != nil {
			return err
		}

		if err := d.Set("certificate", record.CertificateRawData.PublicCertificate); err != nil {
			return err
		}

		return nil
	}

	destroy := func(d *schema.ResourceData, m interface{}) error {
		var (
			ctx   = context.Background()
			appID = d.Get("app_id").(string)
		)

		service, err := appengine.NewService(ctx)
		if err != nil {
			return err
		}

		_, err = service.Apps.AuthorizedCertificates.Delete(appID, d.Id()).Do()
		return err
	}

	return &schema.Resource{
		Create: create,
		Update: update,
		Read:   read,
		Delete: destroy,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:        schema.TypeString,
				Description: "Application Identitifier",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name",
				Required:    true,
			},
			"private_key": {
				Type:        schema.TypeString,
				Description: "Private key",
				Required:    true,
			},
			"certificate": {
				Type:        schema.TypeString,
				Description: "Certificate",
				Required:    true,
			},
		},
	}
}
