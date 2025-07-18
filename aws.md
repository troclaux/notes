
# AWS

> cloud computing platform that provides a wide range of services and tools for building and managing applications and infrastructure in the cloud

- cloud computing: the delivery of computing services over the internet
  - advantages:
    - trade fixed expense for variable expense
    - benefit from massive economies of scale
    - stop guessing capacity
    - increase speed and agility
    - stop spending money running and maintaining data centers
    - go global in minutes
    - on demand resources can be provisioned without human interaction
    - resources are available over the network, and can be accessed by diverse client platforms
    - quickly and easily scale based on demand (stop guessing capacity)
    - measured payment for resources (pay only what you use)
    - trade CAPital EXpense (CAPEX) for OPerational EXpense (OPEX)
      - capital expense: no large upfront costs for servers, buildings, etc
      - operational expense: only pay for what you use
    - easily access global infrastructure
- service: a software that provides functionality and performs a task or set of tasks for your system
  - e.g. apache, nginx, postgresql, auth0
- infrastructure: the physical or virtual resources that support the operation of a system
  - e.g. servers, storage, network, databases, etc
- managed service: aws handles maintenance and scaling, but you still configure some infrastructure components
- fully-managed service: aws handles everything (infrastructure, scaling, patching, backups, etc)
  - you don't manage servers or infrastructure
  - serverless: run code or services without managing any servers at all
    - is a type of fully-managed service, but not all fully-managed services are serverless
    - e.g. aws lambda, aws step functions, Amazon EventBridge, azure functions, google cloud functions

> [!TIP]
> Fully managed (RDS, ElastiCache, EMR): you still pick instance types/sizes. AWS handles patching, scaling, backups.
> Serverless (Lambda, Fargate, DynamoDB, S3): you don't manage or choose servers. Autoscaling is built-in.

- options to manage aws
  - AWS management console
  - AWS Command Line Interface (CLI)
  - AWS Software Developer Kit (SDK): allows developers to interact with aws services using programming languages
    - allows the access and management of aws services programmatically
    - supports python, java, javascript, .NET, ruby, golang, etc
    - Used whenever you want to call APIs

- reliability: ability of a system to perform its intended function consistently over time without failure
- High Availability (HA): design and implementation of systems that are resilient to failures and can continue operating with minimal downtime
  - optimizes its uptime/availability (much higher than in a normal system)
    - able to quickly recover from failures, usually in an automated way
    - there can still be downtime/unavailability while recovering, but the outage will be much shorter than without HA
  - running instances for the same application on at least 2 Availability Zones
  - key characteristics
    - redundancy
    - failover mechanisms
    - geographic distribution
    - health monitoring
    - scalability and elasticity
- fault tolerance: ability to continue operating despite failures in some components
  - no downtime, even if in the process of auto-healing
- durability: ability of a service to retain and protect data over time without loss
  - measure of how likely your data is to survive intact even in the face of hardware failures, disasters or other problems
- scalability: ability of a system to increase or decrease its resources to meet changing demand
  - elasticity: automate scalability

- aws global infrastructure
  - region: separate geographic areas (e.g. `us-east-1` or `sa-east-1`)
    - each region is completely independent from others
    - each region has 3 to 6 Availability Zones, with few exceptions (AZ ⊂ region)
  - Availability Zone (AZ): isolated locations within each region (e.g. `us-east-1a`, `us-east-1b`, `us-east-1c`, `us-east-1d`)
    - each AZ has its own power, cooling and networking, so they do not share single points of failure
  - edge location: physical facility AWS uses to cache and deliver content to users with low latency

- aws global services: services that are NOT tied to a specific region and operate across all regions
  - IAM, Route 53, CloudFront, WAF, Shield, aws organizations, Artifact, DynamoDB, WorkSpaces, Health Dashboard, Trusted Advisor, SNS
  - TIP: if a service manages access, identity or DNS for your entire aws environment, it's likely global
- aws regional services: services that are tied to a specific region
  - e.g. s3, lambda, ec2, rds

- global applications: applications deployed in multiple geographies
  - global applications decrease latency
  - Disaster Recovery (DR): if an AWS Region goes down (earthquake, storms, power shutdown, politics)
    - in this cenario, you can switch to another region and keep the application working
    - DR plan increases the availability of your application
  - attack protection: distributed global infrastructure is harder to attack

> [!WARNING]
> Not all fully-managed services are global.
> Not all global services are fully-managed.

## table of contents

