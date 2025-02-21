
# AWS

- service: a software that provides functionality and performs a task or set of tasks for your system
  - examples:
    - apache
    - nginx
    - postgresql
    - auth0
    - prometheus

- objective: facilitate communication between different systems
  - linux, windows, MacOS

- serverless: servers aren't directly managed
  - e.g. AWS Lambda
- private cloud: cloud infrastructure that is operated solely for a single organization
- public cloud: cloud infrastructure that is available to the general public

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

- advantages of cloud computing
  - On-demand resources can be provisioned without human interaction
  - resources are available over the network, and can be accessed by diverse client platforms
  - quickly and easily scale based on demand (stop guessing capacity)
  - measured payment for resources (pay only what you use)
  - benefits from massive economies of scale
  - trade CAPEX (CAPital EXpense) for OPEX (OPerational EXpense)
  - easily access global infrastructure

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

- load balancing: distributing incoming network traffic across multiple servers
  - improves responsiveness and availability of applications
  - prevents a single server from becoming a single point of failure

- on-site: infrastructure is physically located at your site
- IaaS (Infrastructure as a Service): provides virtualized computing resources over the internet
  - e.g. AWS EC2, Google Compute Engine
- PaaS (Platform as a Service): provides a platform allowing customers to develop, run, and manage applications without the complexity of building and maintaining the infrastructure
  - e.g. AWS Elastic Beanstalk
- SaaS (Software as a Service): software that is centrally hosted and licensed on a subscription basis
  - e.g. Google Drive

|                         | On-site      | IaaS           | PaaS           | SaaS           |
|-------------------------|--------------|----------------|----------------|----------------|
| **Applications**   | You | You | You | Service provider |
| **Data**           | You | You | You | Service provider |
| **Runtime**        | You | You | Service provider | Service provider |
| **Middleware**     | You | You | Service provider | Service provider |
| **O/S**            | You | You | Service provider | Service provider |
| **Virtualization** | You | Service provider | Service provider | Service provider |
| **Servers**        | You | Service provider | Service provider | Service provider |
| **Storage**        | You | Service provider | Service provider | Service provider |
| **Networking**     | You | Service provider | Service provider | Service provider |

- workloads: computing resources required to run an application or service in a cloud computing environment

## WAF (Well-Architected Framework) and CAF (Cloud Adoption Framework)

- WAF: a set of best practices for designing and operating reliable, secure, efficient, and cost-effective systems in the cloud
- CAF: a set of best practices for adopting cloud technologies and services in a structured and organized way

### 5 pillars of cloud computing architecture

1. operational excellence: automation, monitoring, and management
2. security: protect data, systems, and assets
3. reliability: ability to recover from failures and maintain availability
4. performance efficiency: use resources efficiently
5. cost optimization: avoid unnecessary costs

## storage services

- file storage: data is stored in a hierarchical file system (directories and files, like your computer)
  - e.g. AWS EFS
- block storage: data is stored in fixed-size blocks
  - e.g. AWS EBS
  - ideal for databases
- object storage: data is stored as objects
  - e.g. AWS S3
  - object: a collection of data that contains both data and metadata
  - ideal for non-structured data

## AWS CLI

> manage aws resources with CLI commands

- `aws iam list-users`: Lists all IAM users in your AWS account, showing each user's name, path, ID, ARN, and creation date

## IAM (Identity and Access Management)

> manage access to AWS services and resources securely

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

> compute service that allows you to launch virtual servers in the cloud

## Lambda

> serverless service that lets you run code without provisioning or managing servers

## S3 - Simple Storage Service

> scalable object storage service that allows you to store and retrieve data from anywhere on the web

## EBS - Elastic Block Store

> block storage service that allows you to create and attach storage volumes to EC2 instances

## VPC - Virtual Private Cloud

> virtual network that allows you to launch AWS resources in a logically isolated section of the cloud

## RDS - Relational Database Service

> managed database service that allows you to set up, operate, and scale relational databases in the cloud

## DynamoDB

