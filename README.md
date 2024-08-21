# CDNVideo Terraform Provider

![CDNVideo](https://dashboard.cdnvideo.ru/app/inventory/v1/image/425e95ca1cdb404ca0e0b948cf734b35)

The CDNvideo Terraform Provider enables configuration and management of CDN resources using HashiCorp Terraform. This provider simplifies the integration and automation of CDN resource management with Terraform's declarative language.

**Latest provider version:**

- [CDNVideo Provider on Terraform Registry](https://registry.terraform.io/providers/opensource-cdnvideo/cdnvideo/latest)

## Quickstarts

### Minimum Requirements

- [Terraform](https://www.terraform.io/downloads.html) version 1.8 or newer

It's recommended to use the [latest version of Terraform](https://developer.hashicorp.com/terraform/install?product_intent=terraform) for best compatibility with the CDNvideo provider. Versions earlier than 1.2 may have compatibility issues with newer features.

### Installation

1. **Install Terraform**

   Ensure that you have the latest version of Terraform. Download it from the [official Terraform website](https://developer.hashicorp.com/terraform/install?product_intent=terraform).

2. **Configure Terraform if registry not available (Optional)**

   To use Terraform without a VPN, configure a mirror in your `~/.terraformrc` file:

   ```shell
   $ touch ~/.terraformrc
   ```
   
   ```hcl
   provider_installation {
     network_mirror {
       url = "https://terraform-mirror.yandexcloud.net/"
       include = ["registry.terraform.io/*/*"]
     }
     direct {
       exclude = ["registry.terraform.io/*/*"]
     }
   }
   ```

3. **Install the Provider**

   To use the CDNvideo Terraform Provider, create a `main.tf` file with the following configuration:

   ```terraform
   terraform {
     required_version = ">= 1.8"

     required_providers {
       cdnvideo = {
         source  = "opensource-cdnvideo/cdnvideo"
         version = "1.0.1"   # specify the version (select from https://github.com/opensource-cdnvideo/terraform-provider-cdnvideo/releases)
       }
     }
   }

   provider "cdnvideo" {
     account_name = "account_name"
     username     = "example@example.ru"
     password     = "pass"
   }
   ```

   Then, run:

   ```shell
   terraform init
   ```

   This initializes your Terraform configuration directory and downloads the provider. It should be the first command run after creating a new Terraform configuration or cloning an existing one from version control. It is safe to run this command multiple times.

### Writing Module Files

Refer to the examples in the `./examples` folder to create your module files.

