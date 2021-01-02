# terraform-provider-gcp

Google Cloud Provider Extension

## Getting Started

```hcl
provider "gcp" {}

resource "gcp_app_engine_ssl_certificate" {
  app_id      = "application unique identitifer"
  name        = "my certificate"
  private_key = ""
  certificate = ""
}
```
