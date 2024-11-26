
# AWS

service: a software that provides functionality and performs a task or set of tasks for your system
  - examples:
    - apache
    - nginx
    - postgresql
    - auth0
    - prometheus

- objective: facilitate communication between different systems
  - linux, windows, MacOS

- serverless: servers aren't directly managed
  - S3
  - DynamoDB
  - Fargate
  - Lambda
- OLAP
  - used for platforms with intense read operations
  - used for analyzing large quantities of data from multiple perspectives or dimension
- OLTP
  - used for transactional processing of data in real-time
  - used for processing business transactions such as order processing, inventory management, etc

- Compute: CPU
- Memory: RAM
- Storage: Data
- Database: Storing data in a structured way
- Network: Routers, switches, DNS servers

- Advantages
  - On-demand resources can be provisioned without human interaction
  - Resources are available over the network, and can be accessed by diverse client platforms
  - Quickly and easily scale based on demand (stop guessing capacity)
  - Measured payment for resources (pay only what you use)
  - Benefits from massive economies of scale
  - Trade CAPEX (Capital Expense) for OPEX (Operational Expense)
  - Easily access global infrastructure

- problems solved by the cloud
  - flexibility
  - cost-effectiveness
  - scalability
  - elasticity
  - high-availability
  - fault-tolerance
  - agility

- pricing of AWS
  - pay for compute time
  - pay for data stored in the cloud
  - pay only for data transfer out of the cloud
    - data transfer in is free

- users
  - multiple people can belong to a group
  - a person can belong to multiple groups simultaneously

## AWS CLI

- `aws iam list-users`: Lists all IAM users in your AWS account, showing each user's name, path, ID, ARN, and creation date

## IAM (Identity and Access Management)

### Policy structure

```json
{
  "Version": "2012-10-17",
  "Id": "Policy12345",
  "Statement": [
    {
      "Sid": "AllowListBuckets",
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::example-bucket"
    },
    {
      "Sid": "AllowGetObjects",
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::example-bucket/*"
    }
  ]
}
```

- Version: the version of the policy language, defined by date
- ID: Policy ID
- Statement: composed by:
  - SID: Statement ID
  - Effect (Allow, Deny)
  - Principal: Account/user/role to which this policy refers to
  - Action: List of actions the policy allows or denies
  - Resource: List of resources to which the actions apply
  - Condition: Conditions for when this policy is in effect

- IAM Password Policy: You can decide the requirements for password creation

- MFA (Multi Factor Authentication): multi-step account login process that requires users to enter more information than just a password
  - e.g. password && (code sent to email || secret question || fingerprint scan)

- Options to access AWS:
  - AWS Management Console
  - AWS Command Line Interface (CLI)
  - AWS Software Developer Kit (SDK)
    - Used whenever you want to call APIs

- IAM Roles: Allows AWS services to perform tasks by granting credentials

### IAM Security Tools

- IAM Credentials Report (account-level)
- IAM Access Advisor (user-level)
  - Use this information to revise your policies (Least privilege principle)

### IAM Guidelines and Best Practices

- Don’t use the root account (except for AWS account setup)
- Don’t create multiple AWS accounts, create AWS users within an AWS account
- Assign users to groups and assign permissions to groups
  - If you manage company resources, create a group with the policy "Administrator Access" and create an AWS user that belongs to this group
- Create strong password policies and enforce Multi-Factor Authentication (MFA)
- Create and use roles to give permissions to AWS services
- Use Access Keys for Programmatic Access (CLI/SDK)
- Audit permissions by consulting:
  - IAM Credential Report
  - IAM Access Advisor
- Never share IAM users and Access Keys

- Shared Responsibility for IAM
  - AWS:
    - Infrastructure (global network security)
    - Configuration and vulnerability analysis
    - Compliance validation
  - Me:
    - Users, groups, roles, policies, management and monitoring
    - Enable MFA on all accounts
    - Rotate all keys often
    - Use IAM tools to apply appropriate permissions



## EC2 (Elastic Compute Cloud)

When you stop and restart an EC2 instance, the public IPv4 address can change, but the private address won’t

- sizing and configuration options:
  - OS
  - CPU
  - RAM
  - storage space
    - Network-attached (EBS and EFS)
    - Hardware (EC2 Instance Store)
  - Network Card
  - Firewall Rules: Security group
  - EC2 User Data: Bootstrap script launched at the first start of an instance

### Purchasing Options

- On-demand Instances:
  - Short workload, predictable pricing, pay by second (for Linux or Windows)
  - Highest cost but no upfront payment
  - No long-term commitment
  - Recommended for short-term and un-interrupted workloads, where you can’t predict application behavior
