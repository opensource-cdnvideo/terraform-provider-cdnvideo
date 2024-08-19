CDNVideo Terraform Provider
------------------------------

<img alt="CDNVideo" src="https://dashboard.cdnvideo.ru/app/inventory/v1/image/425e95ca1cdb404ca0e0b948cf734b35"/>

================================================================================

The CDNvideo Terraform Provider allows you to configure and manage CDN resources using HashiCorp Terraform. This provider offers a simple way to integrate and automate the management of CDN resources using the declarative Terraform language.

Latest provider version

- [CDNVideo Provider](https://registry.terraform.io/providers/opensource-cdnvideo/cdnvideo/latest)

---

Quickstarts
------------------
### Minimum requirements

- [Terraform](https://www.terraform.io/downloads.html) version 1.8 or newer

We recommend running the [latest version](https://developer.hashicorp.com/terraform/install?product_intent=terraform) for optimal compatibility with the CDNvideo provider. Terraform versions older than 1.2 have known issues with newer features and internals.

### Installation

1.	Install Terraform

Ensure that you have the latest version of Terraform installed. You can download it from the [official Terraform website](https://developer.hashicorp.com/terraform/install?product_intent=terraform).

2.	Install the Provider

To use the CDNvideo Terraform Provider, add the following block to your Terraform configuration file (.tf):

```terraform
terraform {
  required_version = ">= 1.8"

  required_providers {
    cdnvideo = {
       source = "opensource-cdnvideo/cdnvideo"
       version = "{version_number}"   # specify the version (select from https://github.com/opensource-cdnvideo/terraform-provider-cdnvideo/releases)
    }
  }
}
```

Then, run the following command:
```shell
terraform init
```

This will download and install the provider.

### Configuring Terraform

#### Create Terraform configuration file
```shell
$ touch ~/.terraformrc
```
To use Terraform without a VPN, configure a mirror in the `~/.terraformrc` file as follows:
```
provider_installation {

   network_mirror {
      url = "https://terraform-mirror.yandexcloud.net/"
      include = ["registry.terraform.io/*/*"]
   }
   # For all other providers, install them directly from their origin provider 
   # registries as normal. If you omit this, Terraform will _only_ use 
   # the dev_overrides block, and so no other providers will be available. 
   direct {
      exclude = ["registry.terraform.io/*/*"]
   }
}
```
### Project initializing

#### Set up the module configuration file
Follow the guidelines at https://developer.hashicorp.com/terraform/language/providers/requirements to add provider settings. Every Terraform module must specify the providers it needs for Terraform to install and use them.

Create a provider.tf file in your module directory with the following configuration:

```terraform
terraform {
  required_version = ">= 1.8"

  required_providers {
    cdnvideo = {
       source = opensource-cdnvideo/cdnvideo"
       version = "{version_number}"   # specify the version (select from https://github.com/opensource-cdnvideo/terraform-provider-cdnvideo/releases)
    }
  }
}

provider "cdnvideo" {
  account_name = "account_name"
  username     = "example@example.ru"
  password     = "pass"
}
```
#### Initialize working directory
Run terraform init in your module directory:
```shell
terraform init
```
This command sets up your working directory with Terraform configuration files. It should be the first command run after creating a new Terraform configuration or cloning an existing one from version control. It is safe to run this command multiple times.

### Writing modules files

Use the examples provided in the ./examples folder to create your module files.
