
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

> manage aws resources with CLI commands

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

> comp

## S3 - Simple Storage Service
## RDS - Relational Database Service
## Lambda
## API Gateway
## CloudFront
## Route 53
## CloudWatch
## AMI - Amazon Machine Image
## ECR - Elastic Container Registry
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
