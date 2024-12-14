
# Terraform

> Infrastructure as Code (IaC) tool that lets you define cloud resources in human-readable configuration files

## Core Concepts

- Provider: plugin that allows Terraform to interact with cloud platforms (AWS, Azure, GCP)
- Resource: infrastructure object you want to manage (EC2 instance, S3 bucket)
- State: Terraform's record of all managed infrastructure
- Module: reusable Terraform configurations

## Basic Structure

```hcl
# Provider configuration
provider "aws" {
  region = "us-west-2"
}

# Resource example
resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
}
```

## setup steps

1. install terraform
2. Configure provider authentication:

```bash
export AWS_ACCESS_KEY_ID="<insert key>"
export AWS_SECRET_ACCESS_KEY="<insert key>"
```

3. create terraform files (`.tf` extension)

## main commands

- `terraform init`: initialize working directory, download providers
- `terraform plan`: preview changes before applying
- `terraform apply`: apply changes to create/modify resources
- `terraform destroy`: remove all resources managed by Terraform
- `terraform fmt`: format code per HCL standards
- `terraform validate`: validate configuration files

## best practices

- version control your terraform code
- use variables for reusable values
- organize code into modules
- always review plan before applying
- store state files securely
- use workspaces for different environments
