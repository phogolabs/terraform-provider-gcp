# terraform-provider-gcp

Google Cloud Provider Extension

## Getting Started

```hcl
provider "gcp" {}

resource "gcp_app_engine_ssl_certificate" {
  app_id      = "<APP ENGINE ID>"
  name        = "<DISPLAY NAME>"
  private_key = "<PRIVATE KEY>"
  certificate = "<CERTIFICATE>"
}
```