> fully managed NoSQL database service that provides fast and predictable performance with seamless scalability

## ElastiCache

> fully managed in-memory caching service that makes it easy to deploy, operate, and scale popular open-source compatible in-memory data stores

## Redshift

> fully managed data warehouse service that allows you to run complex queries on large datasets

## Route 53

> scalable domain name system (DNS) web service that translates domain names into IP addresses

## CloudFront

> content delivery network (CDN) service that delivers data, videos, applications and APIs to customers globally with low latency and high transfer speeds

## KMS (Key Management Service)

> managed service that allows you to create and control the encryption keys used to encrypt your data

## API Gateway

> fully managed service that makes it easy for developers to create, publish, maintain, monitor, and secure APIs at any scale

## SNS - Simple Notification Service

> fully managed messaging service that allows you to send notifications to a large number of recipients

## SQS - Simple Queue Service

> fully managed message queuing service that allows you to decouple and scale microservices, distributed systems, and serverless applications

## Step Functions

> fully managed service that allows you to coordinate multiple AWS services into serverless workflows

## CodeCommit

> fully managed source control service that makes it easy for teams to host secure and highly scalable private Git repositories

## CodePipeline

> fully managed continuous integration and continuous delivery (CI/CD) service that automates the build, test, and deploy phases of your release process

## CodeBuild

> fully managed build service that compiles source code, runs tests, and produces software packages that are ready to deploy

## CodeDeploy

> fully managed deployment service that automates software deployments to a variety of compute services such as EC2, Lambda, and on-premises servers

## ECS (Elastic Container Service)

> fully managed container orchestration service that allows you to run, stop, and manage Docker containers on a cluster of EC2 instances

## ECR - Elastic Container Registry

> fully managed Docker container registry that makes it easy to store, manage, and deploy Docker container images

## EKR - Elastic Kubernetes Service

> fully managed Kubernetes service that allows you to run, manage, and scale containerized applications using Kubernetes

## CloudWatch

> monitoring and observability service that provides data and insights

- monitor your applications
- respond to system-wide performance changes
- optimize resource utilization
- get a unified view of operational health

## CloudFormation

> fully managed service that allows you to create and manage AWS resources using templates

## AMI - Amazon Machine Image

> a template that contains a software configuration (e.g. operating system, application server, applications) that is used to launch EC2 instances

## Elastic Beanstalk

> fully managed service that allows you to deploy and manage web applications and services

## Kinesis

> fully managed service that allows you to collect, process, and analyze real-time streaming data

## Cognito

> fully managed identity service that allows you to add user sign-up, sign-in, and access control to your web and mobile apps quickly and easily

## Glacier

> low-cost storage service for data archiving and long-term backup

## Lightsail

> deploy and manage applications/websites with pre-configured cloud resources

## Redshift

> fully managed data warehouse service that allows you to run complex queries on large datasets
## Athena

> interactive query service that allows you to analyze data in Amazon S3 using standard SQL

## EMR - Elastic MapReduce

> big data processing service that allows you to run Apache Hadoop, Spark, and other big data frameworks on AWS

---

- Read replica: is a copy of a database that can be used to offload read operations from the primary database
- server provisioning: the process of setting up physical or virtual hardware; installing and configuring software, such as the operating system and applications; and connecting it to middleware, network, and storage components
- failover: ability of a system or service to automatically switch to a backup or secondary system when the primary system becomes unavailable or experiences a failure. This is important for ensuring high availability and minimizing downtime for critical applications and services.
- VPC (Virtual Private Cloud): virtual network infrastructure that allows users to provision a logically isolated section where they can launch AWS resources in a virtual network
- topology: the arrangement of nodes or devices in a network
  - e.g. star, ring, bus, mesh, tree
- cluster: a group of interconnected computers or servers that work together to perform a specific task or function
  - commonly used in high-performance computing, data processing, and distributed systems to improve performance, scalability, and fault tolerance
  - resources are shared and coordinated to achieve a common goal
- grid: distributed computing system that connects multiple computers or servers to work together on a common task or problem
  - decentralized
