
# AWS

> cloud computing platform that provides a wide range of services and tools for building and managing applications and infrastructure in the cloud

- cloud computing: the delivery of computing services over the internet
- serverless: server doesn't requires provisioning and scaling
  - e.g. aws lambda, azure functions, google cloud functions
- service: a software that provides functionality and performs a task or set of tasks for your system
  - e.g. apache, nginx, postgresql, auth0
- infrastructure: the physical or virtual resources that support the operation of a system
  - e.g. servers, storage, network, databases, etc

- types of pricing in AWS
  - pay for compute time
  - pay for data stored in the cloud
  - pay only for data transfer out of the cloud
    - data transfer in is free

- options to manage aws
  - aws management console
  - AWS Command Line Interface (CLI)
  - AWS Software Developer Kit (SDK): allows the access and management of aws services programmatically
    - Used whenever you want to call APIs

- reliability: ability to recover from failures and maintain availability
  - dynamically acquire computing resources to meet demand
- high availability: design and implementation of systems that are resilient to failures and can continue operating with minimal downtime
  - running instances for the same application on at least 2 Availability Zones
  - key characteristics
    - redundancy
    - failover mechanisms
    - geographic distribution
    - health monitoring
    - scalability and elasticity

- aws global infrastructure
  - region: separate geographic areas (e.g. `us-east-1` or `sa-east-1`)
    - each region is completely independent from others
    - each region has 3 to 6 Availability Zones, with few exceptions (AZ ⊂ region)
  - Availability Zone (AZ): isolated locations within each region
  - edge networks (Points of Presence): for content delivery as close as possible to users

- aws global services: services that are NOT tied to a specific region and operate across all regions
  - e.g. IAM, route 53, cloudfront, WAF, shield, aws organizations, aws artifact
  - TIP: if a service manages access, identity or DNS for your entire aws environment, it's likely global

- root user: the original account owner with full administrative access to all AWS services and resources
- IAM users: individual
- IAM groups: collections of IAM users
- AWS organizations: allows management of multiple AWS accounts under one umbrella
  - allows management of permissions
  - AWS organizations Service Control Policies (SCPs): centrally manage and restrict permissions across all accounts

- shared responsibility model: the line of responsibility shifts based on the level of abstraction provided by the service
  - aws is responsible for security **OF** the cloud
    - buildings, computers, storage, network systems, base software (hypervisors, etc)
  - I am responsible for security **IN** the cloud
    - you protect what you do inside the cloud
      - files, data, etc

## table of contents

### compute