- [AWS Documentation](https://docs.aws.amazon.com/): official aws service documentation provided by amazon
- [AWS Knowledge Center](#aws-knowledge-center): offers official advice to common support questions and troubleshooting
- [AWS Prescriptive Guidance](#aws-prescriptive-guidance): provides best practices and architectural patterns from aws experts
- [AWS Security Blog](#security-blog): covers AWS security topics, updates and best practices
- [AWS Security Center](#security-center): official resource for security and compliance information from aws
- [AWS re:Post](#repost): community-driven Q&A site with AWS expert and user responses

### analytics and big data

- [Amazon Athena](#athena): query s3 objects with sql
- [Amazon EMR (Elastic MapReduce)](#emr-elastic-mapreduce): big data processing
- [AWS Glue](#glue): ETL operations
- [Amazon Kinesis](#kinesis): real-time streaming data
- [Amazon OpenSearch Service](#opensearch-service): keyword-based data analytics
- [Amazon QuickSight](#quicksight): data visualization for BI
- [Amazon Redshift](#redshift): data warehouse

- [Amazon QLDB (Quantum Ledger Database)](#qldb-quantum-ledger-database-out-of-scope)

### application integration

- [Amazon EventBridge](#eventbridge): event bus, task scheduler
- [Amazon SNS (Simple Notification Service)](#sns-simple-notification-service)
- [Amazon SQS (Simple Queue Service)](#sqs-simple-queue-service)
- [AWS Step Functions](#step-functions): workflow orchestration

- [Amazon API Gateway](#api-gateway)

### business applications

- [Amazon Connect](#amazon-connect): set up call centers or customer service systems
- [Amazon SES (Simple Email Service)](#ses-simple-email-service): send emails

### cloud financial management

- [AWS Budgets](#budgets): track usage, costs, reserved instances and get alerts
- Cost and Usage Reports: most comprehensive billing dataset, most granular report in aws (view billing activity of last month)
- Cost Explorer: view detailed current/historical usage, forecast future costs, create custom reports and identify under-utilized resources
  - used for analyzing costs after migrating to aws
  - estimate costs **after** migrating to aws
- [AWS Marketplace](#marketplace): search, buy and deploy 3rd-party software and services

- Billing Alarms: notifications to monitor billing
- Billing Dashboard: high level overview + free tier dashboard
- [Consolidated Billing](#consolidated-billing): centralized billing across all aws accounts in an aws organization
- [Cost Allocation Tags](#cost-allocation-tags): tag resources to create detailed reports
- [Cost Anomaly Detection](#cost-anomaly-detection): detect unusual spending using machine learning
- Pricing Calculator: estimates costs in the cloud, also get detailed pricing **before migration**
- Savings Plans: easy way to save based on long-term usage of aws of compute services (ec2, fargate, lambda)
- [Service Quotas](#service-quotas): notifies you when you're close to service quota threshold

### compute

- [AWS Batch](#batch): batch processing
- [Amazon EC2 (Elastic Compute Cloud)](#ec2-elastic-compute-cloud): virtual servers in the cloud
- [AWS Elastic Beanstalk](#elastic-beanstalk): deploy and manage web applications and services
- [Amazon Lightsail](#lightsail): deploy and manage applications/websites with pre-configured cloud resources
- [AWS Outposts](#outposts): run aws services locally on your own hardware, while still being managed from the aws cloud

### containers

- [Amazon ECR (Elastic Container Registry)](#ecr-elastic-container-registry): docker container registry
- [Amazon ECS (Elastic Container Service)](#ecs-elastic-container-service): docker container orchestration
- [Amazon EKS (Elastic Kubernetes Service)](#eks-elastic-kubernetes-service): manage containers with kubernetes

### customer enablement

- [AWS Support Plans](#aws-support-plans)

### database

- [Amazon Aurora](#aurora): proprietary relational database engine that is part of rds
- [Amazon DocumentDB](#documentdb): fully-managed aws implementation of mongodb
- [Amazon DynamoDB](#dynamodb): fully-managed NoSQL (key-value/document) database service with automatic scalability
- [Amazon ElastiCache](#elasticache): fully-managed in-memory database (redis/memcached)
- [Amazon Neptune](#neptune): fully-managed graph database
- [Amazon RDS (Relational Database Service)](#rds-relational-database-service): fully-managed relational database service

### developer tools

- [AWS CLI](#aws-cli)
- [AWS CodeBuild](#codebuild): compiles source code, runs tests and produces software packages for deployment
- [AWS CodePipeline](#codepipeline): orchestrate full CI/CD workflow
- [AWS X-Ray](#x-ray): visual analysis of application services

### end user computing

- [Amazon AppStream 2.0](#appstream-20): streams desktop applications to users over the internet
- [Amazon WorkSpaces](#amazon-workspaces): provides virtual desktops to users that can be accessed from anywhere
- [Amazon WorkSpaces Secure Browser](#amazon-workspaces-secure-browser)

### frontend web and mobile

- [AWS Amplify](#amplify): build and deploy full-stack web and mobile application quickly
- [AWS AppSync](#appsync): develop graphql APIs

### internet of things (iot)

- [AWS IoT Core](#iot-core): allows connected IoT devices to interact with aws services and other devices securely

### machine learning

- [Amazon Comprehend](#comprehend): uses Natural Language Processing (NLP) to extract insights from documents
- [Amazon Kendra](#kendra): document search powered by ML
- [Amazon Lex](#lex): chatbots that speak natural language to applications
- [Amazon Polly](#polly): converts text to speech using deep learning
- [Amazon Q](#amazon-q): AI assistant that helps users analyze data, write code and answer questions using natural language
- [Amazon Rekognition](#rekognition): image recognition and video analysis with ML
- [Amazon SageMaker](#sagemaker-ai): trains machine learning models that can make predictions
- [Amazon Textract](#textract): extracts text, handwriting and data from any scanned documents using AI and ML
- [Amazon Transcribe](#trascribe): converts speech to text using ML
- [Amazon Translate](#translate): natural and accurate language translation of text using ML

### management and governance

- [AWS Auto Scaling](#auto-scaling-group-asg)
- [AWS CloudFormation](#cloudformation)
- [AWS CloudTrail](#cloudtrail): get history of events or API calls made within your aws account
- [AWS Compute Optimizer](#aws-compute-optimizer): recommendations to optimize your aws resources
- [Amazon CloudWatch](#cloudwatch): real-time system health and performance monitoring
- [AWS Config](#aws-config): assess, audit and evaluate the configurations of your aws resources
- [AWS Control Tower](#control-tower): automatically set up aws organizations to organize accounts and implements SCPs (Service Control Policies)
- [Health Dashboard](#health-dashboard-your-account-personalized-view)
- [AWS License Manager](#license-manager): helps you manage software licenses
- [AWS Service Catalog](#service-catalog): internal control over approved AWS resources for your organization
- [Service Quotas](#service-quotas)
- [AWS SSM (SyStems Manager)](#ssm-systems-manager): centrally view, manage and operate nodes (ec2 or on-premises)
- [AWS Trusted Advisor](#trusted-advisor): analyzes your aws accounts with predefined rules and provides recommendations
- [AWS Well-Architected Framework](#aws-well-architected-framework)

- [AWS CAF](#caf-cloud-adoption-framework)

### migration and transfer

- [AWS Application Discovery Service](#application-discovery-service): plan migrations to aws
- [AWS Application Migration Service](#application-migration-service): lift-and-shift service to migrate application
- [AWS Database Migration Service (DMS)](#dms-database-migration-service): migrates databases to aws quickly and securely
- [Migration Evaluator](#migration-evaluator): assess cost and feasibility of migrating to aws
- [AWS Migration Hub](#migration-hub): central dashboard to track and manage application migrations to aws
- [AWS Schema Conversion Tool (AWS SCT)](#sct-schema-convertion-tool): migrates databases from one database engine to another
- [AWS Snow Family](#snow-family)

### networking and content delivery

- [Amazon API Gateway](#api-gateway): helps developers manage APIs at scale
- [Amazon CloudFront](#cloudfront): CDN service that delivers data globally with low latency and high trasfer speeds
- [AWS Direct Connect](#hybrid-connectivity): dedicated private network connection between on-premises infrastructure and aws
- [AWS Global Accelerator](#global-accelerator): improve application availability and performance using the aws global network
- [AWS PrivateLink](#privatelink): privately connect to a service in a 3rd party vpc
- [Amazon Route 53](#route-53): DNS web service
- [AWS Transit Gateway](#transit-gateway): connect to multiple VPCs and/or on-premises infrastructure through a centralized hub
- [Amazon VPC (Virtual Private Cloud)](#vpc-virtual-private-cloud)
- [AWS VPN](#hybrid-connectivity)
- [AWS site-to-site VPN](#hybrid-connectivity)
- [AWS client VPN](#hybrid-connectivity)

### security, identity and compliance

- [AWS Artifact](#aws-artifact): provides aws compliance documentation, agreements and audits
- [AWS Audit Manager](#audit-manager): automate audit preparation for compliance with regulations and standards
- [AWS Certificate Manager](#aws-certificate-manager): easily manage SSL/TLS certificates
- [AWS CloudHSM](#cloudhsm-cloud-hardware-security-module): dedicated, single-tenant hardware to manage cryptographic keys
- [Amazon Cognito](#cognito)
- [Amazon Detective](#detective): identifies the root cause of security issues using ML
- [AWS Directory Service](#aws-directory-service)
- [AWS Firewall Manager](#aws-firewall-manager): centrally manage and configure firewall rules across aws accounts & resources
- [Amazon GuardDuty](#guardduty): threat detection service that continuously monitors aws data sources
- [AWS IAM (Identity and Access Management)](#iam-identity-and-access-management)
- [AWS IAM Identity Center](#aws-iam-identity-center)
- [Amazon Inspector](#inspector): find software vulnerabilities in compute resources
- [AWS KMS (Key Management Service)](#kms-key-management-service)
- [Amazon Macie](#macie): uses ML to discover and protect sensitive data
- [AWS RAM (Resource Access Manager)](#ram-resource-access-manager): grants access to the resources you own to other accounts
- [AWS Secrets Manager](#aws-secrets-manager): store secrets and automate secrets rotation
- [AWS Security Hub (CSPM)](#security-hub): centralized security monitoring that makes compliance checks
- [AWS Shield](#shield): protects networks and applications against DDoS attacks
- [AWS WAF (Web Application Firewall)](#waf-web-application-firewall): create security rules that control bot traffic and block common attack patterns

- [AWS Security Center](#security-center)
- [Security Groups](#security-groups)

### serverless

- [AWS Fargate](#fargate): compute engine for running containers on AWS
- [AWS Lambda](#lambda): compute service that lets you run code in response to events

### storage

- [AWS Backup](#aws-backup): backup data across aws services
- [Amazon Elastic Block Store (Amazon EBS)](#ebs-elastic-block-store)
- [Amazon EFS (Elastic File System)](#efs-elastic-file-system)
- [AWS Elastic Disaster Recovery](#drs-disaster-recovery-service-or-elastic-disaster-recovery): recover after a disaster
- [Amazon FSx](#fsx): launch high-performance file systems for windows-based applications
- [Amazon S3 (Simple Storage Service)](#s3-simple-storage-service)
- [Amazon S3 Glacier](#storage-classes)
- [AWS Storage Gateway](#storage-gateway): extends storage of on-premises environments to aws cloud storage

## AWS CLI

> manage aws resources with CLI commands

- list all IAM users in your AWS account: `aws iam list-users`
- get secret from aws secrets manager: `aws secretsmanager get-secret-value --secret-id peso-env-secret --region sa-east-1`

## EC2 (Elastic Compute Cloud)

> compute service that allows you to launch virtual servers in the cloud

- every EC2 instance is launched in a VPC

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

### purchasing options

- on-demand instances: ideal for short-term use, expensive
  - predictable pricing
  - highest cost but no upfront payment
  - no long-term commitment
- reserved instances (RI): commit to 1 or 3 years, cheaper
  - up to 72% discount compared to on-demand
  - allow instance size flexibility within the same instance family
    - but only for Linux and in a Region and when using regional RIs (not zonal)
  - if you pick convertible reserved instances, you can change instance family later, OS, tenancy or region
    - as long as the new instance is of equal or greater value
- savings plans: flexible pricing model where you commit to an amount of usage for 1 or 3 years
  - discount based on long-term usage
  - best for users who want cost savings, but with more flexibility than reserved instances
  - has largely **superseded** RI as it offers more flexibility and similar or better savings
- spot instances: bid for unused capacity, cheapest
  - no guaranteed availability, aws can terminate them when the spot price exceeds your bid price
- capacity reservations: reserve instance capacity in a specific AZ
  - ensures that you have compute capacity available when you need it
- dedicated hosts: get an entire physical server to yourself
  - an isolated server with configurations you can control
  - ideal for BYOL model, because licenses are tied to the physical server
- dedicated instances: instances that run on hardware dedicated to your account, but aws manages the host
  - cheaper than dedicated host
  - you don't have control over hardware configuration

### instance types

- general pupose (t3, t4g, m5): balance of compute, memory and networking resources
  - best for high-performance web servers, batch processing, multiplayer gaming servers
- compute optimized (c5, c6g)
- memory optimized (r5, x1e): best for in-memory databases (e.g. redis), real-time data analytics, high-performance databases
- storage optimized (i3, d2): best for data warehouses, large-scale transactional databases, distributed file systems
- accelerated computing (p3, inf1): best for gpu-intensive workloads

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

### image builder

> automate the creation, testing and distribution of AMIs

- can be run on a schedule (weekly, whenever packages are updated, etc)
- fully-managed

1. ec2 image builder initiates temporary builder ec2 instance
1. temporary builder ec2 instance builds the image
1. once the image is built, it is saved as an AMI
1. new ec2 instance is launched to test ec2 instance, if it fails, the pipeline stops here
1. if the test is successful, the AMI is made available to other aws regions or accounts

### EC2 Auto Scaling

> provides automatic scaling for EC2 instances

> [!NOTE]
> Other AWS services (e.g. ECS, DynamoDB, Aurora) also support Auto Scaling but not through ASGs.

- relies on predictive scaling using ML

#### Auto Scaling Group (ASG)

> feature within EC2 Auto Scaling that allows aws to automatically manage the number of ec2 instances based on demand

- scaling strategies
  - manual scaling: you update the size of an ASG manually
  - dynamic scaling
    - simple/step scaling
      - when CloudWatch alarm is triggered (example CPU > 70%), then add 2 units
      - when CloudWatch alarm is triggered (example CPU < 30%), then remove 1 unit
    - target tracking scaling
      - average ASG CPU must stay around 40%
    - scheduled scaling: anticipate scaling based on known usage patterns
  - predictive scaling: uses machine learning to predict future traffic

> [!IMPORTANT]
> Auto Scaling is also available in other AWS services (without using ASGs), such as: ECS, DynamoDB and Aurora

### EBS (Elastic Block Store)

> provides block-level storage that can be attached to ec2 instances

- tightly integrated with Amazon EC2
- block-level storage: refers to a type of data storage where data is saved in fixed-sized chunks called blocks
  - each block has its own address
  - the system can read/write to these blocks individually
- EBS snapshots: backup for an ebs volume
  - use ebs snapshots as a buffer to copy ebs volumes across AZ

- volume types
  - gp3 (General Purpose SSD): default for most use cases
  - gp2 (Legacy General Purpose SSD): similar to gp3, but older generation
  - io2/io2 Block Express (Provisioned IOPS SSD): for high-performance databases
  - st1 (Throughput Optimized HDD): big data, data warehouses, log processing
  - sc1 (Cold HDD): for infrequently accessed data (archive)

### instance store volumes

> fast, ephemeral block-level storage that is physically attached to the host machine running your ec2 instance

- block-level storage: disk is composed by blocks
  - each block has a unique address
  - os can read/write any block
- high-performance hardware disk
- **loses storage if stopped (ephemeral)**
  - if your workload needs durable storage, use EBS instead of instance store
- good buffer/cache/temporary content
- risk of data loss if hardware fails
- backups and replication are your responsibility

## IAM (Identity and Access Management)

> manage access to AWS services and resources securely

### types of identity management

- aws account: container for all your aws resources and identities
  - when you create it, you start by setting up a root account/user
  - root user is tied to the aws account, but the aws account is not the root user
- root account/user: has complete and unrestricted access to all aws services, resources and billing features in the account
  - original identity created when you sign up for aws account
  - root user is not an IAM user
  - exclusive capabilities
    - change the aws account's settings (name, password, email, access keys, MFA)
    - close the aws account
    - active IAM access to the billing and cost management console
    - view certain tax invoices
    - register as seller in the ec2 reserved instance marketplace
    - change or cancel your aws support plan
    - if the only IAM admin accidentally revokes their own permissions, sign into root account to restore IAM user permissions
    - modify root user MFA settings
    - edit or delete policy that denies all principals for both SQS (only resource-based policies) or S3 bucket
  - security protection methods for root account:
    - enable MFA
    - create IAM users and avoid using root user for daily tasks
    - store the root user credentials securely, and use it only when absolutely necessary
    - use aws organizations for centralized management to reduce the need to access root account often
- IAM user: individual user identity created and managed within an AWS account using the IAM service
- IAM groups: collections of IAM users, permissions applied to the group apply to all members
- AWS organizations: allows management of multiple AWS accounts under one umbrella
  - allows management of permissions
  - AWS organizations Service Control Policies (SCPs): centrally manage and restrict permissions across all accounts
  - you can get discounts if consuming lots of resources across the accounts
- IAM roles: identity in aws that you can assume temporarily to get specific permissions
  - use cases
    - grant temporary access to aws services without long-term credentials
    - assign permissions to services like ec2 or lambda
    - to switch roles across accounts securely
  - key components of a role:
    - trust policy: defines who can assume the role (e.g. ec2, lambda, another account)
    - permissions policy defines what the role can do (e.g. access s3, write logs)
    - session duration
    - assume role
- federated access
  - federated identity: allows external users (e.g. organization's active directory or Google) to access aws without needing IAM user accounts
  - often done using AWS Single Sign-On (SSO) or SAML 2.0

### policy structure

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

### shared responsibility

- aws
  - infrastructure
  - configuration and vulnerability analysis
  - compliance validation
- me
  - management and monitoring of users, groups, roles and policies
  - enable and enforce MFA on all accounts
  - rotate all keys often
  - use iam tools to apply appropriate permissions
  - analyze access patterns and review permissions

### IAM Security Tools

- IAM Credentials Report (account-level)
  - generates a csv report that lists all IAM users in your account and the status of their credentials
- IAM Access Advisor (user-level)
  - shows to a user which aws services they have permission to access and when those services were last accessed
  - use this information to revise your policies (least privilege principle)
- IAM policy simulator: test and debug IAM policies to check what actions are allowed/denied for specific users, groups, roles

### IAM Guidelines and Best Practices

- don't use the root account (except for aws account setup)
- don't create multiple aws accounts, create aws users within an aws account
- assign users to groups and assign permissions to groups
  - if you are responsible for the company resources
    - create a groups with the policy "administrator access" and create an aws user that belongs to this group
- create strong password policy (use and enforce MFA)

- create and use roles to give permissions to aws services
- to give user access to S3, use IAM policy to give access
- to give access to an EC2 instance, always use EC2 instance roles with IAM permissions
- to give access to IAM User in another AWS account, create a bucket policy that allows access to that specific user


### IAM access analyzer

> used to find out which resources are accessible to entities outside your AWS account or organization (called a Zone of Trust)

- checks these types of resources:
  - s3 buckets, iam roles, kms keys, lambda functions, sqs queues, Secrets Manager secrets
- finding: alert that indicates a resource in your aws environment is accessible outside your Zone of Trust
  - IAM Access Analyzer creates a **finding** if a resource (such as an S3 bucket or IAM role) allows access from:
    - an aws not in your organization
    - the public internet
    - anonymous users
    - unrelated external aws account

## security groups

> virtual firewalls that control inbound and outbound traffic at the instance level

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
> It's good to keep 1 separate security group for SSH access.
> If your application is not accessible (time out), then it's a security group issue.
> If you application gives a "connection refused" error, then it's an application error or it's not launched.

> [!IMPORTANT]
> All inbound traffic is blocked by default.
> All outbound traffic is authorized by default.

#### security groups vs firewalls

SGs protect at resource-level, NACLs protect at subnet-level, AWS Network Firewall protects at VPC-level

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

## shared responsibility model

> defines which security and compliance tasks are handled by AWS and which are handled by the customer

> [!IMPORTANT]
> Varies slightly depending on the aws service.
> Depends partially in which category the service belongs: on-premises, IaaS, PaaS, SaaS.
> Not every IaaS service has exactly the same responsibility model. The same applies to PaaS and SaaS categories.
> Not every AWS service fits cleanly into the traditional on-premises/IaaS/PaaS/SaaS classification.

- aws & me
  - patch management:
    - aws is responsible for patching the infrastructure (hardware, software, networking and facilities)
    - the customer is responsible for patching their own guest operating systems and applications
  - configuration management
    - AWS maintains the configuration of its infrastructure
    - the customer is responsible for configuring their own instances, databases and applications

|                    | On-premises | IaaS             | PaaS                 | SaaS             |
|--------------------|-------------|------------------|----------------------|------------------|
| **Data**           | You         | You              | You                  | You              |
| **Applications**   | You         | You              | You/Service provider | Service provider |
| **Runtime**        | You         | You              | Service provider     | Service provider |
| **Middleware**     | You         | You              | Service provider     | Service provider |
| **OS**             | You         | You              | Service provider     | Service provider |
| **Virtualization** | You         | Service provider | Service provider     | Service provider |
| **Servers**        | You         | Service provider | Service provider     | Service provider |
| **Storage**        | You         | Service provider | Service provider     | Service provider |
| **Networking**     | You         | Service provider | Service provider     | Service provider |

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

![IaaS shared responsibility model](./images/aws_shared_responsibility_model.jpeg)

- examples of each category:
  - on-premises: aws outposts
  - IaaS: EC2, VPC, ELB, EBS
  - PaaS: Elastic Beanstalk, Amazon API Gateway, Fargate, RDS, Aurora
  - SaaS: WorkSpaces, QuickSight

## VPC (Virtual Private Cloud)

> virtual network that allows you to launch AWS resources in a logically isolated section of the cloud

- global service
- **NOT** fully managed, you have to manage networking components (e.g. subnets, route tables, ip address ranges, etc)
- allows management over IP addresses, subnets, routing and security
  - allows the creation of public and private subnets
    - subnet: segments of a VPC's IP address range where you can place groups of isolated resources
      - each subnet must be associated with a route table
      - use CIDR notation (e.g. `192.168.1.0/24`)

- elastic ip: fixed public IPv4 address attached to ec2 instance
  - ongoing cost if not attached to ec2 instance or if the instance is stopped

- **public vs private subnet**
  - public subnets: subnet that is configured to allow direct access to the internet
    - need an Internet Gateway to allow internet access, but not all public subnets have an Internet Gateway
    - this means:
      - instances can send traffic to the internet
      - instances can receive traffic from the internet, if security rules allows it
  - private subnets do not have route to internet gateway
    - this means:
      - instances CANNOT initiate outbound internet traffic (unless via NAT gateway)
      - instances CANNOT receive ANY inbound traffic directly from the internet

| Category           | Scope    | Purpose                                            | Layer               | Stateful? | Controls                                     | Applied To                  |
| ------------------ | -------- | -------------------------------------------------- | ------------------- | --------- | -------------------------------------------- | --------------------------- |
| Internet Gateway | VPC | Enable internet connectivity for resources in public subnets | Layer 3 (Network) | No       | Provides external access point               | Entire VPC                  |
| NAT Gateway        | Subnet   | Allow outbound internet access for private subnets | Layer 3 (Network)   | Yes       | Outbound internet for private resources      | Private subnet traffic      |
| Route Tables       | Subnet   | Direct traffic                                     | Layer 3 (Network)   | No        | Where traffic goes                           | Subnets                     |
| NACL (Network ACL) | Subnet   | Allow or deny traffic at the subnet level          | Layer 3 (Network)   | No        | Whether traffic is allowed in/out of subnets | Subnets                     |
| Security Groups    | Instance | Allow or deny traffic at the instance level        | Layer 4 (Transport) | Yes       | Whether traffic is allowed                   | EC2 and RDS instances, ENIs |

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
  - does **not** encrypt traffic
  - required for ec2 instances in public subnets to:
    - download packages
    - be accessed via SSH or a browser
- NAT (Network Address Translation) gateway: managed AWS service that allows instances in a private subnet to initiate outbound connections to the internet (e.g. for updates or API calls), but prevents the internet from initiating inbound connections to those instances
  - use case: allows a private ec2 to download updates or access external APIs without being publicly exposed
  - usually are placed in a public subnet and route private subnet traffic through it
  - does **not** encrypt traffic
- NAT instances (self-managed): is an ec2 instance configured manually to perform the same function as NAT gateway
- NACL (Network Access Control List): security layer used in networks to control inbound/outbound traffic at the subnet level

- stateful: tracks active connections and allows return traffic automatically
- stateless: doesn't track connection state
  - you must define both inbound and outbound rules explicitly, like with NACLs
  - supports ALLOW and DENY rules
  - rules only include IP addresses

> [!TIP]
> Route Tables: control the roads each neighborhood (subnet) uses to reach other places.
> Security Groups: control the doors and who can enter/exit each house (instance).
> Internet Gateway: acts like a main entrance to the city (VPC) that lets traffic in/out from the internet, but only if neighborhoods (subnets) build roads (routes) to it.

### ENI (Elastic Network Interface)

> manages networking for ec2 instance in a vpc

- network interface: virtual network card that connects ec2 instances and other aws resources to a vpc network
- allows you to create multiple network interfaces for an ec2 instance
- use cases
  - creating a management network interface
  - creating a network interface for traffic that needs to be isolated from other traffic

### vpc flow logs

> feature that captures information about the ip traffic going to and from network interfaces

- captures network information from aws managed interfaces like ELB, elasticache, rds, aurora, etc
- vpc flow logs data can go to S3, cloudwatch logs, kinesis data firehose
- what the vpc flow logs capture
  - source ip and destination ip
  - source port and destination port
  - protocol (tcp, udp)
  - traffic acceptance (accept or reject)
  - bytes transferred

### endpoints

> privately connect your VPC to aws services without using public IPs or going through the internet

- a VPC endpoint provides better security and lower latency to access aws services
- use VPC endpoint gateway if you want to connect to s3 or [DynamoDB](#dynamodb)
- use VPC endpoint interface if you want to connect to other aws services

### PrivateLink

> privately connect to a service in a 3rd party vpc

- use to expose your own service (like an API) to:
  - other VPCs in your organization
  - external AWS accounts
  - [AWS Marketplace](#marketplace)
- most secure and scalable way to expose a service to 1000s of vpcs
- doesn't require
  - vpc peering
  - internet gateway
  - nat
  - route tables
- requires
  - network load balancer
  - [ENI](#eni-elastic-network-interface)

### hybrid connectivity

- aws vpn: establishes an encrypted connection between your on-premises network and AWS over the public internet
  - client vpn: allows **individual** users to securely access AWS resources and on-premises private networks
    - the connection goes over the public internet, but the traffic is encrypted
    - the client device behaves as if it is part of the private network
  - site-to-site vpn: creates an encrypted connection between your on-premises network and your VPC over the public internet
- direct connect (DX): dedicated private network connection between your on-premises infrastructure and aws
  - bypasses the public internet
  - cannot connect 2 VPCs
  - reduces network latency
  - increases bandwidth throughput

### VPC Peering

> connect 2 VPCs privately using aws' network and make them behave as if they are the same network

- IP address ranges can't overlap
- less scalable than transit gateway

### transit gateway

> connect to multiple VPCs and/or on-premises infrastructure through a centralized hub

- simplifies network architecture
- makes it easier to manage connectivity at scale
- works with everything (connect gateway, vpn connections, vpc, subnets, internet gateway, nat gateway/instances, nacl, security groups, vpc peering, elastic ip, vpc endpoints, PrivateLink, pc flow logs, site-to-site vpn, client vpn, direct connect, transit gateway)

## lambda

> serverless compute service that lets you run code in response to events

- short-lived functions: each invocation runs in an isolated environment with max duration of 15 minutes
- stateless: each function runs independently
- event-driven, examples of events are file uploads, http requests, database changes
- automatically scales based on number of events
- price based on number of calls and duration of compute time
- you must allocate memory to your lambda function, from 128 MB (default) to 10 GB
  - increasing ram also improves cpu and network
- supports: node.js, python, golang, java, rust, ruby, c# (.NET core and powershell)

### workflow

1. write a function in Python, Node.js, Java, Go, etc
1. deploy it to Lambda (manually, with the CLI or using IaC like Terraform)
1. configure a trigger (e.g. HTTP endpoint, S3 upload)
1. lambda runs your code when the event occurs
1. get billed per request and per compute time (in milliseconds)

use case example: resize images uploaded to an S3 bucket

1. upload triggers an S3 event
1. aws Lambda receives the event, processes the image (e.g. resizes it)
1. stores the output back to S3

## RDS (Relational Database Service)

> fully-managed relational database service

- supports multiple engines:
  - [postgresql](./postgresql.md)
  - [mysql](./mysql.md)
  - [aurora](#aurora)
  - [mariadb](./mariadb.md)
  - oracle
  - microsoft sql server

- you still need to manage scaling
- aws handles backups (enabled by default) and patching, but you can customize their settings
- read replica: copy of a database that improves performance by offloading read operations from the primary database
  - doesn't contribute to high availability, since they are all located in a single AZ
  - improves database scalability
- multi-AZ: failover in case of AZ outage (high availability)
- multi-region (uses read replicas): spans multiple aws regions (e.g. `sa-east-1` and `us-east-1`)
  - disaster recovery in case of region issue
  - how it works:
    - Your main DB lives in one region (e.g. `eu-west-1` in Ireland)
    - you create read replicas in other regions (e.g. U.S. or Asia)
    - local apps can read from these replicas, reducing latency
    - writes still go to the main DB only

## S3 (Simple Storage Service)

> scalable object storage service that allows you to store and retrieve data from anywhere on the web

- use cases
  - object storage, **NOT** for database storage
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
  - event-driven: can trigger lambda, SQS, SNS on object creation or deletion
  - durable: 11 9s (99.999999999%) durability

- versioning (not enabled by default): option to retain multiple versions of objects
  - is **not** free, you are charged for every version of every object stored
- doesn't encrypt data at rest by default, aws provides encryption tools, but customers must use and configure them properly
- server-side encryption is enabled by default, client-side encryption is the customer's responsibility
  - client-side encryption: data is encrypted on the client (user's) side before it is sent to the server
  - server-side encryption: data is encrypted by the server after it is uploaded and decrypted by the server before it is served back to the client
- S3 Transfer Acceleration: enables fast, secure transfers of files over long distances between your client and S3 bucket
  - leverages cloudfront's globally distributed edge locations
- S3 data migration: process of moving data into, out of or between S3 buckets
  - use cases: large data cloud migrations, decommission (means data center retirement), disaster recovery
- there are no real folders
- max object size is 5 TB, if a file is bigger than 5 TB, upload is segmented into multiple parts

- key: unique identifier for an object within a bucket (often a filename or path)
  - must be unique in the bucket (e.g. `documents/2025/invoices/invoice_001.pdf`)
  - similar to full file path
    - aren't directories, just long names that contain slashes ("/")
  - prefix: path without file name
- object (similar to file): stored data (file + metadata + unique key)
- url structure to an s3 object: `https://<bucket-name>.s3.<region>.amazonaws.com/<object-key>`
  - `https://my-unique-bucket.s3.us-east-1.amazonaws.com/documents/2025/invoices/invoice_001.pdf`
- bucket (similar to directory): container that stores objects
  - must have unique name globally across all existing bucket names in S3 (not just your account)
  - created in a region (bucket location), which affects latency
  - pay attention to naming requirements
    - 3 to 63 characters
    - no spaces or underscores
    - no leading or trailing dots or hyphens
    - names can only contain lowercase letters, numbers, periods and hyphens

- bucket policies examples:
  - you can enable public access with an S3 Bucket Policy
  - to give user access to S3, use IAM policy to give access
  - to give access to an EC2 instance, always use EC2 instance roles with IAM permissions
  - to give access to IAM User in another AWS account, create a bucket policy that allows access to that specific user

> [!WARNING]
> to fix 403 error, give public access to your bucket

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

### shared responsibility

- aws is responsible for
  - infrastructure
    - global security
    - durability
    - availability
    - sustain concurrent loss of data in 2 facilities
  - configuration and vulnerability analysis
  - compliance validation
- me
  - access and permissions
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
  - no retrieval charges for accessing your data
- standard IA (Infrequent Access)
  - pay to retrieve data
  - 99% availability
  - use cases: disaster recovery, backups
- one zone-IA: for data that is accessed less frequently but requires rapid access when needed
  - high durability in a single AZ
  - cheaper than standard IA
  - less resiliant than standard IA (no AZ redundancy)
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
      - lowest cost
      - retrieval options:
        - standard: 12 hours (medium cost)
        - bulk: under 48 hours (lowest cost)

### replication (CRR and SRR)

> automatically copy objects from one S3 bucket to another

- you must enable Versioning in source and destination buckets
- buckets can be in different aws accounts
- copying is asynchronous, i.e. data is not copied immediately, there might be a small delay after upload
- 2 types:
  - Cross-Region Replication (CRR): Replicates data to a bucket in a different AWS region
    - use cases: compliance, lower latency access, replication across accounts
  - Same-Region Replication (SRR): Replicates data to a bucket in the same AWS region
    - use cases: log aggregation, live replication between production and testing accounts

#### snow family

- snowcone
  - 8 TB of HDD storage
  - 14 TB of SSD storage
- snowball edge: move less than 10 PBs of data
  - storage optimized: 80 TB HDD
  - compute optimized: 42 TB HDD capacity + 28 TB NVMe SSD
- snowmobile: literally a truck
  - use when transfering more than 10 PB
  - each has 100 PB (100,000 TB) of capacity (1 PB = 1000 TBs)

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

## ALB (Application Load Balancer)

> modern layer 7 (HTTP/HTTPS/gRPC) load balancer

- classic [ELB](#elb-elastic-load-balancer) is legacy and supports basic TCP/HTTP traffic
- similar to [nginx](./nginx.md)

## Amazon Connect

> AI-powered application that provides contact center

- set up call centers or customer service systems

## Amazon Managed Blockchain (out-of-scope)

> fully-managed service that makes it easy to create, manage and scale blockchain networks

- transactions without the need for a trusted, central authority
- use cases
  - join public blockchain networks
  - create your own scalable private network
- compatible with hyperledger and ethereum frameworks

## Amazon Q

> generative AI assistant design to help users analyze data, write code and answer questions using natural language

- includes:
  - IDE assistant (Q Developer)
  - business insights tools (Q Business)
- integrates with [Amazon Connect](#amazon-connect)

## Amazon WorkSpaces

> provides virtual desktops to users that can be accessed from anywhere

## Amazon WorkSpaces Secure Browser

> offers secure access to web content without exposing the network and without full desktop

## AMI (Amazon Machine Image)

> a template that contains a software configuration (e.g. operating system, application server, applications) that is used to launch EC2 instances

> complete snapshot of a virtual machine

- **not** a global service, you cannot use an AMI from one region in another region directly
- an AMI is neither a docker image nor a container image
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

## Amplify

> build and deploy full-stack web and mobile application quickly

- features:
  - hosting dynamic and static web apps
  - authentication
  - graphql/rest apis
  - works with frameworks like react, angular, vue and mobile platforms (iOS/android)

## API Gateway

> fully-managed service that makes it easy for developers to create, publish, maintain, monitor and secure APIs at any scale

client (e.g. browser) <= REST API => API Gateway <= proxy requests => Lambda <= CRUD => DynamoDB

- scalable and serverless API
- supports
  - rest
  - websocket apis
  - security
  - user authentication
  - api throttling: limiting the number of api requests a client can make in a time period
  - api keys
  - monitoring

## Application Discovery Service

> plan migrations to aws by collecting information about on-premises data centers

- often used with [Migration Hub](#migration-hub) to centralize migration tracking

## Application Migration Service

> lift-and-shift (rehost) service that simplifies, expedites and reduces the cost of migrating application to aws

## AppStream 2.0

> fully-managed service that lets you stream desktop applications to users over the internet

## AppSync

> fully-managed service that simplifies developing graphql APIs

> store and sync data across mobile and web apps in real-time

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

## Audit Manager

> automate audit preparation for compliance with regulations and standards

- automates evidence collection
- continuously collects and organizes aws resource configurations and activity data

## Aurora

> proprietary relational database engine that is part of rds

- fully-managed
- highly available
- costs more than rds, but is more efficient
- supports postgresql and mysql
- continuous auto scaling

## aws artifact

> provides customers on-demand access to aws compliance documentation, agreements and audits

- doesn't do audits, only provides access to the results of audits done on aws

## aws backup

> fully-managed backup service that makes it easy to centralize and automate the backup of data across aws services

- integrates with AWS organizations to provide centralized backup management for all accounts within the organization
- can automate backups for:
  - EBS volumes
  - RDS databases
  - DynamoDB tables
  - EFS file systems
  - S3
  - AWS storage gateways

## aws certificate manager

> easily provision, manage and deploy SSL/TLS certificates

- SSL/TLS certificate provides encryption for websites (HTTPS)
- supports both public and private tls certificates
- free of charge for public tls certificates
- automatic tls certificate renewal
- integrates with ELB, cloudfront distributions, APIs on API gateway

## aws compliance

> helps customers meet security, regulatory and compliance requirements

- exists so AWS can demonstrate and publish how its infrastructure and services meet various regulatory and industry standards
  - helps AWS be compliant **for** the cloud
  - helps customers be compliant **in** the cloud
- while AWS infrastructure is compliant, customers must configure their services and applications properly to remain compliant
- achieves certifications and passes assessments to demonstrate compliance with global standards, such as:
  - ISO 27001
  - SOC 1, SOC 2, SOC 3
  - PCI DSS (for payment data)
  - HIPAA (for healthcare data)
  - GDPR (EU data protection)

## aws compute optimizer

> service that helps you optimize your aws resources by recommending the most cost-effective and performance-efficient options

- supports: ec2 instances, auto scaling groups, ebs volumes and lambda functions
- analyzes historical usage (e.g. CPU, memory, I/O)
- provides right-sizing recommendations

## aws config

> assess, audit and evaluate the configurations of your aws resources

- continuously monitors and records AWS resource configurations and evaluates them against desired configurations using rules
  - aws config rules: check if your aws resources comply with specific desired configurations or policies

- use cases
  - enforce encryption on RDS databases
  - ensure IAM policies don't allow wildcard actions
  - check whether all resources are in approved regions
  - make sure CloudTrail is enabled and logging correctly

## aws directory service

> helps organizations use Microsoft Active Directory (AD) with aws services

- directory services: stores information about users, groups, devices and other resources in the network
  - usually on on-premises systems
- Microsoft Active Directory: directory service for windows domain networks
  - windows domain network: type of computer network in that manages user accounts, computers and other resources
    - features
      - centralized authentication: users log in with one set of credentials across all computers in the domain
      - domain controller: special server that runs active directory and handles logins, access rights and security policies

## aws firewall manager

> allows you to **centrally** manage and configure firewall rules across aws accounts and resources

- organization-wide security policy enforcement
- helps you manage
  - WAF rules
  - aws Shield advanced protections
  - vpc security groups
  - route 53 resolver dns firewall rules
  - network firewalls

## aws iam identity center

> one login for multiple AWS accounts and applications

- manage access for workforce users (employees) across AWS accounts and business applications
- supports federated access
  - integrates with external identity providers (IdPs) like Microsoft Entra ID (Azure AD), Okta, etc (using SAML 2.0 or OIDC)
  - use case: enterprise users accessing AWS services through identity providers (e.g. Okta, Azure AD, Google Workspace)
- can integrate with STS to generate temporary security credentials for users
- one login for all your
  - aws accounts in aws organizations
  - business cloud applications (e.g. salesforce, box, microsoft 365)
  - ec2 windows instances

## aws knowledge center

- basically an official FAQ
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

## aws network firewall (out-of-scope)

> deploy essential network protections for your VPCs

> filter traffic entering and leaving a VPC

- VPC-level protection
- stateful firewall
- features
  - traffic filtering
  - inspect the content of packets, not just headers
  - logging and monitoring

## aws organizations

> centrally manage multiple aws accounts

- allows management of permissions
- AWS organizations Service Control Policies (SCPs): centrally manage and restrict permissions across all accounts
  - SCPs do not grant permissions, they define the maximum available permissions and can restrict actions, even if IAM policies allow them
  - does not apply to management account
  - applies to all the users and roles of the account, including root
- api is available to automate aws account creation

- Organizational Unit (OU): used to group accounts within the organization

## aws prescriptive guidance

> collection of strategies, guides and best practices to help customers plan, implement and scale their cloud adoption

## aws professional services and partner network

- aws professional services: global group of experts available to work alongside your team and help achieve cloud goals faster
  - often partner with APN (Aws Partner Network) members
- APN (Aws Partner Network): global community of 3rd party companies that help customers build, market and sell their offerings on aws
  - offerings = products, services or solutions that APN provides to customers using aws infrastructure
  - types of APN partners
    - APN technology partners: provide hardware, connectivity or software solutions
    - APN consulting partners: help customers design, architect, build, migrate and manage workloads on aws
    - APN training partners: deliver aws-authored training to organizations and individuals
  - aws competency program: recognize APN partners that have demonstrated technical expertise
  - aws navigate program: helps APN partners improve specialized skills
- APN Portal: provides partners with resources, tools and content to support their business growth
- Partner Volume Discounts: reduced pricing based on the volume of aws services that an APN partner manages

### aws partner training and certification

> help individuals/organizations build cloud skills through online courses and credentials

## aws secrets manager

> fully-managed service that stores secrets

- can force rotation of secrets after a period of time
- automatic rotation of secrets (uses lambda)
- integrates with [rds](#rds-relational-database-service)
- secrets are encrypted using [KMS](#kms-key-management-service)

## aws support plans

- basic/free: available to all aws customers by default
  - offers access to aws documentation and discussion forums (AWS Support Forums)
- developer: cheapest paid tier for testing
  - offers access to aws documentation and discussion forums (AWS Support Forums)
- business: for production systems
  - AWS Support API: allows you to integrate aws support into your applications
    - allows you to programmatically interact with your aws support cases
- enterprise: mission-critical + TAM (Technical Account Manager) + Concierge Support (billing and account assistance)

| Plan       | Cost (Starting)     | Access Type          | Use Case         | Response Time (Critical) | Trusted Advisor |
| ---------- | ------------------- | -------------------- | ---------------- | ------------------------ | --------------- |
| Basic      | Free                | AWS Docs & Forums    | Learning/Test    | Not available            | Limited access  |
| Developer  | $29/mo or 3%        | Email (Business hrs) | Dev/Test         | < 12–24 hrs              | Limited access  |
| Business   | $100/mo or % usage  | 24/7 Chat/Phone      | Production       | < 1 hr                   | Full access     |
| Enterprise | $15k/mo or % usage  | TAM + 24/7 support   | Mission-Critical | < 15 mins                | Full access     |

## aws well-architected framework

> set of best practices and guidelines created by amazon to help cloud architects design and operate secure, resilient and efficient infrastructure for their applications

6 pillars:

1. operational excellence
2. security
3. reliability
4. performance efficiency
5. cost optimization
6. sustainability

- [well-architected tool](#well-architected-tool): free tool to review your architectures against the 6 pillars of well-architected framework
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

### security

> protect information, systems and assets

- identity and access management
- data protection (at rest and in transit)
- incident response
- threat detection

### reliability

> ensure a system can recover from failures and meet customer demands

- automated recovery
- failure management (e.g. backups)
- distributed system design
- load balancing
- scalability

### performance efficiency

> use computing resources efficiently to meet system requirements as demand changes

- scalability
- use serverless and managed services where possible
- monitor and improve performance
- test different instance types and configurations

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

### sustainability

- optimize energy consumption
- improve efficiency across system lifecycle
- use managed services to reduce waste

## Batch

> fully-managed batch processing at any scale

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
| Platform    | Cloud architecture/DevOps     | Architects, engineers |
| Security    | Data and asset protection     | Security teams        |
| Operations  | Manage/monitor cloud services | IT ops, admins        |

- CAF Transformation Phases:
  - Envision: demonstrate how the cloud will accelerate business outcomes
    - identify transformation opportunities
    - create a foundation for your digital transformation
  - Align: identify capability gaps across the 6 AWS CAF Perspectives
    - results in an action plan
  - Launch: build and deliver pilot initiatives in production
    - demonstrate incremental business value
  - Scale: expand pilot initiatives to the desired scale while realizing the desired business benefits

## CDK (Cloud Development Kit) (out-of-scope)

> define cloud infrastructure using familiar programming languages

- compatible with
  - javascript/typescript
  - python
  - java
  - .NET
- code is compiled into CloudFormation template (json/yaml)

## CloudFormation

> fully-managed service that allows you to create and manage AWS resources using templates

- IaC, similar to terraform
- template: configuration file that defines infrastructure
- provision the same aws infrastructure across multiple AWS accounts and regions
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
  - origin server: original source of the content that CDN retrieves when the content is not already cached at the edge
- AWS Edge Locations: cache copies of content close to users to reduce latency and improve performance
  - part of cloudfront's CDN
- PoP (Points of Presence): edge locations used by cloudfront
- improves content delivery by caching the content at edge locations
- scales automatically to accommodate traffic spikes
- does **not** directly improve database performance
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
- CloudHSM keys
  - the customer has full control over key management
  - required for highly regulated environments

## CloudTrail

> get history of events or API calls made within your aws account via aws cli, aws console, aws sdk or aws services

- features:
  - governance: monitoring account activity to ensure **compliance**
  - auditing
  - API call logging:
    - view history of AWS API calls across your account
    - log data events such as S3 object-level API activity
  - IAM integration that provides detailed information on who performed the actions
- enabled by default
- to store logs long-term, put logs from CloudTrail into [CloudWatch Logs](#cloudwatch-logs) or S3
- cloudtrail insights: automates the analysis of cloudtrail events to detect unusual activity

> [!TIP]
> cloudtrAIl => ApI

## CloudWatch

> real-time system health and performance monitoring

- monitors resources, applications, performance and operational health
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

#### CloudWatch Logs for EC2

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

## Cloud9 (out-of-scope)

> cloud IDE that can be used in the browser

- allows code collaboration in real-time

## CodeArtifact (out-of-scope)

> fully-managed artifact repository service

- use cases:
  - publish your executable/binary in common dependency management tools (e.g. npm, yarn, pip, maven, gradle)
  - manage versioned software artifacts

- artifact: package or binary like `.jar`, `tgz` (tarball)

## CodeBuild

> fully-managed build service that compiles source code, runs tests and produces software packages that are ready to deploy

- similar to github actions and gitlab ci
- fully-managed, serverless
- pay-as-you-go pricing
- can be triggered by [CodePipeline](#codepipeline)
- use cases
  - compile source code
  - run tests
  - produce packages that are ready to be deployed (by CodeDeploy, for example)
- integrates with github, codecommit and bitbucket

## CodeCommit (out-of-scope)

> fully-managed source control service that makes it easy for teams to host secure and highly scalable private Git repositories

- aws' github competitor
- source control service
- fully-managed

## CodeDeploy (out-of-scope)

> fully-managed deployment service that automates software deployments to a variety of compute services

- works with
  - ec2 instance
  - lambda
  - on-premise servers
- hybrid service
- servers/instances must be provisioned and configured ahead of time with CodeDeploy agent

## CodeGuru (out-of-scope)

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

- fully-managed
- can activate [codebuild](#codebuild)
- compatible with:
  - [CodeCommit](#codecommit)
  - [CodeBuild](#codebuild)
  - [CodeDeploy](#codedeploy)
  - [elastic beanstalk](#elastic-beanstalk)
  - [CloudFormation](#cloudformation)
  - [github](./github.md)
  - 3rd party services
  - custom plugins

codecommit => codebuild => codedeploy => compute resource (can be ec2 instance, ecs, etc)

## CodeStar (out-of-scope)

> unified user interface to easily manage software development activities

- quick way to setup CodeCommit, CodePipeline, CodeBuild, CodeDeploy, Elastic Beanstalk, EC2, etc
- edit the code "in the cloud" using AWS Cloud9

## Cognito

> fully-managed identity service that allows you to add user sign-up, sign-in and access control to your apps quickly and easily

- authenticate end-users of web or mobile applications
- similar to auth0 and firebase
- supports encryption of data at rest and while it's in transit
- identity for web/mobile application's users
- don't create an IAM user for the clients of your application, use Cognito
- instead of giving your app's users aws iam accounts (which are meant for admins and systems), use Cognito to manage their identities securely
- supports federated access: sign in with google/facebook accounts

## Comprehend

> uses Natural Language Processing (NLP) to extract insights about the content of documents

- fully-managed and serverless service
- uses machine learning to find insights and relationships in text
- capabilities
  - analyzes text using tokenization
  - organize a collection of text files by topic
  - analyze customer interactions as positive or negative

## Consolidated Billing

- when enabled, combine billing for multiple aws accounts into one

- no extra cost to use
- get discounts:
  - volume discounts applied across accounts (e.g. data transfer, s3 storage)
  - RI discount sharing: AWS organizations share Reserved Instance discounts across all accounts in the organization
  - savings plans discounts

## Control Tower

> automatically set up aws organizations to organize accounts and implements SCPs (Service Control Policies)

- helps manage multiple AWS accounts, not individual IAM users
- benefits
  - automate the setup of multiple aws accounts in a few clicks
  - automatically apply recommended security, compliance and operational settings
  - automate ongoing policy management using guardrails

## Cost Allocation Tags

> organize and track aws costs by assigning metadata to your aws resources

- how it works: you apply tags (key-value pairs) to resources, like `Environment=Production` or `Team=Finance`
- types of tags: aws-generated and user-defined

## Cost Anomaly Detection

> continuously monitor your cost and usage using ML to detect unusual spending

## Detective

> investigates and identifies the root cause of security issues or suspicious activities using ML and graphs

- generates visualizations with details and context to get to the root cause
- helps analyzing findings from [GuardDuty](#guardduty) and other services

## DMS (Database Migration Service)

> fully-managed service that helps you migrate databases to aws quickly and securely

- primary function is migrating data, NOT the conversion of database schemas, use [SCT](#sct-schema-convertion-tool) for that
- resilient and self healing
- source database remains available during the migration
- supports
  - homogeneous migration (e.g. oracle to oracle)
  - heterogeneous migration (e.g. sql server to aurora)

## DocumentDB

> fully-managed aws implementation of mongodb

- nosql database
- store, query and index json data
- highly available with replication across 3 AZ

## DRS (Disaster Recovery Service or Elastic Disaster Recovery)

> recover your applications and systems quickly after a disaster

- provides disaster recovery for on-premises and cloud-based applications by continuously replicating your servers to aws
- enables fail over to aws and fail back once systems are restored
- recovery time is faster than [AWS Backup](#aws-backup), and focused on servers

## DynamoDB

> fully-managed NoSQL database service with automatic scalability

- TIP: dYnamodb => keY-value database
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

- global tables: replicate data automatically across your choice of AWS Regions
  - automatically scale capacity to accommodate your workloads
  - globally distributed applications can access data locally in selected regions to get single-digit millisecond read/write performance

### DynamoDB streams

> change log for table

- captures real-time changes: inserts, updates, deletes
- it's optional

### DAX (DynamoDB Accelerator)

> fully-managed in-memory cache for DynamoDB

- 10x performance improvement
- secure, highly scalable and highly available

## ECS (Elastic Container Service)

> fully-managed container orchestration service that allows you to run, stop and manage Docker containers on a cluster of EC2 instances

- launch containers with ec2 instances or [fargate](#fargate)
- use ECS Service Auto Scaling for elasticity (not enabled by default)

| Feature           | ECS on EC2                       | ECS with Fargate                  |
| ----------------- | -------------------------------- | --------------------------------- |
| Server Management | You manage EC2 instances         | No server management (serverless) |
| Scalability       | You scale EC2 cluster manually   | Fargate scales automatically      |
| Pricing           | Pay for EC2 instances            | Pay per vCPU and memory used      |
| Flexibility       | More control over infrastructure | Simpler, less overhead            |

## ECR (Elastic Container Registry)

> fully-managed Docker container registry that makes it easy to store, manage and deploy Docker container images

- private docker registry
- stored docker images can be run by [ECS](#ecs-elastic-container-service) or [fargate](#fargate)

## EFS (Elastic File System)

> fully-managed cloud-based file storage service

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

> fully-managed Kubernetes service that allows you to run, manage and scale containerized applications using Kubernetes

## ElastiCache

> fully-managed service that deploys, operates and scales popular open-source compatible in-memory data stores in the cloud

- used to manage [redis](./redis.md) or memcached
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

> legacy service that automatically distributes incoming application traffic across multiple targets

- similar to [nginx](./nginx.md)
- targets can be ec2 instances, ip addresses, containers
- improves fault tolerance
- can have high availability (multi AZ)
  - multi-AZ ELB: can distribute traffic across multiple AZs
  - single-AZ ELB: can only send traffic to its own AZ
- improves scalability
- does regular health checks to your instances
- types of load balancers
  - [ALB](#alb-application-load-balancer) (network layer): uses HTTP/HTTPS/gRPC protocols
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

> fully-managed service that allows you to deploy and manage web applications and services

- PaaS
- developer is only responsible for the application
- elastic beanstalk uses CloudFormation to provision the application, but you don't need to managed it (all automated)
- managed service
  - instance configuration and OS is handled by Beanstalk
  - deployment strategy is configurable but managed by Beanstalk
- automatically handles: provisioning (e.g. ec2 instances), load balancing, auto scaling and monitoring
- 3 architecture models:
  - single instance deployment
  - LB + ASG: good for production or pre-production web applications
  - ASG only: good for non-web apps in productions (workers, etc)

- support for:
  - golang
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

> task scheduler

- uses events to connect application components together
  - event: json message that describes something that happened in a system
    - events come from aws services, custom apps or SaaS providers (e.g. auth0)
- event bus: pipeline where events are sent and rules are applied to process them
- event source: origin of the event
- rules: match certain events and define what should happen when a match occurs
- targets: aws services or resources that receive matched events and act on them
  - common targets: lambda (run code), step functions (start workflow), sqs (enqueue message), sns (publish to topic), etc

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

> fully-managed serverless compute engine for running containers on AWS

- allocates the exact cpu and ram requested
- provides compute runtime for [ECS](#ecs-elastic-container-service) and [EKS](#eks-elastic-kubernetes-service)
- removes the need to provision and manage ec2 instances

## FIS (Fault Injection Simulator) (out-of-scope)

> test the resilience of your applications (chaos engineering)

- similar to netflix's chaos monkey
- supports EC2, ECS, EKS, RDS, etc

## Forecast (out-of-scope)

> fully-managed service that uses ML to deliver highly accurate forecasts

- example scenario: predict the future sales of a raincoat
- use cases
  - product demand planning
  - financial planning
  - resource planning

## FSx

> fully-managed service that allows you to launch high-performance file systems in the cloud optimized for windows-based applications

- amazon FSx for Windows File Server: for windows-based applications, Active Directory and SMB protocol
  - built on Microsoft Windows Server
- amazon FSx for Lustre (out-of-scope): used for machine learning, analytics, video processing, financial modeling

## Global Accelerator

> improve application availability and performance using the aws global network

- offers static IP addresses (doesn't change over time)

[cloudfront](#cloudfront) vs global accelerator:

| Feature       | Amazon CloudFront                                 | AWS Global Accelerator                                 |
| ------------- | ------------------------------------------------- | ------------------------------------------------------ |
| Purpose       | Content delivery (CDN)                            | Improves global application performance & availability |
| Optimized for | Static and dynamic content (e.g., images, videos) | Any type of application (web apps, APIs, etc.)         |
| Routing Type  | HTTP/HTTPS only                                   | TCP and UDP traffic                                    |
| IP Type       | Uses domain name (e.g., `d123.cloudfront.net`)    | Uses static IP addresses                               |
| Best For      | Fast content delivery to end users worldwide      | Fast and reliable routing to regional AWS endpoints    |

## Glue

> managed Extract, Transform, Load (ETL) service

- "glue" data together from different sources
- prepare and transform data for analytics
- fully serverless
- glue data catalog: catalog of datasets
- glue databrew: tool that enables users to clean and normalize data without writing any code
- can be used by [athena](#athena), [redshift](#redshift), [EMR](#emr-elastic-mapreduce)

## GuardDuty

> threat detection service that continuously monitors aws data sources

- uses machine learning algorithms and anomaly detection
- analyses the following data sources
  - cloudtrail events logs
  - vpc flow logs
  - dns logs

## Health Dashboard (Your Account – Personalized View)

> provides personalized alerts and remediation guidance for aws outages that might impact your aws account or resources

- global service
- provides
  - alerts and remediation guidance when AWS is experiencing events that may impact you
  - personalized view into the performance and availability of AWS services you are using
  - relevant or timely information to help you manage events in progress
  - proactive notification to help you plan for scheduled activities
  - aggregate data from entire AWS Organization
  - subscription to an RSS feed to be notified of the status of all aws service interruptions

### Service Health Dashboard (Public View)

- shows
  - all regions
  - all services health
  - historical information for each day

## Infrastructure Composer (out-of-scope)

> no-code tool that helps you design and deploy aws infrastructure using cloudformation

- uses cloudformation templates
- ideal for begginers that are unfamiliar with writing YAML/JSON templates manually

## Inspector

> find software vulnerabilities in compute resources

- scans aws workloads for vulnerabilities and unintended network exposure
- inspects running OS against known vulnerabilities
- only for ec2 instances, ecr container images and lambda functions
- not used for tracking user actions or api calls

## IoT core

> managed service that lets connected IoT devices interact with aws services and other devices securely

- enables IoT (Internet of Things) devices to connect, send data and receive commands using aws
- offers secure communication
- register, manage and monitor IoT devices
- route data to aws services (e.g. lambda, s3, dynamodb) automatically

## Kendra

> fully-managed document search service powered by ML

- extract answers from within a document (text, pdf, html, powerpoint, etc)
- natural language search capabilities
- can learn from user interactions/feedback to promote preferred results
- ability to manually fine-tune search results

## Kinesis

> fully-managed service that allows you to collect, process and analyze real-time streaming data

- streaming data: data that arrives continuously and in small chunks
- unlimited scalability
- cloud-native service: proprietary protocol from aws

## KMS (Key Management Service)

> managed service that allows you to create and control the encryption keys used to encrypt your data

- keys can be managed by customers or aws
- customer is responsible for controlling who has access to encrypted data
- uses CMKs (Customer Master Key): primary resource in KMS used to encrypt and decrypt data
  - customer-managed CMKs: created, managed and used by the customer
  - aws-managed CMKs: created, managed and used by aws
  - aws-owned CMKs: collection of CMKs that aws services own and managed, used across multiple accounts

### workflow

you can use aws management console or aws cli

1. you create a customer master key (CMK) in KMS
1. you use the CMK to encrypt your data
1. when you need to decrypt your data, you use the same CMK

## Lex

> add chatbots that speak natural language to applications

- uses Automatic Speech Recognition (ASR) to convert speech to text
- uses natural language understanding to recognize the intent of text, callers

## License Manager

> helps you manage software licenses from vendors like Microsoft, Oracle, IBM across aws and on-premises environments

## Lightsail

> deploy and manage applications/websites with pre-configured cloud resources

- great for people with little cloud experience
- simpler alternative to using ec2, rds, elb, ebs, route 53...
  - sets up a virtual server, database, storage and networking
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

> fully-managed data security service that uses machine learning and pattern matching to discover and protect sensitive data

- helps identify and alert you to sensitive data, such as Personally Identifiable Information (PII) in s3 buckets
- uses s3 buckets as data source

## Marketplace

> digital catalog where customer can find, buy and deploy third-party software, services and data that run on aws

- sell SaaS solutions to aws customers
- buy software that has been bundled into customized AMIs by the aws marketplace sellers
- provides cost, governance and entitlement management
  - entitlement: manage who has access to software licenses

## Migration Evaluator

> assess cost and feasibility of migrating to aws

- helps estimating cost savings, right-sizing opportunities and optimal aws services for your workloads
- gathers data about your existing infrastructure and analyzes it

## Migration Hub

> provides central dashboard to track and manage application migrations to aws

- tracks status of servers and applications as they move to aws
- no extra charge for using migration hub

## MQ

> managed message broker service for rabbitmq and activemq

- doesn't scale as much as [sqs](#sqs-simple-queue-service) and [sns](#sns-simple-notification-service)
- **not** serverless

## Neptune

> fully-managed graph database

- high availability
- replication across 3 AZ enabled by default
- great for knowledge graphs

## OpenSearch Service

> managed service that searches, analyzes and visualizes large amounts of data

- used for log analytics, monitoring, search and real-life data analysis
- keyword-based search, unlike [Kendra](#kendra)

## OpsWorks (out-of-scope)

> manages chef and puppet

- chef and puppet are tools that perform server configuration automatically
- alternative to [SSM](#ssm-systems-manager)

## Outposts

> fully-managed service that extends aws infrastructure, services and tools to virtually any on-premises or edge locations

> run aws services locally on your own hardware, while still being managed from the aws cloud

- you are responsible for the outposts rack physical security
- outposts require an active connection to an aws region for service control plane operations (like management and monitoring)
- benefits
  - low latency access to on-premise systems
  - local data processing
  - data residency
  - easier migration from on-premises to the cloud
  - fully-managed service
- compatible services: ec2, ebs, s3, eks, ecs, rds, emr

## Personalize (out-of-scope)

> fully-managed ML service to build apps with real-time personalized recommendations

- same technology used by amazon
- integrates with
  - existing websites
  - applications
  - sms
  - email marketing systems
- implement in days, not months
- no need to build, train and deploy ML solutions
- use cases: retail stores, media, entertainment, etc

## Pinpoint (out-of-scope)

> customer engagement service that allows businesses to communicate with their customers through multiple channels

- supports email, SMS, push notifications (mobile apps), voice messages, in-app messaging
- capable of receiving replies
- use cases: run campaigns by sending marketing, bulk, transactional SMS messages

- what is the difference when comparing with SNS or SES?
  - in SNS and SES, you managed each message's audience, content and delivery schedule
  - in Pinpoint, you create
    - message templates
    - delivery schedules
    - highly-targeted segments
    - full campaigns

## Polly

> fully-managed service that converts text to speech using deep learning

- allows development of applications that talk

## QLDB (Quantum Ledger DataBase) (out-of-scope)

> fully-managed ledger database

> track all changes to your application data over time and maintain a verifiable history of changes

- serverless
- centralized database, with one source of truth
- high availability
- replication across 3 AZ
- immutable system: no entry can be removed or modified, cryptographically verifiable

## QuickSight

> serverless machine learning-powered business intelligence (BI) service that provides visual tools for data analytics

- per-session pricing
- use cases
  - create and publish interactive dashboards
  - visualize data form various sources (S3, RDS, redshift, Aurora, Athena, etc)
  - business analytics
  - perform ad-hoc analysis

## RAM (Resource Access Manager)

> allows you to grant access to the resources you own to other accounts

## Redshift

> fully-managed data warehouse service that allows you to run complex queries on large datasets

- based on postgresql, but it's not used for OLTP
- used for [OLAP](/business_intelligence.md#olap-online-analytical-processing) (analytics and data warehouse)
- requires a well-defined schema
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

### registering domain

visitor enters URL → Domain Registrar → Nameservers → Hosted Zone → DNS Records → IP Address

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

## SageMaker AI

> fully-managed service for developers to train machine learning models that can make predictions

lets say you want to build a model that predicts your exam score

- SageMaker can help in all steps below and do more:
  - gather past data
  - label with a score (0 to 10 in the exam)
  - build a ML model
  - train and tune ML model
  - SageMaker now make predictions

## SCT (Schema Convertion Tool)

> help migrate databases from one database engine to another

- e.g. oracle to Aurora, sql server to postgresql
- often used with aws DMS (Database Migration Service)
- not needed when the source and target databases are the same engine (homogeneous migration)
- SCT vs [DMS](#dms-database-migration-service)
  - DMS migrates data between databases
  - SCT
    - converts database schema from one engine to another
    - doesn't move data, only schema and sql code conversion

## Security Blog

> official source of security guidance, best practices and service updates from aws

## Security Center

> official resource for security and compliance information from aws

## Security Hub

> centralized security monitoring that makes compliance checks

- helps identify security issues across aws accounts and services
- continuously monitors security best practices and compliance standards
- aggregates alerts from various aws services across multiple AWS accounts
  - config, GuardDuty, inspector, macie, iam access analyzer, aws firewall manager, aws health, aws partner network solutions

### CSPM (Cloud Security Posture Management)

> subset of security hub capabilities

- audit configurations
- run compliance checks

Security Hub CSPM vs [Detective](#detective) vs [guardduty:](#guardduty)

| Feature                     | **Security Hub (CSPM)**                        | **Amazon Detective**                        | **Amazon GuardDuty**                         |
| --------------------------- | ---------------------------------------------- | ------------------------------------------- | -------------------------------------------- |
| Purpose                     | Monitors security posture and compliance       | Investigates and analyzes security findings | Detects threats and suspicious activity      |
| Does it detect threats?     | ❌ No (relies on other services)               | ❌ No (investigation only)                  | ✅ Yes (main purpose)                        |
| Does it analyze findings?   | ❌ Not deeply                                  | ✅ Yes (deep analysis & root cause)         | ❌ No (just alerts)                          |
| Does it aggregate findings? | ✅ Yes (e.g. from GuardDuty, Macie, Inspector) | ❌ No (uses findings from GuardDuty)        | ❌ No                                        |
| Compliance checks?          | ✅ Yes (e.g. CIS, PCI-DSS)                     | ❌ No                                       | ❌ No                                        |
| Data sources used           | From other services                            | CloudTrail, VPC Flow Logs, GuardDuty        | CloudTrail, VPC Flow Logs, DNS logs          |
| Typical use                 | Posture & compliance dashboard                 | Investigate a GuardDuty alert               | Alert on suspicious API calls or IP behavior |

## Service Catalog

> allows organizations to manage/distribute approved resources (can be applications or aws resources)

- typically used by IT administrators to control what AWS resources and services end-users (like developers) can deploy
- permissions management through IAM roles and policies
- ensures users can only deploy compliant and pre-approved configurations

## Service Quotas

> notifies you when you are close to a service quota value threshold

## SES (Simple Email Service)

> fully-managed service for sending emails

## Shield

> protects networks and applications by analyzing network security configurations and providing managed DDoS protection

- shield standard: free protection from DDoS attacks, SYN/UDP floods, reflection attacks
- shield advanced: paid 24/7 DDoS protection
  - 24/7 access to aws DDoS response team and detailed attack diagnostics

## SNS (Simple Notification Service)

> fully-managed messaging service that allows you to send notifications to a large number of recipients

- unlimited scalability
- cloud-native service: proprietary protocol from aws
- no message retention, unlimited queues

- [pub-sub](/system_design.md#pubsub-architecture) paradigm
- event publishers: sends message to one SNS topic
- event subscribers: listens to the sns topic notifications
  - each subscriber to a topic will get all the messages
  - services that can be target subscribers: sqs, lambda, kinesis data firehose, emails, sms, mobile notifications, http(s) endpoints

## SQS (Simple Queue Service)

> message queuing service that allows you to decouple and scale microservices, distributed systems and serverless applications

- fully-managed
- unlimited scalability
- message: unit of data that is sent from one component of a system to another through a queue
  - messages are kept for 14 days
  - messages can be xml or json

- producer service: produces requests to sqs
- consumer service: consumers requests from sqs
- decoupling: independent scaling of apps and isolated failing of apps

## SSM (SystemS Manager)

> centrally view, manage and operate nodes (ec2 or on-premises) at scale in AWS, on-premises and multicloud environments

> provides you with browser-based shell and CLI experience

- fully-managed, hybrid service
- similar to [ansible](./ansible.md)
- not just a single service, it is a suite of products
- uses IAM for access control
- features
  - automate common administrative tasks: running commands, patching (for compliance) and configuring servers without using SSH
  - quickly identify any issues that might impact applications using a specific resource by checking:
    - API activity
    - resource configuration changes
    - related notifications
    - operational alerts
    - software inventory
  - SSM Parameter Store: store configuration data and secrets

### workflow

1. install SSM agent onto the systems we control
1. its installed by default on Amazon Linux AMI and Ubuntu AMI
1. if an instance can't be controlled with SSM, it's probably an issue with SSM agent

## Step Functions

> serverless orchestration service that lets you coordinate multiple AWS services into automated workflows (aka state machines)

- runs tasks in sequence or parallel
- automatically handles retries, timeouts, and error handling
- automate business processes, data processing pipelines or application workflows
- integrates with services like Lambda, ECS, SQS, DynamoDB, etc

## Storage Gateway

> extends storage of on-premises environments to aws cloud storage, helping with backup, archiving and disaster recovery

- hybrid cloud storage service, bridge between on-premise data and cloud data in S3
- data encryption enabled by default
- works with s3, glacier and ebs
- use cases
  - back up and archive on-premises data to aws
  - provide low-latency access to frequently used data via local caching
  - migrate data to aws over time
  - extend on-premises storage without buying new hardware
- types of gateways
  - file gateway: stores files as objects in S3
  - tape gateway
  - volume gateway

## STS (Security Token Service) (out-of-scope)

> create temporary, short-term credentials to access your aws resources with limited privileges

## Textract

> extracts text, handwriting and data from any scanned documents using AI and ML

- read and process any type of document (PDFs, images, etc)

## Translate

> natural and accurate language translation of text using ML

## Trascribe

> fully-managed service that automatically converts speech to text using ML

- uses deep learning process called Automatic Speech Recognition
- automatically removes Personally Identifiable Information (PII) using redaction

## Trusted Advisor

> analyzes your aws accounts with predefined rules and provides recommendations

- provides recommendations on 5 categories:
  - cost optimization (e.g. identity unattached or underutilized EBS elastic volumes)
  - performance
  - security
  - fault tolerance
  - service limits

- [AWS Support Plan](#aws-support-plans) benefits:
  - Basic: limited checks (e.g. service limits and security)
  - Business/Enterprise: full set of checks and recommendations

## WAF (Web Application Firewall)

> create security rules that control bot traffic and block common attack patterns

> protects web application from common web exploits (layer 7)

- deploy on [ALB](#alb-application-load-balancer), [API gateway](#api-gateway), cloudfront and [AppSync](#appsync)
  - ALB is part of [ELB](#elb-elastic-load-balancer)
- you are responsible for
  - setting WAF rules and conditions that match your security requirements

## Well-Architected Tool

[Well-Architected Framework](#aws-well-architected-framework)

> free tool to review your architectures against the 6 pillars of the Well-Architected Framework and adopt best practices

1. select your workload and answer questions
1. review your answers against the 6 pillars
1. obtain advice

## X-ray

> provides visual analysis of application services by collecting data about requests your application serves, helping developers gain insights into performance and behavior, especially in microservices or distributed systems

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

## migration strategies

- active-active: multiple environments (regions or data centers) are actively running and serving traffic at the same time
- active-passive: one environment is active, the other is on standby (only takes over if the active one fails)
- big bang migration: all systems switch to AWS at once, often during a scheduled downtime window
- phased migration (trickle or hybrid)
- blue/green deployment

## pricing

- always free AWS services: Security Groups, Auto Scaling, CloudFormation
  - IAM
  - VPC subnets
  - Budgets
  - Lambda (1M requests/month)
  - DynamoDB (25 GB storage)
  - CloudWatch (10 custom metrics and alarms)
  - Amazon Q Developer (Unlimited autocomplete suggestions)
  - Aurora DSQL (100k DPUs)
  - CloudFront (1 TB of data trasfer out)
  - aws blogs, forums, guides, quick starts

- 12-months free AWS services for new AWS customers:
  - EC2, S3, RDS, EBS, ECR, CloudFront, ElastiCache, DynamoDB*, Glacier*, Lambda*

- What AWS services are never free (need to pay to use): WAF, Inspector, Route 53, EBS volumes, ELB (WIREE)

- types of pricing in AWS
  - pay for compute time: EC2, Lambda, Fargate
  - pay for data stored in the cloud: S3, EBS, RDS
  - pay only for data transfer out of the cloud
    - data transfer into aws is free

### models

- pay as you go
- pay per use
- pay less by using more
- pay less when you reserve
- pay less when AWS grows
- no up-front investment

### data transfer

- inbound (data going into aws services): free
- outbound (data going from aws to the internet): first 100 GB/month is free and after that, it's charged

- within the same Availability Zone via private IP: free
- between different AZs (within same region):
  - private IP: ~$0.01/GB (approximate; actual pricing may vary)
  - public IP or elastic IP: ~$0.02/GB (also approximate)
- between regions: charged

> [!NOTE]
> Use private IPs and same AZ for lower costs and better network performance.
> Using different AZs increases availability but adds networking costs.

### cloud licensing strategies

- on-premises: buy long-term licenses
- BYOL (Bring Your Own License): use existing licenses in the cloud
- license included: aws bundles license cost into usage fee (e.g. windows/sql server)

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

### disaster recovery strategies

cost comparison: (cheap) Backup and Restore < Pilot Light < Warm Standby < Multi-Site/Hot-Site (expensive)

- backup and restore
- pilot light: keeps core functions of the app ready to scale, with minimal setup
- warm standby: keeps full version of the app, but at minimum size
  - when there's a problem, it just increases the size of the app
- multi-site/hot-site: full version of the app, at full size

### security

- services that don't require prior approval for penetration testing
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
- aws prohibits certain attacks
  - DNS zone walking via route 53 hosted zones
  - Denial of Service (DoS)
  - Distributed Denial of Service (DDoS)
  - Simulated DoS
  - Simulated DDoS
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

- workload: set of components that together deliver business value
- aws workloads: applications, services or processes that are running on aws infrastructure
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
- principal: an entity that can make a request to aws (e.g. IAM users, IAM roles, federated users, aws services)
- AWS organizations Service Control Policies (SCPs): centrally manage and restrict permissions across all accounts
  - SCPs do not grand permissions, they **limit** them
  - does not apply to management account
  - applies to all the users and roles of the account, including root
- PII (Personally Identifiable Information): information that can identify a specific individual
  - e.g. full name, email address, phone number, credit card numbers
- aws whitepapers: official documents published by aws that provide best practices, technical guidance and foundational knowledge for using aws cloud services
  - some examples of aws whitepapers: WAF, CAF