- Reserved Instances (1 and 3 years):
  - Long workloads, up to 72% discount compared to on-demand
  - Scope: Regional or zonal (reserve capacity in an AZ)
  - Recommended for steady-state usage applications (e.g., databases)
  - Can be bought and sold in the Reserved Instance Marketplace
  - Convertible Reserved Instance: Can change EC2 instance type, instance family, OS, scope, and tenancy. Up to 66% discount (less discount, more flexibility).
- Savings Plans:
  - Long workload, commitment to an amount of usage ($10/hour for 1 or 3 years)
  - Discount based on long-term usage (up to 72%)
  - Locks you into a specific instance family and AWS region (e.g., M% in us-east-1)
  - Flexible for instance size, OS, and tenancy
- Spot Instances:
  - Short workloads, cheap (up to 90% discount compared to on-demand)
  - Most cost-efficient instances on AWS
  - You can lose these instances (less reliable) if your max price is lower than the current spot price
  - Recommended for workloads resilient to failure (e.g., batch jobs, data analysis, image processing)
- Dedicated Hosts:
  - Reserve an entire physical server
  - Most expensive
  - Control instance placement
  - Purchasing options: On-demand, reserved (1 or 3 years)
  - Recommended for software with complicated licensing models or companies with strong regulatory or compliance needs
- Dedicated Instances:
  - Instances run on hardware dedicated to you. No other customers will share your hardware
  - No control over instance placement
- Capacity Reservations:
  - Reserve capacity in a specific Availability Zone (AZ) for any duration
  - Normally used to ensure you have the capacity to run workloads even during times of high demand
  - No time commitment (create/cancel anytime)
  - No Billing Discounts
  - Charges at on-demand rates whether you run instances or not
  - Recommended for short-term uninterrupted workloads that need to be in a specific AZ

### EC2 Instance Tenancy

- Shared (default): Multiple AWS accounts can share the same physical hardware
- Dedicated Instance: Your instance runs on single-tenant hardware
- Dedicated Host: Your instance runs on a physical server fully dedicated to your use
  - An isolated server with configurations you can control

## Security Groups

> Facilitates managing network traffic

- Acts as a "firewall" for EC2 instances
- Regulate:
  - Access to ports
  - Authorized IP ranges (IPv4 and IPv6)
  - Control of inbound and outbound network traffic
- Can be attached to multiple instances
- Locked down to a region/VPC combination
- Lives "outside" the EC2
  - If traffic is blocked, the EC2 instance won’t see it
- It's good to keep 1 separate security group for SSH access
- If your application is not accessible (time out), then it's a security group issue
- If you application gives a "connection refused" error, then it's an application error or it's not launched

> [!IMPORTANT]
> All inbound traffic is blocked by default
> All outbound traffic is authorized by default

- Ports:
  - 22 = SSH (Secure Shell): log into a Linux instance
  - 3389 = RDP (Remote Desktop Protocol): log into a Windows instance
  - 21 = FTP (File Transfer Protocol)
  - 22 = SFTP (Secure File Transfer Protocol)
  - 80 = HTTP : access unsecured websites
  - 443 = HTTPS: access secured websites

### SSH Access

- Create an identity file `.pem` for SSH:
  - EC2 Dashboard > Key Pair > Create Key Pair
  - Copy the public IP of the instance (used in the SSH command)
  - Move the `.pem` file to a secure location and restrict permissions:
    - `chmod 0400 <file>.pem`
  - SSH into the instance:
    - `ssh -i <file>.pem ec2-user@<public_IP>`

## EC2 - Elastic Compute Cloud
## S3 - Simple Storage Service
## RDS - Relational Database Service
## Lambda
## API Gateway
## CloudFront
## Route 53
## CloudWatch
## SNS - Simple Notification Service
## SQS - Simple Queue Service
## DynamoDB
## Elastic Beanstalk
## CloudFormation
## Kinesis
## Cognito
## CodePipeline
## CodeBuild
## CodeDeploy
## CodeCommit

---

- Read replica: is a copy of a database that can be used to offload read operations from the primary database
- server provisioning: the process of setting up physical or virtual hardware; installing and configuring software, such as the operating system and applications; and connecting it to middleware, network, and storage components
- failover: ability of a system or service to automatically switch to a backup or secondary system when the primary system becomes unavailable or experiences a failure. This is important for ensuring high availability and minimizing downtime for critical applications and services.
- VPC (Virtual Private Cloud): virtual network infrastructure that allows users to provision a logically isolated section where they can launch AWS resources in a virtual network