- [Auto Scaling Group](#auto-scaling-group)
- [Batch](#batch)
- [EC2 (Elastic Compute Cloud)](#ec2-elastic-compute-cloud)
- [Elastic Beanstalk](#elastic-beanstalk)
- [Fargate](#fargate)
- [Lambda](#lambda)
- [Lightsail](#lightsail)

### storage

- [EFS (Elastic File System)](#efs-elastic-file-system)
- [FSx](#fsx)
- [S3 (Simple Storage Service)](#s3-simple-storage-service)
- [Storage Gateway](#storage-gateway)

### database

- [Aurora](#aurora)
- [DocumentDB](#documentdb)
- [DynamoDB](#dynamodb)
- [ElastiCache](#elasticache)
- [Neptune](#neptune)
- [QLDB](#qldb-quantum-ledger-database)
- [RDS (Relational Database Service)](#rds-relational-database-service)

### analytics and big data

- [Athena](#athena)
- [EMR (Elastic MapReduce)](#emr-elastic-mapreduce)
- [Glue](#glue)
- [Kinesis](#kinesis)
- [QLDB (Quantum Ledger Database)](#qldb-quantum-ledger-database)
- [Redshift](#redshift)

### networking and content delivery

- [API Gateway](#api-gateway)
- [CloudFront](#cloudfront)
- [Direct Connect](#direct-connect-site-to-site-vpn)
- [Elastic Load Balancer (ELB)](#elb-elastic-load-balancer)
- [Global Accelerator](#global-accelerator)
- [Route 53](#route-53)
- [VPC (Virtual Private Cloud)](#vpc-virtual-private-cloud)
- [Wavelength](#wavelength)

### security

- [Artifact](#aws-artifact)
- [CloudHSM](#cloudhsm-cloud-hardware-security-module)
- [Cognito](#cognito)
- [Detective](#detective)
- [Directory Services](#aws-directory-services)
- [GuardDuty](#guardduty)
- [IAM (Identity and Access Management)](#iam-identity-and-access-management)
- [Inspector](#inspector)
- [KMS (Key Management Service)](#kms-key-management-service)
- [Macie](#macie)
- [STS (Security Token Service)](#sts-security-token-service)
- [Secrets Manager](#aws-secrets-manager)
- [Security Groups](#security-groups)
- [Security Hub (CSPM)](#security-hub-cspm-cloud-security-posture-management)
- [Shield](#shield)
- [WAF (Web Application Firewall)](#waf-web-application-firewall)

### billing and costing management

- [Billing Alarms](#billing-alarms): notifications to monitor billing
- [Billing Dashboard](#billing-dashboard): high level overview + free tier dashboard
- [Budgets](#budgets): track usage, costs, reserved instances and get alerts
- [Compute Optimizer](#compute-optimizer): recommends resources configurations to reduce cost
- [Consolidated Billing](#consolidated-billing): centralized billing across all aws accounts in an aws organization
- [Cost Allocation Tags](#cost-allocation-tags): tag resources to create detailed reports
- [Cost Anomaly Detection](#cost-anomaly-detection): detect unusual spending using machine learning
- [Cost Explorer](#cost-explorer): view detailed current usage and forecast usage
- [Cost and Usage Reports](#cost-and-usage-reports): most comprehensive billing dataset
- [Pricing Calculator](#pricing-calculator): estimates costs in the cloud
- [Savings Plans](#savings-plans): easy way to save based on long-term usage of aws of compute services (ec2, fargate, lambda)
- [Service Quotas](#service-quotas): notifies you when you're close to service quota threshold

### devops

- [CDK (Cloud Development Kit)](#cdk-cloud-development-kit)
- [Cloud9](#cloud9)
- [CloudFormation](#cloudformation)
- [CodeBuild](#codebuild)
- [CodeCommit](#codecommit)
- [CodeDeploy](#codedeploy)
- [CodePipeline](#codepipeline)
- [CodeStar](#codestar)
- [ECR (Elastic Container Registry)](#ecr-elastic-container-registry)
- [Elastic Container Service (ECS)](#ecs-elastic-container-service)
- [Elastic Kubernetes Service (EKS)](#eks-elastic-kubernetes-service)
- [EventBridge](#eventbridge)
- [OpsWorks](#opsworks)
- [Step Functions](#step-functions)
- [X-Ray](#x-ray)

### ai and machine learning

- [aws connect](#aws-connect)
- [CodeGuru](#codeguru)
- [Comprehend](#comprehend)
- [Forecast](#forecast)
- [GuardDuty](#guardduty)
- [Kendra](#kendra)
- [Lex](#lex)
- [Macie](#macie)
- [Personalize](#personalize)
- [Polly](#polly)
- [QuickSight](#quicksight)
- [Rekognition](#rekognition)
- [SageMaker](#sagemaker)
- [Textract](#textract)
- [Transcribe](#transcribe)
- [Translate](#translate)

### application integration

- [API Gateway](#api-gateway)
- [SNS (Simple Notification Service)](#sns-simple-notification-service)
- [SQS (Simple Queue Service)](#sqs-simple-queue-service)

### management and governance

- [AWS Config](#aws-config)
- [CloudTrail](#cloudtrail)
- [CloudWatch](#cloudwatch)
- [Control Tower](#control-tower)
- [Health Dashboard](#health-dashboard)
- [SSM (Systems Manager)](#ssm-systems-manager)
- [Service Quotas](#service-quotas)
- [Trusted Advisor](#trusted-advisor)
- [Well-Architected Framework](#aws-well-architected-framework)

## AWS CLI

> manage aws resources with CLI commands

- list all IAM users in your AWS account: `aws iam list-users`
- get secret from aws secrets manager: `aws secretsmanager get-secret-value --secret-id peso-env-secret --region sa-east-1`

## EC2 (Elastic Compute Cloud)

> compute service that allows you to launch virtual servers in the cloud

- instance: a virtual server in the cloud
- AMI: a template for your instance (OS + configs)
- EBS (Elastic Block Store): persistent block storage for instances
- Security Groups: virtual firewalls that control traffic
- key pairs: used for SSH access to Linux or RDP to Windows instances
- elastic IP: static public IP address for your instance

- inbound traffic: traffic that comes into a resource from the outside
  - instance <= traffic = outside
- outbound traffic: traffic that goes out from a resource to an external destination
  - instance = traffic => outside

### workflow

1. choose an AMI (ubuntu, amazon linux, windows, etc)
1. choose an instance type (e.g. t2.micro, m5.large)
1. configure instance details (VPC, subnet, IAM role, etc.)
1. add storage (EBS volumes)
1. add tags (e.g. MyServer)
1. configure security group (firewall rules)
1. launch (with a key pair for SSH)

### shared responsibility model

> defines which security and compliance tasks are handled by AWS and which are handled by the customer

> [!IMPORTANT]
> Varies slightly depending on the aws service.
> depends partially in which category the service belongs: IaaS, PaaS, SaaS.
> Not every IaaS service has exactly the same responsibility model. The same applies to PaaS and SaaS categories.

|                     | On-site | IaaS             | PaaS                 | SaaS             |
|---------------------|---------|------------------|----------------------|------------------|
| **Data**            | You     | You              | You                  | You              |
| **Applications**    | You     | You              | You/Service provider | Service provider |
| **Runtime**         | You     | You              | Service provider     | Service provider |
| **Middleware**      | You     | You              | Service provider     | Service provider |
| **OS**              | You     | You              | Service provider     | Service provider |
| **Virtualization**  | You     | Service provider | Service provider     | Service provider |
| **Servers**         | You     | Service provider | Service provider     | Service provider |
| **Storage**         | You     | Service provider | Service provider     | Service provider |
| **Networking**      | You     | Service provider | Service provider     | Service provider |

- data: information your application uses (e.g. customer records, documents, transactions)
  - you are ALWAYS responsible for your data
- applications: the actual software you use or build (e.g. web app, analytics tool)
- runtime: the environment where code runs (e.g. JVM, node.js, python interpreter)
- middleware: software that connects applications or services (e.g. message brokers, APIs, databases, caching systems)
- os: the system software that runs on a server or VM (e.g. Linux, Windows Server)
- virtualization: software that abstracts physical hardware into virtual machines (e.g. a hypervisor like VMware or Xen)
- servers: physical or virtual machines that process workloads (CPUs, RAM, etc)
- storage: where your data is saved (e.g. hard drives, SSDs, cloud object storage like Amazon S3)
- networking: The communication layer (e.g. routers, switches, firewalls, internet access, VPCs, load balancers)
  - it allows components to talk to each other or external users

- you/customer: responsibility for security in the cloud
  - customer data
  - platform, applications, IAM
  - OS, networking traffic protection (encryption, integrity, identity)
  - client-side and server-side encryption
  - authentication
- aws: responsibility for security of the cloud
  - compute, storage, database, networking

### purchasing options

- on-demand: ideal for short-term use, expensive
  - predictable pricing
  - highest cost but no upfront payment
  - no long-term commitment
- reserved: commit to 1 or 3 years, cheaper
  - long workloads
  - up to 72% discount compared to on-demand
- spot instances: bid for unused capacity, cheapest
  - no guaranteed availability, aws can terminate them when the spot price exceeds your bid price
- savings plans: commit to an amount of usage
  - discount based on long-term usage
- capacity reservations: reserve instance capacity in a specific AZ
- dedicated instances: instances that run on hardware dedicated to your account, but aws manages the host
- dedicated hosts: get an entire physical server to yourself

### sizing and configuration options

- os
- cpu
- ram
- storage space
  - network-attached (ebs and efs)
  - hardware (ec2 instance store)
- network card
- firewall rules: security group
- ec2 user data: script that is launched at the first start of an instance

#### right sizing

> process of looking at deployed instances and identifying opportunities to eliminate or downsize without compromising capacity or other requirements, which results in lower costs

- when to right size:
  - before a cloud migration
  - continuously after the cloud onboarding process (requirements change over time)

- tools that can help: cloudwatch, cost explorer, trusted advisor

### ec2 instance tenancy

- shared (default): multiple aws accounts may share the same physical hardware
- dedicated instance: instance runs on single-tenant hardware
- dedicated host: instance runs on a physical server with ec2 instance capacity fully dedicated to your use

### ec2 image builder

> automate the creation, testing and distribution of AMIs or container images

- can be run on a schedule (weekly, whenever packages are updated, etc)

1. ec2 image builder initiates temporary builder ec2 instance
1. temporary builder ec2 instance builds the image
1. once the image is built, it is saved as an AMI
1. new ec2 instance is launched to test ec2 instance, if it fails, the pipeline stops here
1. if the test is successful, the AMI is made available to other aws regions or accounts

### local ec2 instance store

> fast, ephemeral block-level storage that is physically attached to the host machine running your ec2 instance

- block-level storage: disk is composed by blocks
  - each block has a unique address
  - os can read/write any block
- high-performance hardware disk
- **loses storage if stopped (ephemeral)**
- good buffer/cache/temporary content
- risk of data loss if hardware fails
- backups and replication are your responsibility

### EBS (Elastic Block Store)

> provides block-level storage that can be attached to ec2 instances

- tightly integrated with Amazon EC2
- block-level storage: refers to a type of data storage where data is saved in fixed-sized chunks called blocks
  - each block has its own address
  - the system can read/write to these blocks individually

#### EBS snapshots

> backup for an ebs volume

- it's recommended to detach ebs volume to do a snapshot
- use ebs snapshots as a buffer to copy ebs volumes across AZ

## IAM (Identity and Access Management)

> manage access to AWS services and resources securely

- root account/user: has complete and unrestricted access to all aws services and billing features in the account
  - original identity created when you dign up for aws account
  - root user is not an IAM user
  - capabilities
    - change account settings
    - close your aws account
    - change or cancel your aws support plan
    - register as seller in the reserved instance marketplace
- IAM user: individual user identity created and managed within an AWS account using the IAM service

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

- `Version`: the version of the policy language, defined by date
- `ID`: optional identifier for the policy
- `Statement`: contains individual permission rules
  - `"Sid": "AllowListBuckets"` (Statement ID): unique identifier for this permission block
  - `Effect`: defines if this permission is allowed or not
  - `"Action": "s3:GetObject"`: allows the use to list objects in the root of the specified S3 bucket
  - `Resource`: list of resources to which the actions apply
  - `Condition`: conditions for when this policy is in effect

- IAM Password Policy: You can decide the requirements for password creation

- MFA (Multi Factor Authentication): multi-step account login process that requires users to enter more information than just a password
  - e.g. password && (code sent to email || secret question || fingerprint scan)

- IAM Policy: set of permissions for aws resources
  - a policy can exist without a role, but a role can't do anything useful without a policy
- IAM Roles: identity that grants temporary permissions to aws services or users to perform tasks
  - set of IAM policies

### IAM role

> identity in aws that you can assume temporarily to get specific permissions

- use cases
  - grant temporary access to aws services without long-term credentials
  - assign permissions to services like ec2 or lambda
  - to switch roles across accounts securely

- key components of a role:
  - trust policy: defines who can assume the role (e.g. ec2, lambda, another account)
  - permissions policy defines what the role can do (e.g. access s3, write logs)
  - session duration
  - assume role

### shared responsibility

- aws
  - infrastructure
  - configuration and vulnerability analysis
  - compliance validation
- me
  - users, groups, roles, policies, management and monitoring
  - enable MFA on all accounts
  - rotate all keys often
  - use iam tools to apply appropriate permissions

### IAM Security Tools

- IAM Credentials Report (account-level)
- IAM Access Advisor (user-level)
  - use this information to revise your policies (least privilege principle)
- IAM policy simulator: test and debug IAM policies to check what actions are allowed/denied for specific users, groups, roles

### IAM Guidelines and Best Practices

- don't use the root account (except for aws account setup)
- don't create multiple aws accounts, create aws users within an aws account
- assign users to groups and assign permissions to groups
  - if you are responsible for the company resources
    - create a groups with the policy "administrator access" and create an aws user that belongs to this group
- create strong password policy
  - use and enforce MFA
- create and use roles to give permisssions to aws srevices

- dedicated host: your instance runs on a physical server fully dedicated to your use
  - an isolated server with configurations you can control

### IAM access analyzer

> used to find out which resources are shared outside your trusted boundaries (called a Zone of Trust)

- checks these types of resources:
  - s3 buckets, iam roles, kms keys, lambda functions, sqs queues, Secrets Manager secrets
- iam access analyzer creates a finding if a resource (like an s3 bucket or iam role) allows access from:
  - an aws not in your organization
  - the public internet
  - anonymous users
  - unrelated external aws account

## security groups

> facilitates managing network traffic

- acts as a "firewall"
- supports ALLOW rules only
- define:
  - access to ports
  - authorized IP ranges (IPv4 and IPv6)
  - control of inbound and outbound network traffic
- can be attached to multiple instances
- locked down to a region/VPC combination
- lives "outside" the EC2
  - if traffic is blocked, the EC2 instance won't see it

> [!TIP]
> it's good to keep 1 separate security group for SSH access.
> if your application is not accessible (time out), then it's a security group issue.
> if you application gives a "connection refused" error, then it's an application error or it's not launched.

> [!IMPORTANT]
> all inbound traffic is blocked by default
> all outbound traffic is authorized by default

- Ports:
  - 22 = SSH (Secure Shell): log into a linux instance
  - 3389 = RDP (Remote Desktop Protocol): log into a windows instance
  - 21 = FTP (File Transfer Protocol)
  - 22 = SFTP (Secure File Transfer Protocol)
  - 80 = HTTP : access unsecured websites
  - 443 = HTTPS: access secured websites

#### security groups vs firewalls

| Feature  | AWS Security Groups                                       | Traditional Firewalls (e.g., `ufw`, `iptables`) |
|----------|-----------------------------------------------------------|-------------------------------------------------|
| Level    | AWS-level (cloud)                                         | OS-level (server)                               |
| Type     | Virtual firewall for EC2 and services                     | Software firewall on the OS                     |
| Rules    | Stateful (return traffic automatically allowed)           | Stateful or stateless                           |
| Scope    | Applied per EC2 instance                                  | Applies to whole server                         |
| Use Case | Allow inbound on port 22 from your IP, or 443 to everyone | Block/allow ports directly on EC2 OS            |

### SSH Access

1. create an identity file `.pem` for SSH:
1. EC2 Dashboard => Key Pair => Create Key Pair
1. copy the public IP of the instance (used in the SSH command)
1. move the `.pem` file to a secure location and restrict permissions: `chmod 0400 <file>.pem`
1. SSH into the instance: `ssh -i <file>.pem ec2-user@<public_IP>`

## VPC (Virtual Private Cloud)

> virtual network that allows you to launch AWS resources in a logically isolated section of the cloud

- allows management over IP addresses, subnets, routing and security
- must have a CIDR block
- elastic ip: fixed public IPv4 address attached to ec2 instance
  - ongoing cost if not attached to ec2 instance or if the instance is stopped
- allows the creation of public and private subnets
  - subnet: smaller network inside larger network
    - helps organize and manage traffic in a network by dividing it into chunks
    - subnets allow better allocation of IPs from a larger IP range
    - each subnet in aws is associated with one route table
    - use CIDR notation (e.g. `192.168.1.0/24`)

- **public vs private subnet**
  - public subnets in AWS need an Internet Gateway to allow internet access
    - this means:
      - instances can send traffic to the internet
      - instances can receive traffic from the internet, if security rules allows it
  - private subnets do not have route to internet gateway
    - this means:
      - instances CANNOT initiate outbound internet traffic (unless via NAT gateway)
      - instances CANNOT receive ANY inbound traffic directly from the internet

| Feature        | Route Tables       | Security Groups            | Internet Gateway               |
| -------------- | ------------------ | -------------------------- | ------------------------------ |
| **Scope**      | Subnet-level       | Instance-level             | VPC-level                      |
| **Purpose**    | Direct traffic     | Allow or deny traffic      | Enable internet connectivity   |
| **Layer**      | Layer 3 (Network)  | Layer 4 (Transport)        | Layer 3 (Network)              |
| **Stateful?**  | No                 | Yes                        | Yes                            |
| **Controls**   | Where traffic goes | Whether traffic is allowed | Provides external access point |
| **Applied To** | Subnets            | EC2 instances, ENIs        | Entire VPC                     |

- route table: defines how traffic flows within the VPC (between subnets) and between the VPC and external networks (like the internet), by specifying the next hop for outbound traffic
  - hop: one step in the path that network traffic takes from a source to a destination
    - e.g. if a packet goes from a computer → to a router → to another router → to the destination server, that's 3 hops
  - only controls outbound routing (cannot restrict inbound traffic)
  - contains rules like:
    - to reach the internet (0.0.0.0/0), go through the Internet Gateway
    - to reach the private subnet, stay local
  - each subnet is associated with a route table
  - only one route table per subnet is allowed
  - a vpc can have multiple route tables
- internet gateway: allows instances in vpc to send traffic to the internet and enables the instances to receive responses from the internet
  - attached to a VPC
  - required for ec2 instances in public subnets to:
    - download packages
    - be accessed via SSH or a browser
- NAT (Network Address Translation) gateway (aws-managed): managed aws service that allows instances in a private subnet to connect to the internet, but prevents the internet from initiating a connection back
  - e.g. private ec2 can download updates or access external APIs without being publicly exposed
  - usually are placed in a public subnet and route private subnet traffic through it
- NAT instances (self-managed): is an ec2 instance configured manually to perform the same function as NAT gateway

- Network ACL (Access Control List): security layer used in networks to control inbound/outbound traffic at the subnet level
  - stateless: rules must be defined for both inbound and outbound traffic
  - supports ALLOW and DENY rules
  - rules only include IP addresses

### ENI (Elastic Network Interface)

> manages network connectivity for an ec2 instance in a vpc

- allows you to create multiple network interfaces for an ec2 instance
- use cases
  - creating a management network interface
  - creating a network interface for traffic that needs to be isolated from other traffic

### vpc flow logs

> feature that captures information about the ip traffic going to and from network interfaces

- network interface: virtual network card that connects ec2 instances and other aws resources to a vpc network
- captures network information from aws managed interfaces like ELB, elasticache, rds, aurora, etc
- vpc flow logs data can go to S3, cloudwatch logs, kinesis data firehose
- what the vpc flow logs capture
  - source ip and destination ip
  - source port and destination port
  - protocol (tcp, udp)
  - traffic acceptance (accept or reject)
  - bytes transferred

### endpoints

> privately connect VPC to aws services without using public IPs or going through the internet

- provides better security and lower latency to access aws services
- VPC endpoint gateway if you want to connect to s3 or DynamoDB
- VPC endpoint interface if you want to connect to other aws services

### peering

> connect 2 VPCs privately using aws' network and make them behave as if they are the same network

- IP address ranges can't overlap

### PrivateLink

> privately connect to a service in a 3rd party vpc

- most secure and scalable way to
  - expose a service to 1000s of vpcs
- doesn't require
  - vpc peering
  - internet gateway
  - nat
  - route tables
- requires
  - network load balancer
  - [ENI](#eni-elastic-network-interface)

### direct connect + site-to-site vpn

- site-to-site vpn: connect on-premises network to your aws vpc
  - secure "bridge" between physical data center and cloud network
- direct connect (DX): dedicated network connection between your on-premises, bypassing the public internet

### client vpn

> securely access AWS resources and on-premises networks that are not publicly accessible (i.e. private networks)

- the connection goes over the public internet, but the traffic is encrypted end-to-end
- allows your device to act as if it's inside the private network

### transit gateway

> connecto to multiple VPCs and/or on-premises infrastructure through a centralized hub

- works with
  - connect gateway
  - vpn connections
  - vpc
  - subnets
  - internet gateway
  - nat gateway/instances
  - nacl
  - security groups
  - vpc peering
  - elastic ip
  - vpc endpoints
  - PrivateLink
  - vpc flow logs
  - site-to-site vpn
  - client vpc
  - direct connect
  - transit gateway

## lambda

> serverless compute service that lets you run code in response to events

- short-lived functions: each invocation runs in an isolated environment with max duration of 15 minutes
- stateless: each function runs independently
- price based on number of calls and duration of compute time
- increasing ram also improves cpu and network
- event-driven
  - examples of events
    - file uploads
    - http requests
    - database changes
- supports
  - node.js
  - python
  - golang
  - java
  - c# (.NET core and powershell)

### workflow

1. write a function in Python, Node.js, Java, Go, etc
1. deploy it to Lambda (manually, with the CLI, or using IaC like Terraform)
1. configure a trigger (e.g. HTTP endpoint, S3 upload)
1. lambda runs your code when the event occurs
1. get billed per request and per compute time (in milliseconds)

use case example: resize images uploaded to an S3 bucket

1. upload triggers an S3 event
1. aws Lambda receives the event, processes the image (e.g. resizes it)
1. stores the output back to S3

## RDS (Relational Database Service)

> managed database service that allows you to set up, operate and scale relational databases in the cloud

- supports multiple engines:
  - [postgresql](./postgresql.md)
  - [mysql](./mysql.md)
  - [mariadb](./mariadb.md)
  - oracle
  - microsoft sql server
  - aurora

- read replica: is a copy of a database that can be used to offload read operations from the primary database
  - doesn't contribute to high availability, since they are all located in a single AZ
- multi-AZ: failover in case of AZ outage (high availability)
- multi-region (uses read replicas): spans multiple aws regions (e.g. `sa-east-1` and `us-east-1`)
  - disaster recovery in case of region issue
  - how it works:
    - Your main DB lives in one region (e.g. eu-west-1 in Ireland)
    - you create read replicas in other regions (e.g. U.S. or Asia)
    - local apps can read from these replicas — reduces latency
    - writes still go to the main DB only

## S3 (Simple Storage Service)

> scalable object storage service that allows you to store and retrieve data from anywhere on the web

- use cases
  - storage
  - backup and restore
  - application hosting
  - static websites
  - [data lake](/data_warehouse.md#types-of-data-repositories) and big data analytics
  - web hosting (static files)
  - media storage and distribution
  - application assets (images, audio, etc)
  - log storage

- properties
  - serverless: no server or infrastructure to manage
  - scalable: automatically scales to any amount of data
  - durable
  - event-driven: can trigger lambda, sqs, sns on object creation or deletion
  - durable: 11 9s (99.999999999%) durability

- versioning: option to retain multiple versions of objects
- there are no real folders
- max object size is 5 TB
  - if file is bigger than 5 TB, upload is segmented into multiple parts

- server-side encryption is enabled by default (instead of client-side encryption)
- you can enable public access with an S3 bucket policy
- to give user access to S3, use IAM policy to give access
- to give access to an EC2 instance, always use EC2 instance roles with IAM permissions
- to give access to IAM User in another AWS account, create a bucket policy that allows access to that specific user

- key: unique identifier for an object within a bucket (often a filename or path)
  - must be unique in the bucket
  - similar to full file path
    - aren't directories, just long names that contain slashes ("/")
  - prefix: path without file name
- object (similar to file): stored data (file + metadata + unique key)
- url structure to an s3 object: `https://<bucket-name>.s3.<region>.amazonaws.com/<object-key>`
  - `https://my-unique-bucket.s3.us-east-1.amazonaws.com/documents/report.pdf`
- bucket (similar to directory): container that stores objects
  - must have unique name globally
  - created in a region
  - pay attention to naming requirements
    - 3 to 63 characters
    - no spaces or underscores
    - no leading or trailing dots or hyphens
    - names can only contain lowercase letters, numbers, periods and hyphens
- region: bucket location (affects latency)

> [!NOTE]
> When you upload data (objects), server-side encryption is enabled by default, rather than client-side encryption.
> Server-side encryption ensures that this stored data is encrypted so that even if someone gains access to the raw storage, they can't read your data.

- compatible databases:
  - [postgresql](./postgresql.md)
  - [mysql](./mysql.md)
  - [mariadb](./mariadb.md)
  - oracle
  - microsoft sql server
  - aurora

- bucket policies examples:
  - you can enable public access with an S3 Bucket Policy
  - to give user access to S3, use IAM policy to give access
  - to give access to an EC2 instance, always use EC2 instance roles with IAM permissions
  - to give access to IAM User in another AWS account, create a bucket policy that allows access to that specific user

> [!WARNING]
> to fix 403 error, you need to give public access to your bucket

### workflow

1. create a bucket (e.g. `my-app-assets` in `sa-east-1`)
1. upload an object (`logo.png`)
1. access it with: `https://my-app-assets.s3.amazonaws.com/logo.png`
1. use IAM policies or s3 bucket policies to control access

#### aws cli example

```bash
# create a bucket
aws s3 mb s3://my-bucket-name

# upload a file
aws s3 cp ./file.txt s3://my-bucket-name/

# download a file
aws s3 cp s3://my-bucket-name/file.txt .

# list files
aws s3 ls s3://my-bucket-name/
```

### shared responsibility model

- aws is responsible for
  - infrastructure
    - global security
    - durability
    - availability
    - sustain concurrent loss of data in 2 facilities
  - configuration and vulnerability analysis
  - compliance validation
- me
  - s3 versioning
  - s3 bucket policies
  - s3 replication setup
  - logging and monitoring
  - s3 storage classes
  - data encryption at rest and in transit

### storage classes

> determines availability, durability and cost

- durability: how many times an object is going to be lost by S3 (likelihood of data loss)
  - same for all storages classes: 99.999999999% (11 9's)
- availability: percentage of time where a service is available
  - varies depending on storage class

- standard: frequently accesses data
  - use cases: big data analytics, mobile and gaming applications, content distribution
  - instant data retrieval
  - most expensive
- intelligent-tiering: automatically moves data between frequent/infrequent tiers
  - tiers:
    - frequent access tier: default tier
    - infrequent access tier: objects not accessed for 30 days
    - archive instant access tier: objects not accessed for 90 days
    - archive access tier: configurable from 90 days to 700+ days
    - deep archive access tier: configurable from 180 days to 700+ days
- standard IA (Infrequent Access)
  - pay to retrieve data
  - 99% availability
  - use cases: disaster recovery, backups
- one zone-IA
  - high durability in a single AZ
  - 99.5% availability
- glacier: long-term storage
  - cheap
  - types:
    - glacier instant retrieval
    - glacier flexible retrieval
      - retrieval options:
        - expedited: 1 to 5 min (highest cost)
        - standard: 3 to 5 hours (medium cost)
        - bulk: 5 to 12 hours (lowest cost)
    - glacier deep archive
      - retrieval options:
        - standard: 12 hours (medium cost)
        - bulk: under 24 hours (lowest cost)

### Replication (CRR and SRR)

> automatically copy objects from one S3 bucket to another

- you must enable Versioning in source and destination buckets
- buckets can be in different accounts
- buckets can be in different aws accounts
- copying is asynchronous, i.e. data is not copied immediately, there might be a small delay after upload
- 2 types:
  - Cross-Region Replication (CRR): Replicates data to a bucket in a different AWS region
    - use cases: compliance, lower latency access, replication across accounts
  - Same-Region Replication (SRR): Replicates data to a bucket in the same AWS region
    - use cases: log aggregation, live replication between production and testing accounts

### s3 data migration

- use cases: large data cloud migrations, decommission (means data center retirement), disaster recovery

#### snow family

- snow code
  - 8 TB of HDD storage
  - 14 TB of SSD storage
- snowball edge: move less than 10 PBs of data
  - storage optimized: 80 TB HDD
  - compute optimized: 42 TB HDD capacity + 28 TB NVMe SSD
- snowmobile: literally a truck
  - use when transfering more than 10 PB
  - each has 100 PB (equals to 100,000 PB) of capacity (1 PB = 1000 TBs)

### EBS vs EFS vs S3

| Feature         | **EFS** (Elastic File System) | **EBS** (Elastic Block Store)  | **S3** (Simple Storage Service) |
| --------------- | ----------------------------- | ------------------------------ | ------------------------------- |
| **Type**        | File Storage                  | Block Storage                  | Object Storage                  |
| **Access**      | Shared (multiple EC2s)        | Single EC2 instance            | HTTP-based, global access       |
| **Use Case**    | Shared files, home dirs, apps | Databases, OS volumes          | Backup, static files, media     |
| **Protocol**    | NFSv4                         | Attached as a disk (block I/O) | REST API / HTTPS                |
| **Scalability** | Auto scales (no limit)        | Fixed size (resize manually)   | Virtually unlimited             |
| **Performance** | High (latency ms-level)       | Very high (low latency)        | Variable (higher latency)       |
| **Durability**  | High (across AZs)             | High (within an AZ)            | Very high (11 nines)            |
| **Pricing**     | Pay per GB + throughput    | Pay per provisioned (allocated up front) GB | Pay per GB + requests |
| **OS Support**  | Linux (only)                  | Linux and Windows              | Any (via API/SDK/browser)       |

- file storage
- block storage
- object storage

- use cases
  - efs: shared file system across multiple ec2 instances
  - ebs: high-performance disk for a single ec2 instance
    - TIP: EB S => S ingle instance
  - s3: object storage with scalability and durability

## Abuse

> report suspected aws resources used for abusive or illegal purposes

- abuse and prohibited behaviors:
  - spam
  - DoS attacks
  - DDoS attacks
  - intrusion attempts
  - hosting objectionable, prohibited or illegal content (e.g. phishing websites)
  - copyrighted content
  - distributing malware
- contact the aws abuse team via
  - aws abuse form
  - abuse@amazonaws.com

## amazon managed blockchain

> fully managed service that makes it easy to create, manage and scale blockchain networks

- transactions without the need for a trusted, central authority
- use cases
  - join public blockchain networks
  - create your own scalable private network
- compatible with hyperledger and ethereum frameworks

## AMI (Amazon Machine Image)

> a template that contains a software configuration (e.g. operating system, application server, applications) that is used to launch EC2 instances

> complete snapshot of a virtual machine

- customize an ec2 instance
- faster boot and configuration times
- can be copied across regions

- sources of AMIs
  - public AMIs: provided by aws
  - your own AMIs: created and maintained by you
  - aws marketplace AMIs: made by someone else (may be free or sold)

| Feature        | **AMI**                               | **Docker Image**                       |
| -------------- | ------------------------------------- | -------------------------------------- |
| **Used for**   | Launching **EC2 virtual machines**    | Launching **containers**               |
| **Includes**   | OS, configuration, apps               | App + its dependencies (usually no OS) |
| **Managed by** | EC2 (VMs)                             | Docker Engine / ECS / Kubernetes       |
| **Isolation**  | Full virtual machine (hardware-level) | Process-level isolation                |
| **Size**       | Typically large (GBs)                 | Typically small (MBs)                  |
| **Boots**      | An entire VM                          | A single app or service                |
| **Boot Time**  | Slower (typically 30s–2min)           | Fast (usually <1s)                     |
| **Why?** | Full virtual machine boots up: OS, kernel, networking, etc. | Container shares host OS and starts as a process |

## API Gateway

> fully managed service that makes it easy for developers to create, publish, maintain, monitor and secure APIs at any scale

client <= REST API => API gateway <= proxy requests => lambda <= CRUD => DynamoDB

- scalable and serverless API
- supports
  - rest
  - websocket apis
  - security
  - user authentication
  - api throttling
  - api keys
  - monitoring

## Athena

> serverless query service that allows you to perform analytics against s3 objects using SQL

- query files with sql
- analyze data in s3 using serverless sql
- supports csv, json and more
- use compressed or columns data to reduce cost

- use cases:
  - business intelligence
  - analytics
  - reporting
  - analyza and query vpc flow logs

## Aurora

> proprietary relational database engine that is part of rds

- fully managed
- highly available
- costs more than rds, but is more efficient
- supports postgresql and mysql
- auto scales in increments of 10 GB

## Auto Scaling Group

> feature that allows aws to automatically manage the number of instances based on demand

- available in the following services: ec2, cloudwatch, elb
- scaling strategies
  - manual scaling: you update the size of an ASG manually
  - dynamic scaling
    - simple/step scaling
      - when CloudWatch alarm is triggered (example CPU > 70%), then add 2 units
      - when CloudWatch alarm is triggered (example CPU < 30%), then remove 1 unit
    - target tracking scaling
      - average ASG CPU must stay around 40%
    - schedules scaling: anticipate scaling based on known usage patterns
  - predictive scaling: uses machine learning to predict future traffic

## aws artifact

> provides on-demand access to aws compliance documentation, agreements and audits

- doesn't do audits, only provides access to the results of audits done on aws

## aws certificate manager

> easily provision, manage and deploy SSL/TLS certificates

- SSL/TLS certificate provides encryption for websites (HTTPS)
- supports both public and private tls certificates
- free of charge for public tls certificates
- automatic tls certificate renewal
- integrates with ELB, cloudfront distributions, APIs on API gateway

## aws config

> assess, audit and evaluate the configurations of your aws resources

- records configurations and changes over time

- aws config rules: check if your aws resources comply with specific desired configurations or policies

- use cases
  - enforce encryption on RDS databases
  - ensure IAM policies don't allow wildcard actions
  - check whether all resources are in approved regions
  - make sure CloudTrail is enabled and logging correctly

## aws connect

> AI-powered application that provides contact center

- set up call centers or customer service systems

## aws directory services

> suite of managed directory services that makes it easy to set up, manage and scale directory services in the AWS Cloud

- directory services: systems that manage information about users, computers, printers and other resources in the network
  - usually on on-premises systems
- can integrate with microsoft active directory services
- microsoft active directory: directory service for windows domain networks
  - windows domain network: type of computer network in that manages user accounts, computers and other resources
    - features
      - centralized authentication: users log in with one set of credentials across all computers in the domain
      - domain controller: special server that runs active directory and handles logins, access rights and security policies

## aws iam identity center

> easy single login

- one login for all your
  - aws accounts in aws organizations
  - business cloud applications (e.g. salesforce, box, microsoft 365)
  - ec2 windows instances

## aws knowledge center

- contains most frequent and common questions and requests about:
  - popular services
  - compute
  - IoT
  - analytics
  - customer engagement
  - management and governance
  - application integration
  - database
  - migration and trasfer
  - account and billing management
  - developer tools
  - networking and content delivery
  - business applications
  - front-end web and mobile
  - security, identity and compliance
- written and maintained by aws staff

## aws network firewall

> deploy essential network protections for your VPCs

> filter traffic entering and leaving a VPC

- features
  - traffic filtering
  - inspect the content of packets, not just headers
  - logging and monitoring

## aws organizations

> centrally manage multiple aws accounts

- allows management of permissions
- AWS organizations Service Control Policies (SCPs): centrally manage and restrict permissions across all accounts
  - SCPs do not grand permissions, they **limit** them
  - does not apply to master account
  - applies to all the users and roles of the account, including root
- api is available to automate aws account creation

- Organizational Unit (OU): used to group accounts within the organization

## aws professional services and partner network

> global group of experts available to work alongside your team and help achieve cloud goals faster

- often partner with APN (Aws Partner Network) members
- APN: global community of partners that help customers build, market and sell their offerings on aws
  - offerings = products, services or solutions that apn provides to customers using aws infrastructure
- types of APN partners
  - APN technology partners: provide hardware, connectivity or software solutions
  - APN consulting partners: help customers desing, architect, build, migrate and manage workloads on aws
  - APN training partners: deliver aws-authored training to organizations and individuals
- aws competency program: recognize APN partners that have demonstrated technical expertise
- aws navigate program: helps apn partners improve specialized skills

## aws secrets manager

> stores secrets

- can force rotation of secrets after a period of time
- automate generation of secrets on rotations (uses lambda)
- integrates with [rds](#rds-relational-database-service)
- secrets are encrypted using [KMS](#kms-key-management-service)

## aws support plans

- basic/free: available to all aws customers by default
- developer
- business
- enterprise

| Plan       | Cost (Starting)     | Access Type          | Use Case         | Response Time (Critical) |
| ---------- | ------------------- | -------------------- | ---------------- | ------------------------ |
| Basic      | Free                | Docs & Forums        | Learning/Test    | Not available            |
| Developer  | $29/mo or 3%        | Email (Business hrs) | Dev/Test         | < 12–24 hrs              |
| Business   | $100/mo or % usage  | 24/7 Chat/Phone      | Production       | < 1 hr                   |
| Enterprise | $15k/mo or % usage  | TAM + 24/7 support   | Mission-Critical | < 15 mins                |

- TAM (Technical Account Manager)

## aws training and certification

> help individuals/organizations build cloud skills through online courses and credentials

## aws well-architected framework

> set of best practices and guidelines created by amazon to help cloud architects design and operate secure, resilient and efficient infrastructure for their applications

6 pillars:

1. operational excellence
1. security
1. reliability
1. performance efficiency
1. cost optimization
1. sustainability

- well-architected tool: free tool to review your architectures against the 6 pillars of well-architected framework
  - how to use:
    - select your workload and answer questions
    - review your answers against the 6 pillars
    - obtain advise

### operational excellence

> running and monitoring systems effectively and continually improving processes and procedures

- IaC
- automate the creation of annotated documentation after every build
- make frequent, small, reversible changes
- anticipate failure
- learn from all operational failures

> [!TIP]
> operational excellence is the devops pillar

#### services

- prepare
  - CloudFormation
  - AWS Config
- operate
  - CloudFormation
  - AWS Config
  - cloudtrail
  - cloudwatch
  - aws x-ray (tracy api calls/http requests)
- evolve
  - cloudformation
  - codebuild
  - codecommit
  - codedeploy
  - codepipeline

### security

> protect information, systems and assets

- identity and access management
- data protection (at rest and in transit)
- incident response
- threat detection

#### services

- identity and access management
  - IAM
  - STS (Security Token Service): grants temporary credentials to users or services
  - MFA token
  - AWS organizations: centrally manage multiple aws accounts
- detective controls
  - aws config
  - CloudTrail
  - CloudWatch
- infrastructure protection
  - CloudFront
  - VPC
  - shield
  - WAF (Web Application Firewall)
  - Inspector
- data protection
  - KMS
  - S3
  - ELB
  - EBS (Elastic Block Store)
  - RDS
- incident response
  - IAM
  - CloudFormation
  - CloudWatch events

### reliability

> ensure a system can recover from failures and meet customer demands

- automated recovery
- failure management
- distributed system design

#### services

- failure management
  - CloudFormation
  - route 53
  - s3

### performance efficiency

> use computing resources efficiently to meet system requirements as demand changes

- use serverless and managed services where possible
- monitor and improve performance
- test different instance types and configurations

#### services

- selection
  - auto scaling
  - lambda
  - ebs
  - s3
  - rds
- review
  - cloudformation
  - aws news blog
- monitoring
  - cloudwatch
- tradeoffs
  - rds
  - elasticache
  - snowball
  - cloudfront

### cost optimization

> ability to run systems to deliver business value at the lowest price point

- use cost-effective resources
  - spot instance
  - reserved instance
  - s3 glacier
- turn off unused resources
- monitor usage and budget
  - aws services to track expenditure
    - budgets
    - cost and usage report
    - cost explorer
    - reserved instance reporting

- match supply and demand
  - auto scaling
  - lambda
- optimize overtime
  - trusted advisor: analyzes your aws account provides recommendations to optimize your aws environment
    - improves security, performance, fault tolerance, service limits and cost optimization
  - cost and usage report
  - news blog

### sustainability

- optimize energy consumption
- improve efficiency across system lifecycle
- use managed services to reduce waste

## batch

> fully managed batch processing at any scale

- batch: tasks that can be run in parallel and don't require real-time interaction
  - typically large volumes of data
- batch jobs are defined as docker images and run on [ECS](#ecs-elastic-container-service)
- uses ec2 instances
  - you can configure whether aws batch should use spot instances, on-demand instances, or a mix

### batch vs lambda

- lambda
  - maximum execution time of 15 min, after that the function is forcefully terminated, regardless of whether it's done
  - supports specific runtimes (python, node.js, java, go, etc) that aws explicitly provides and maintains
  - limited temporary disk space (512 MB)
  - serverless
- batch
  - no time limit
  - able to run any runtime as long as it's packaged as a Docker image
  - can rely on EBS volumes or instance store for disk space
  - not serverless

## Budgets

> send alarms when cost exceeds budget

- types of metrics for budget restriction: usage, cost, reservation, savings plans

## CAF (Cloud Adoption Framework)

> framework provided by aws to assist adoption of cloud computing for your infrastructure

six key perspectives:

| Perspective | Focus Area                    | Example Stakeholders  |
| ----------- | ----------------------------- | --------------------- |
| Business    | Business value and goals      | Executives, finance   |
| People      | Skills and change management  | HR, training teams    |
| Governance  | Risk and compliance           | Risk officers, legal  |
| Platform    | Cloud architecture            | Architects, engineers |
| Security    | Data and asset protection     | Security teams        |
| Operations  | Manage/monitor cloud services | IT ops, admins        |

## CDK (Cloud Development Kit)

> define cloud infrastructure using familiar programming languages

- compatible with
  - javascript/typescript
  - python
  - java
  - .NET
- code is compiled into CloudFormation template (json/yaml)

## CloudFormation

> fully managed service that allows you to create and manage AWS resources using templates

- IaC, similar to terraform
- template: configuration file that defines infrastructure
- uses json or yaml
- easily generate diagrams of your templates
- no need to define ordering and orchestration
- use existing templates on the web
- support for almost all aws resources
- easily see the cost of each service

### workflow

1. you write a template
1. upload template to CloudFormation
1. CloudFormation provisions and manages the resources

## CloudFront

> CDN service that delivers data, videos, applications and APIs to customers globally with low latency and high transfer speeds

- CDN (Content Delivery Network): caches static content (images, css, js) at edge locations globally
- improves content delivery by caching the content at edge locations
- DDoS protections
- integrated with
  - shield
  - aws web application firewall
  - s3 bucket
  - custom origin (http)
    - which can be:
      - application load balancer
      - ec2 instance
      - s3 website (you must first enable bucket static s3 website)
      - any http backend

### cloudfront vs s3 cross region replication

- cloudfront
  - global edge network
  - files are cached for a TTL (Time To Live) (maybe a day)
  - great for static content that must be available everywhere
- S3 Cross Region Replication
  - must be setup for each region if you want replication to happen
  - files are updated in near real-time
  - read only
  - great for dynamic content that need to be available at low latency in few regions

## CloudHSM (Cloud Hardware Security Module)

> provides dedicated, single-tenant hardware to securely generate, store and manage cryptographic keys

- dedicated: the hardware is exclusively assigned to your organization
- single-tenant: only your organization uses the hardware

- types of CMKs (Customer Master Key)
  - customer-managed CMKs: created, managed and used by the customer
  - aws-managed CMKs: created, managed and used by aws
  - aws-owned CMKs: collection of CMKs that aws services own and managed, used across multiple accounts
  - CloudHSM keys: CMKs generated and stored in dedicated CloudHSM hardware
    - the customer has full control over key management
    - required for highly regulated environments

## CloudTrail

> get history of events or API calls made within your aws account via aws cli, aws console, aws sdk or aws services

- provides governance and auditing
  - governance: monitoring activity, ensuring compliance, auditability and accountability
- enabled by default
- to store logs long-term, put logs from cloudtrail into: cloudwatch logs or s3

> [!TIP]
> cloudtrAIl => ApI

### insights

> automates the analysis of cloudtrail events to detect unusual activity

## CloudWatch

> monitoring and observability service that provides data and insights

- monitors resources, applications,  performance and operational health
  - unified view of operational health
- respond to system-wide performance changes
- optimize resource utilization

### CloudWatch Alarm

> monitor metric or math expression and perform one or more actions based on it over a specific time periond

> trigger notifications for any metric

- alarm actions
  - auto scaling: increase/decrease ec2 instances "desired" count
  - ec2 actions: stop, terminate, reboot or recover ec2 instance
  - sns notifications: send notification into sns topic
- can choose period on which to evaluate the alarm
- statistic types: %, max, min, sum, sampleCount, etc
- alarm states
  - OK
  - INSUFICIENT_DATA
  - ALARM

### CloudWatch Logs

> logging service that collects, monitors and stores log data from various aws resources, applications and services in real time

- log groups: container for logs from similar sources (e.g. ecs app in production, specific lambda function)
- log streams: sequence of log events from a single source within a log group

- log group = folder
- log stream = file inside the folder
- log event = line in the file

- adjustable cloudwatch logs retention: can be 1 week, month, year, etc

### CloudWatch Logs for EC2

- by default, no logs from ec2 instance will go to cloudwatch
- run cloudwatch agent on ec2 to push log files
- make sure iam permissions are correct
- cloudwatch log agent can be setup on-premises too

### CloudWatch Metrics

> provides metrics from aws resources

- ec2 instances
  - cpu utilization
  - status checks
  - network
    - default metrics every 5 min (if you pay more, every 1 min)
- ebs volumes
  - disk reads/writes
- s3 buckets
  - BucketSizeBytes
  - NumberOfObjects
  - AllRequests
- billing
  - total estimated charge
- service limits
  - service api usage
- custom metrics

## Cloud9

> cloud IDE that can be used in the browser

- allows code collaboration in real-time

## CodeArtifact

> fully managed artifact repository service

- use cases:
  - publish your executable/binary in common dependency management tools (e.g. npm, yarn, pip, maven, gradle)
  - manage versioned software artifacts

- artifact: package or binary like `.jar`, `tgz` (tarball)

## CodeBuild

> fully managed build service that compiles source code, runs tests and produces software packages that are ready to deploy

- similar to github actions and gitlab ci
- fully managed, serverless
- pay-as-you-go pricing
- can be triggered by [CodePipeline](#codepipeline)
- use cases
  - compile source code
  - run tests
  - produce packages that are ready to be deployed (by CodeDeploy, for example)

## CodeCommit

> fully managed source control service that makes it easy for teams to host secure and highly scalable private Git repositories

- aws' github competitor
- source control service
- fully managed

## CodeDeploy

> fully managed deployment service that automates software deployments to a variety of compute services

- works with
  - ec2 instance
  - lambda
  - on-premise servers
- hybrid service
- servers/instances must be provisioned and configured ahead of time with CodeDeploy agent

## CodeGuru

> machine learning-powered service for automated code reviews and application performance recommendations

- features
  - codeguru reviewer: automated code reviews for static code analysis (development)
    - identifies critial issues, security vulnerabilities, bugs
  - codeguru profiles: visibility/recommendations about application performance during runtime (production)
    - visibility: visual representation of performance (e.g. cpu usage)
    - helps understand runtime behavior of your application
    - minimal overhead on application
    - features
      - identify and remove code inefficiencies
      - improve application performance (e.g. reduce CPU utilization)
      - decrease compute costs
      - anomaly detection

## CodePipeline

> service that orchestrates full CI/CD workflow

- fully managed
- can activate [codebuild](#codebuild)
- compatible with:
  - [CodeCommit](#codecommit)
  - [CodeBuild](#codebuild)
  - [CodeDeploy](#codedeploy)
  - [elastic Beanstalk](#elastic-beanstalk)
  - [CloudFormation](#cloudformation)
  - [github](./github.md)
  - 3rd party services
  - custom plugins

codecommit => codebuild => codedeploy => compute resource (can be ec2 instance, ecs, etc)

## CodeStar

> unified user interface to easily manage software development activities

- quick way to setup CodeCommit, CodePipeline, CodeBuild, CodeDeploy, Elastic Beanstalk, EC2, etc
- edit the code "in the cloud" using AWS Cloud9

## Cognito

> fully managed identity service that allows you to add user sign-up, sign-in and access control to your apps quickly and easily

- similar to auth0 and firebase
- identity for web/mobile application's users
- don't create an IAM user for the clients of your application, use Cognito
- instead of giving your app's users aws iam accounts (which are meant for admins and systems), you use cognito to manage their identities securely
- also capable of signing in with google/facebook/twitter accounts

## Comprehend

> uses Natural Language Processing (NLP) to extract insights about the content of documents

- fully managed and serverless service
- uses machine learning to find insights and relationships in text
- capabilities
  - analyzes text using tokenization
  - organize a collection of text files by topic
  - analyze customer interactions as positive or negative

## Consolidated Billing

- when enabled, combine usage across all aws accounts in the aws organization to share
  - volume pricing
  - reserved instances
  - savings plans discounts

## Control Tower

> automatically set up aws organizations to organize accounts and implements SCPs (Service Control Policies)

- easy way to set up and govern a secure multi-account aws environment
- benefits
  - automate the setup of multiple aws accounts in a few clicks
  - automatically apply recommended security, compliance and operational settings
  - automate ongoing policy management using guardrails

## Cost Anomaly Detection

> continuously monitor your cost and usage using ML to detect unusual spending

## Detective

> investigates and identifies the root cause of security issues or suspicious activities using ML and graphs

- generates visualizations with details and context to get to the root cause

## DMS (Database Migration Service)

> migrate databases to aws

- resilient and self healing
- source database remains available during the migration
- supports
  - homogeneous migration (e.g. oracle to oracle)
  - heterogeneous migration (e.g. sql server to aurora)

## DocumentDB

> fully managed aws implementation of mongodb

- nosql database
- store, query and index json data
- highly available with replication across 3 AZ

## DynamoDB

> fully managed NoSQL database service that provides fast and predictable performance with seamless scalability

- key-value database
- highly available with replication across 3 AZ
- auto scalability
- distributed "serverless" database, scales to massive workloads
- single-digit millisecond latency retrieval
- integrated with IAM for
  - security
  - authorization
  - administration

### DynamoDB tables

> data storage

- read/write data
- auto scaling
- high availability

### DynamoDB streams

> change log for table

- captures real-time changes: inserts, updates, deletes
- it's optional

### DAX (DynamoDB Accelerator)

> fully managed in-memory cache for DynamoDB

- 10x performance improvement
- secure, highly scalable and highly available

## ECS (Elastic Container Service)

> fully managed container orchestration service that allows you to run, stop, and manage Docker containers on a cluster of EC2 instances

## ECR (Elastic Container Registry)

> fully managed Docker container registry that makes it easy to store, manage, and deploy Docker container images

- private docker registry
- stored docker images can be run by [ECS](#ecs-elastic-container-service) or [fargate](#fargate)

## EFS (Elastic File System)

> fully managed cloud-based file storage service

> provides a network file system

- network file system:
- scalable
- allows shared access (100s of ec2 instances)
- accessible concurrently by multiple ec2 instances or other aws services
- can only be multi-AZ when using only ec2 instances

### pricing

- storage classes:
  - standard: for active, frequently accessed files
  - Infrequent Access (IA): lower cost for files accessed less often

- billing is based on:
  - storage used (GB/month)
  - throughput
  - requests and data transfer (for IA)

## EKS (Elastic Kubernetes Service)

> fully managed Kubernetes service that allows you to run, manage, and scale containerized applications using Kubernetes

## ElastiCache

> fully managed service that deploys, operates and scales popular open-source compatible in-memory data stores in the cloud

- used to manage redis or memcached
- in-memory db
- high performance, low-latency
- managed by aws
  - os maintenance
  - optimizations
  - setup
  - configuration
  - monitoring
  - failure recovery
  - backups

## ELB (Elastic Load Balancer)

> service that automatically distributes incoming application traffic across multiple targets

- targets can be ec2 instances, ip addresses, containers
- can have high availability (multi AZ)
  - multi-AZ ELB: can distribute traffic across multiple AZs
  - single-AZ ELB: can only send traffic to its own AZ
- improve fault tolerance
- improve scalability
- do regular health checks to your instances
- types of load balancers
  - application load balancer (network layer): uses HTTP/HTTPS/gRPC protocols
  - network load balancer (transport layer): ultra-high performance, allows for tcp
    - balances incoming network traffic
    - routes traffic based on IP, doesn't inspect request content
  - gateway load balancer (application layer): used for security, detect intrusions

## EMR (Elastic MapReduce)

> big data processing service that allows you to run Apache Hadoop, Spark and other big data frameworks on AWS

> create hadoop clusters to analyze and process vast amount of data

- clusters can be made of hundreds of ec2 instances
- takes care of all provisioning and configuration
- offers support for apache spark
- offers auto-scaling
- can be integrated with spot instances

- use cases
  - data processing
  - machine learning
  - web indexing
  - big data

## Elastic Beanstalk

> fully managed service that allows you to deploy and manage web applications and services

- developer is only responsible for the application
- PaaS
- uses CloudFormation to provision the application
- managed service
  - instance configuration and OS is handled by Beanstalk
  - deployment strategy is configurable but managed by Beanstalk
- load balancing and auto-scaling
- 3 architecture models: 
  - singe instance deployment
  - LB + ASG
    - good for production or pre-production web applications
  - ASG only
    - good for non-web apps in productions (workers, etc)

- support for:
  - go
  - java se
  - node.js
  - php
  - python
  - ruby
  - single container docker
  - multi-container docker
  - pre-configured docker
- CloudWatch integration for monitoring
  - checks app health
  - publishes health events

## EventBridge

> serverless event bus service that makes it easy to build event-driven applications at scale

- serverless
- use events to connect application components together

- event: json message that describes something that happened in a system
  - events come from aws services, custom apps or saas providers
- event bus: pipeline where events are sent and rules are applied to process them
- event source: origin of the event
- rules: match certain events and define what should happen when a match occurs
- targets: aws services or resources that receive matched events and act on them
  - common targets: lambda (run code), step functions (start workflow), sqs (enqueue messsage), sns (publish to topic), etc

event example:

```json
{
  "source": "aws.ec2",
  "detail-type": "EC2 Instance State-change Notification",
  "detail": {
    "instance-id": "i-1234567890abcdef0",
    "state": "terminated"
  }
}
```

## Fargate

> launch docker containers on aws

- fully managed
- allocates the exact cpu and ram requested

## Forecast

> fully managed service that uses ML to deliver highly accurate forecasts

- example scenario: predict the future sales of a raincoat
- use cases
  - product demand planning
  - financial planning
  - resource planning

## FSx

> fully managed service that allows you to launch high-performance file systems in the cloud

- amazon FSx for Windows File Server: for windoes-based applications that require SMB protocol and Active Directory integration
- amazon FSx for Lustre: used for machine learning, analytics, video processing, financial modeling
  - high performance computing
  - scalable file storage
  - lustre = linux + cluster

## global accelerator

> improve application availability and performance using the aws global network

- uses private network to optimize the route to your application

### cloudfront vs global accelerator

- both
  - use AWS global network
  - use edge locations around the world
  - integrate with AWS Shield for DDoS protection
- Cloudfront is a Content Delivery Network
  - improves performance for cacheable content
  - content is served at the edge
- Global Accelerator
  - no caching
  - proxying packets at the edge to the application over TCP or UDP
  - good for HTTP use cases that require static IP addresses
  - good for HTTP use cases that require deterministic, fast regional failover

## global applications

> applications deployed in multiple geographies

- global applications decrease latency
- Disaster Recovery (DR): If An AWS Region Goes Down (Earthquake, Storms, Power Shutdown, Politics)
  - in this cenario, you can switch to another region and keep the application working
  - DR plan increases the availability of your application
- attack protection: distributed global infrastructure is harder to attack

### global applications architecture

> how an application is set up to run across different regions

- active-active: multiple instances in different regions executes read/write operations
  - lowers read's latency
  - lowers write's latency
- active-passive: when 1 instance executes read/write, while the other only executes read
  - lowers read's latency
- multi-region
- single-region

## Glue

> managed Extract, Transform, Load (ETL) service

- prepare and transform data for analytics
- fully serverless
- glue data catalog: catalog of datasets
- can be used by [athena](#athena), [redshift](#redshift), [EMR](#emr-elastic-mapreduce)

## GuardDuty

> threat detection service that continuously monitors aws data sources

- uses machine learning algorithms and anomaly detection
- analyses the following data sources
  - cloudtrail events logs
  - vpc flow logs
  - dns logs

## Health Dashboard

> global service that provides personalized alerts and remediation guidance for aws outages that might impact your aws account or resources

### service health (public view)

- shows
  - all regions
  - all services health
  - historical information for each day

### your account (personalized view)

- provides
  - alerts and remediation guidance when AWS is experiencing events that may impact you
  - personalized view into the performance and availability of AWS services you are using
  - relevant or timely information to help you manage events in progress
  - proactive notification to help you plan for scheduled activities
  - aggregate data from entire AWS Organization

## Inspector

> automated security assessment service that helps improve security and compliance

- only for ec2 instances, ecr container images and lambda functions
- scans aws workloads for vulnerabilities and unintended network exposure
  - aws workloads: applications, services or processes that are running on aws infrastructure

## IQ

> find professional help for your aws projects

- engage and pay aws certified 3rd party experts for on-demand project work

for customers:

1. submit request
2. review responses
3. select expert
4. work securely
5. pay per milestone

for experts:

1. create profile
2. connect with customers
3. start a proposal
4. work securely
5. get paid

## Kendra

> fully managed document search service powered by ML

- extract answers from within a document (text, pdf, html, powerpoint, etc)
- natural language search capabilities
- can learn from user interactions/feedback to promote preferred results
- ability to manually fine-tune search results

## Kinesis

> fully managed service that allows you to collect, process and analyze real-time streaming data

- unlimited scalability
- cloud-native service: proprietary protocol from aws

## KMS (Key Management Service)

> managed service that allows you to create and control the encryption keys used to encrypt your data

### workflow

you can use aws management console or aws cli

1. you create a customer master key (CMK) in KMS
1. you use the CMK to encrypt your data
1. when you need to decrypt your data, you use the same CMK

## Lex

> add chatbots that speak natural language to applications

- uses Automatic Speech Recognition (ASR) to convert speech to text
- uses natural language understanding to recognize the intent of text, callers

## Lightsail

> deploy and manage applications/websites with pre-configured cloud resources

- simpler alternative to using ec2, rds, elb, ebs, route 53...
- high availability
- no auto-scaling
- limited aws integrations, kinda of a standalone service
- use cases
  - simple web applications
    - has templates for: LAMP, nginx, MEAN, node.js
  - websites
    - has templates for: wordpress, magento, etc

## Local Zones

> extend VPC to more locations

- local zones are an extension of an aws region
- compatible with ec2, rds, ecs, ebs, elasticache, direct connect

## Macie

> fully managed data security service that uses machine learning and pattern matching to discover and protect sensitive data

- helps identify and alert you to sensitive data, such as Personally Identifiable Information (PII) in s3 buckets

## MQ

> managed message broker service for rabbitmq and activemq

- doesn't scale as much as [sqs](#sqs-simple-queue-service) and [sns](#sns-simple-notification-service)

## Neptune

> fully managed graph database

- high availability
- replication across 3 AZ
- great for knowledge graphs

## OpsWorks

> manages chef and puppet

- chef and puppet are tools that perform server configuration automatically
- alternative to [SSM](#ssm-systems-manager)

## Outposts

> fully managed service that extends aws infrastructure, services and tools to virtually any on-premises or edge locations

> run aws services locally on your own hardware, while still being managed from the aws cloud

- you are responsible for the outposts rack physical security
- benefits
  - low latency access to on-premise systems
  - local data processing
  - data residency
  - easier migration from on-premises to the cloud
  - fully managed service
- compatible services: ec2, ebs, s3, eks, ecs, rds, emr

## Personalize

> fully managed ML service to build apps with real-time personalized recommendations

- same technology used by amazon
- integrates with
  - existing websites
  - applications
  - sms
  - email marketing systems
- implement in days, not months
- no need to build, train and deploy ML solutions
- use cases: retail stores, media, entertainment, etc

## Polly

> convert text to speech using deep learnign

- allows development of applications that talk

## QLDB (Quantum Ledger DataBase)

> fully managed ledger database

> track all changes to your application data over time and maintain a verifiable history of changes

- serverless
- centralized database, with one source of truth
- high availability
- replication across 3 AZ
- immutable system: no entry can be removed or modified, cryptographically verifiable

## QuickSight

> serverless machine learning-powered business intelligence serve to create interactive dashboards

- per-session pricing
- use cases
  - business analytics
  - building visualizations
  - perform ad-hoc analysis
- integrated with
  - RDS
  - Aurora
  - Athena
  - S3

## RAM (Resource Access Manager)

> allows you to grant access to the resources you own to other accounts

## Redshift

> fully managed data warehouse service that allows you to run complex queries on large datasets

- based on postgresql, but it's not used for OLTP
- used for OLAP (analytics and data warehouse)
- 10x better performance in comparison to other data ware houses
- load data every hour, not every second
- has sql interface for queries
- can be integrated with aws tools, such as quicksight, tableau

## Rekognition

> image recognition and video analysis with ML

- capable of:
  - face analysis for
    - user verification
    - identifying gender, age range, emotions
    - face search
  - people counting
  - labeling
  - content moderation

## re:Post

> aws-managed Q&A service

> forum to ask technical questions to experts

- free, similar to stackoverflow
- don't use it to make questions that are time-sensitive or involve any proprietary information

## Route 53

> scalable domain name system (DNS) web service that translates domain names into IP addresses

- managed dns
- hosted zone: logical container within aws route 53 that holds DNS records for domains
  - DNS records: maps domain name to a specific resource (e.g. IP address, mail server)
    - types:
      - `A` (Address): points a domain/subdomain directly to a public IPv4 address
      - `AAAA`: url => ipv6
      - `CNAME` (Canonical NAME): redirects one domain or subdomain to another domain or subdomain
      - `MX` (Mail Exchange): specifies mail exchange servers for receiving email
      - `TXT` (TeXT): holds text information, often used for verification purposes or SPF records (email authentication)
    - what is a dns record?
    - explain each type of dns record
    - logical container: not a docker container, it's a logical group/repository with dns records
  - NS (NameServer records): indicate authoritative server for the domain
  - authoritative NS record: specifies which nameservers have definitive information about dns records for the domain
  - domain registrar: responsible for managing official domain records, registration, renewals and assignments

```
Visitor enters URL → Domain Registrar → Nameservers → Hosted Zone → DNS Records → IP Address
```

### registering domain

1. access route 53 in AWS console
1. select "Register Domains"
1. insert/search/choose a domain
1. enter contact information
1. review details, pricing and accept terms (select to renew domain yearly, if you want to)
1. finalize registration

- registering a domain automatically cretes a hosted zone
- if you delete it, you can easily copy the registered domain's name servers and create a new record

example of Name Servers:

```
ns-123.awsdns-01.com
ns-1234.awsdns-01.org
ns-123.awsdns-01.net
ns-1234.awsdns-01.co.uk
```

### routing policies

> control how DNS queries are answered

- simple routing policy: route traffic to a single resource (e.g. a web server or an S3 bucket)
  - no health checks
  - no request distribution
- weighted routing policy: distributes requests across resources (ex: EC2 instances)
  - has health checks
- latency routing policy: uses latency as criteria
  - has health checks
- failover routing policy: for high availability and disaster recovery
  - for disaster recovery
  - has primary instance and failover instance
  - has health checks

## SageMaker

> fully managed service for developers to train machine learning models that can make predictions

lets say you want to build a model that predicts your exam score

- SageMaker can help in all steps below and do more:
  - gather past data
  - label with a score (0 to 10 in the exam)
  - build a ML model
  - train and tune ML model
  - SageMaker now make predictions

## Security Hub CSPM (Cloud Security Posture Management)

> gives centralized view of aws security status

- helps identify security issues across aws accounts and services
- continuously monitors security best practices and compliance standards
- aggregates alerts from various aws services
  - config, GuardDuty, inspector, macie, iam access analyzer, aws firewall manager, aws health, aws partner network solutions

## Service Catalog

> allows organizations to manage/distribute approved resources (can be applications or aws resources)

## Service Quotas

> notifies you when you are close to a service quota value threshold

## Shield

> protects networks and applications by analyzing network security configurations and providing managed DDoS protection

- shield standard: free protection from DDoS attacks, SYN/UDP floods, reflection attacks
- shield advanced: paid 24/7 DDoS protection
  - 24/7 access to aws DDoS response team

## SNS (Simple Notification Service)

> fully managed messaging service that allows you to send notifications to a large number of recipients

- unlimited scalability
- cloud-native service: proprietary protocol from aws
- no message retention

- event publishers: sends message to one SNS topic
- event subscribers: listens to the sns topic notifications
  - each subscriberto the topic will get all the messages
  - services that can be target subscribers:
    - sqs
    - lambda
    - kinesis data firehose
    - emails
    - sms
    - mobile notifications
    - http(s) endpoints

## SQS (Simple Queue Service)

> message queuing service that allows you to decouple and scale microservices, distributed systems and serverless applications

- fully managed
- unlimited scalability
- low latency
- messages are kept for 14 days
- messages can be
  - xml
  - json

- producer service: produces requests to sqs
- consumer service: consumers requests from sqs
- decoupling: when both producers and consumers scale independently from each other

## SSM (SystemS Manager)

> centrally view, manage, and operate nodes (ec2 or on-premises) at scale in AWS, on-premises and multicloud environments

- run commands, patch and configure servers
- hybrid service
- not just a single service, it is a suite of products
- features
  - patching automation for enhanced compliance
  - run commands across entire fleet of servers
  - store parameter configuration with the SSM Parameter Store

### workflow

1. install SSM agent onto the systems we control
1. its installed by default on Amazon Linux AMI and Ubuntu AMI
1. if an instance can't be controlled with SSM, it's probably an issue with SSM agent

## Step Functions

> fully managed service that allows you to coordinate multiple AWS services into serverless workflows

## Storage Gateway

> hybrid cloud storage service that enables on-premises applications to seamlessly use aws cloud storage

- bridge between on-premise data and cloud data in S3
- use cases
  - back up and archive on-premises data to aws
  - provide low-latency access to frequently used data via local caching
  - migrate data to aws over time
  - extend on-premises storage without buying new hardware

## STS (Security Token Service)

> create temporary, short-term credentials to access your aws resources with limited privileges

## Textract

> extracts text, handwriting and data from any scanned documents using AI and ML

- read and process any type of document (PDFs, images, etc)

## Translate

> natural and accurate language translation

## Trascribe

> automatically converts speech to text

- uses deep learning process called Automatic Speech Recognition
- fully managed
- automatically removes Personally Identifiable Information (PII) using Redaction

## Trusted Advisor

> analyze your aws accounts and provides recommendations

- provides recommendations on 5 categories:
  - cost optimization
  - performance
  - security
  - fault tolerance
  - service limits

## WAF (Web Application Firewall)

> create security rules that control bot traffic and block common attack patterns

> protects web application from common web exploits (layer 7)

- deploy on ALB, API gateway, cloudfront

## Wavelength

> brings aws services to the edge of the 5G networks

- ultra-low latency applications through 5G networks
- traffic doesn't leave Communication Service Provider's (CSP) network
- no additional charges or service agreements
- use cases
  - smart cities
  - ML-assisted diagnostics
  - connected vehicles
  - interactive video streams
  - AR/VR
  - online gaming

## x-ray

> visual analysis of applications services

> helps developers gain insights into the performance/behavior of applications running in a microservices architecture or distributed system

- distributed tracing service
  - tracing service: tool that helps developers track/visualize the flow of requests through an application

- advantages
  - troubleshooting performance (bottlenecks)
  - understand dependencies in a microservice architecture
  - pinpoint service issues
  - review request behavior
  - find errors and exceptions
  - identify users that are impacted by problems

## aws common tasks

### install docker on amazon linux 2

```bash
#!/bin/bash
sudo yum update -y
sudo yum -y install docker
sudo service docker start
sudo usermod -a -G docker ec2-user
sudo chmod 666 /var/run/docker.sock
```

## pricing

- global AWS services: IAM, CloudFront, Amazon Route 53, WAF, Amazon Chime, DynamoDB, WorkDocs, WorkMail, WorkSpaces, WorkLink, Service Certificates

- AWS services that include a free tier for 12 months for new AWS customers:
  - EC2, S3, RDS, EBS, ECR, CloudFront, ElastiCache, DynamoDB*, Glacier*, Lambda*

- Free AWS services: Security Groups, Auto Scaling, CloudFormation
  - Lambda (1M requests/month)
  - DynamoDB (25 GB storage)
  - CloudWatch (10 custom metrics and alarms)
  - Amazon Q Developer (Unlimited autocomplete suggestions)
  - Aurora DSQL (100k DPUs)
  - CloudFront (1 TB of data trasfer out)
  - aws blogs, forums, guides, quick starts

- What AWS services are never free (need to pay to use): WAF, Inspector, Route 53, EBS volumes, ELB (WIREE)

### models

- pay as you go
- pay per use
- pay less by using more
- pay less when you reserve
- pay less when AWS grows
- no up-front investment

### data transfer

- inbound (data going into aws services): free
- within the same Availability Zone (AZ): free
- from aws to the internet (first 100GB/month): free

### networking costs per GB

> [!TIP]
> use private IP to save money and better network performance.
> use same AZ for maximum savings (at the cost of high availability)

TODO

## tutorials

- [setup kubernetes in ec2 instance tutorial](https://varunmanik1.medium.com/setting-up-a-kubernetes-cluster-on-aws-ec2-with-ubuntu-22-04-lts-and-kubeadm-5c54930a4659)

## best practices

### account

- least privilege principle: concede the minimum amount of privileges as possible
- don't use root account (except for aws account setup)
- don't create multiple aws accounts, create aws users within an aws account
- operate multiple accounts using aws organizations
- use SCP (Service Control Policies) to restrict account power
- easily setup multiple accounts with best-practices with AWS Control Tower
- use tags and cost allocation tags for easy management and billing
- iam guidelines: MFA, least-privilege, password policy, password rotation
- config to record all resources configurations and compliance over time
- cloudformation to deploy stacks across accounts and region
- trusted advisor to get insights, choose a support plan adapted to your needs
  - support plan: subscription for trusted advisor features
    - basic/free, developer, business, enterprise
- send service logs and access logs to s3 or cloudwatch logs
  - if your account is compromised: change the root password, delete and rotate all passwords/keys, contact the aws support
- cloudtrail to record api calls made within your account
- if your account is compromised: change the root password, delete and rotate all passwords/keys, contact AWS support

### architecting and ecosystem

- stop guessing your capacity needs
- test systems at production scale
- automate to make architectural architectural experimentation easier
- allow for evolutionary architecture
  - design based on changing requirements

### security

- penetrationg testing
  - you dont need prior approval to run penetrationg tests for the following services
    - ec2 instances
    - nat gateways
    - ELB
    - RDS
    - cloudfront
    - aurora
    - api gateways
    - lambda and lambda edge functions
    - lightsail resources
    - elastic beanstalk environments
    - this list can increase overtime
  - prohibited activities
    - DNS zone walking via route 53 hosted zones
    - Denial of Service (DoS)
    - Distributed Denial of Service (DDoS)
    - Simulated DoS
    - Simullated DDoS
    - Port flooding
    - protocol flooding
    - request flooding
      - login request flooding
      - API request flooding
    - for any other simulated events, contact aws-security-simulated-event@amazon.com

### DDoS protection on aws

- aws shield standard
- aws shield advanced
- aws waf
- cloudfront and route 53

---

- server provisioning: the process of setting up physical or virtual hardware; installing and configuring software, such as the operating system and applications; and connecting it to middleware, network, and storage components
- failover: ability of a system/service to automatically switch to a backup/secondary system when the primary system becomes unavailable or experiences a failure
  - important for ensuring high availability and minimizing downtime for critical applications and services
- topology: the arrangement of nodes or devices in a network
  - e.g. star, ring, bus, mesh, tree
- cluster: a group of interconnected computers or servers that work together to perform a specific task or function
  - commonly used in high-performance computing, data processing and distributed systems to improve performance, scalability, and fault tolerance
  - resources are shared and coordinated to achieve a common goal
- grid: distributed computing system that connects multiple computers or servers to work together on a common task or problem
  - decentralized
- hybrid service: service that can operate both in the cloud and on-premises

