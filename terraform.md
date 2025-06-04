
# terraform

> Infrastructure as Code (IaC) tool that lets you define cloud resources in human-readable configuration files

## core concepts

- provider: plugin that allows Terraform to interact with cloud platforms (AWS, Azure, GCP)
- resource: infrastructure object you want to manage (EC2 instance, S3 bucket)
- state: Terraform's record of all managed infrastructure
- module: directory with `.tf` files
  - enables reusability
  - 2 types:
    - root module: directory where `terraform init` and `terraform apply` is run, the entry point
    - child module: any module referenced from another modul

## basic structure

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
1. export provider credential variables: e.g. `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`
1. run `terraform init` in project's root directory
1. create terraform files (`.tf` extension)
1. run `terraform plan -var-file="terraform.tfvars" -out=tfplan`
1. run `terraform apply tfplan`

## main commands

- `terraform init`: initialize working directory, download providers
- `terraform plan`: preview changes before applying
  - `terraform plan -var-file="terraform.tfvars" -out=tfplan`
    - `-var-file="terraform.tfvars"`: loads variable values from a file named `terraform.tfvars`
    - `-out=tfplan`: seves the plan to a file named `tfplan`, apply it with `terraform apply tfplan`
- `terraform apply`: apply changes to create/modify resources
- `terraform destroy`: remove all resources managed by Terraform
- `terraform fmt`: format code per HCL standards
- `terraform validate`: validate configuration files
- `terraform output ec2_public_ip`: print terraform variables

```bash
terraform apply -var-file="terraform.tfvars"
```

you can destroy resources with:

```bash
terraform apply -destroy -target=aws_instance.app_instance -auto-approve
terraform apply -auto-approve
```

## best practices

- version control your terraform code
- use variables for reusable values
- organize code into modules
- always review plan before applying
- store state files securely
- use workspaces for different environments

## types of resources

- `aws_instance`: EC2 instance
- `aws_vpc`: Virtual Private Cloud
- `aws_subnet`: subnet within a VPC
- `aws_security_group`: security group for EC2 instances
- `aws_s3_bucket`: S3 bucket
- `aws_iam_role`: Identity and Access Management role

