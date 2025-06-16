
# AWS

> cloud computing platform that provides a wide range of services and tools for building and managing applications and infrastructure in the cloud

- cloud computing: the delivery of computing services over the internet
- service: a software that provides functionality and performs a task or set of tasks for your system
  - examples:
    - apache
    - nginx
    - postgresql
    - auth0
    - prometheus
- infrastructure: the physical or virtual resources that support the operation of a system
  - e.g. servers, storage, network, databases, etc

- types of pricing in AWS
  - pay for compute time
  - pay for data stored in the cloud
  - pay only for data transfer out of the cloud
    - data transfer in is free

- options to manage aws
  - aws management console
  - aws CLI
  - aws SDK: allows the access and management of aws services programmatically

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

- region: separate geographic areas (e.g. `us-east-1` or `sa-east-1`)
  - each region is completely independent from others
  - each region has 3 to 6 Availability Zones, with few exceptions (AZ ⊆ region)
    - Availability Zone (AZ): isolated locations within each region

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

## AWS CLI

> manage aws resources with CLI commands

- list all IAM users in your AWS account: `aws iam list-users`
- get secret from aws secrets manager: `aws secretsmanager get-secret-value --secret-id peso-env-secret --region sa-east-1`

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

- Options to access AWS:
  - AWS Management Console
  - AWS Command Line Interface (CLI)
  - AWS Software Developer Kit (SDK)
    - Used whenever you want to call APIs

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

## VPC (Virtual Private Cloud)

> virtual network that allows you to launch AWS resources in a logically isolated section of the cloud

- allows management over IP addresses, subnets, routing and security
- must have a CIDR block
- allows the creation of public and private subnets
  - subnet: smaller network inside larger network
    - helps organize and manage traffic in a network by dividing it into chunks
    - subnets allow better allocation of IPs from a larger IP range
    - each subnet in aws is associated with one route table
    - use CIDR notation (e.g. `192.168.1.0/24`)

- **public vs private subnet**
  - public subnets have route to internet gateway
    - this means:
      - instances can send traffic to the internet
      - instances can receive traffic from the internet, if security rules allows it
  - private subnets do not have route to internet gateway
    - this means:
      - instances CANNOT initiate outbound internet traffic (unless via NAT gateway)
      - instances CANNOT receive ANY inbound traffic directly from the internet

- route table: defines how traffic flows inside VPC
  - contains rules like:
    - to reach the internet (0.0.0.0/0), go through the Internet Gateway
    - to reach the private subnet, stay local
  - each subnet is associated with a route table
  - only one route table per subnet is allowed

| Feature    | Route Tables       | Security Groups            |
| ---------- | ------------------ | -------------------------- |
| Scope      | Subnet-level       | Instance-level             |
| Purpose    | Direct traffic     | Allow or deny traffic      |
| Layer      | Layer 3 (Network)  | Layer 4 (Transport)        |
| Stateful?  | No                 | Yes                        |
| Controls   | Where traffic goes | Whether traffic is allowed |
| Applied To | Subnets            | EC2 instances, ENIs        |

- internet gateway: allows public subnets to access the internet and receive traffic from it
  - attached to a VPC
  - required for ec2 instances in public subnets to:
    - download packages
    - be accessed via SSH or a browser
- NAT (Network Address Translation) gateway: allows private subnets to access the internet, but prevents the internet from initiating a connection back
  - e.g. private ec2 can download updates or access external APIs without being publicly exposed
  - usually are placed in a public subnet and route private subnet traffic through it

## security groups

> facilitates managing network traffic

- acts as a "firewall"
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

| Feature | AWS Security Groups | Traditional Firewalls (e.g., `ufw`, `iptables`) |
|--------|--------------------------|-----------------------------------------------------|
| Level | AWS-level (cloud) | OS-level (server) |
| Type | Virtual firewall for EC2 and services | Software firewall on the OS |
| Rules | Stateful (return traffic automatically allowed) | Stateful or stateless |
| Scope | Applied per EC2 instance | Applies to whole server |
| Use Case | Allow inbound on port 22 from your IP, or 443 to everyone | Block/allow ports directly on EC2 OS |

### SSH Access

- Create an identity file `.pem` for SSH:
  - EC2 Dashboard > Key Pair > Create Key Pair
  - Copy the public IP of the instance (used in the SSH command)
  - Move the `.pem` file to a secure location and restrict permissions:
    - `chmod 0400 <file>.pem`
  - SSH into the instance:
    - `ssh -i <file>.pem ec2-user@<public_IP>`

## EC2 (Elastic Compute Cloud)

> compute service that allows you to launch virtual servers in the cloud

- instance: a virtual server in the cloud
- AMI: a template for your instance (OS + configs)
- EBS (Elastic Block Store): persistent block storage for instances
- Security Groups: virtual firewalls that control traffic
- key pairs: used for SSH access to Linux or RDP to Windows instances
- elastic IP: static public IP address for your instance

1. choose an AMI (ubuntu, amazon linux, windows, etc)
1. choose an instance type (e.g. t2.micro, m5.large)
1. configure instance details (VPC, subnet, IAM role, etc.)
1. add storage (EBS volumes)
1. add tags (e.g. MyServer)
1. configure security group (firewall rules)
1. launch (with a key pair for SSH)

- inbound traffic: instance <= traffic = outside
- outbound traffic: instance = traffic => outside

### shared responsibility model

- aws
  - infrastructure (global network security)
  - isolation on physical host
  - replacing faulty hardware
  - compliance validation
- me
  - security groups rules
  - OS patches and updates
  - software and utilities installed on the EC2 instance
  - IAM roles assigned to EC2 instances
  - IAM user access management
  - data security on your instance

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
- savings plans
  - commitment to an amount of usage
  - discount based on long-term usage
- capacity reservations
- dedicated instances
- dedicated hosts

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

- tools that can help
  - cloudwatch
  - cost explorer
  - trusted advisor

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
        - bulk: 48 hours (lowest cost)

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

## lambda

> serverless compute service that lets you run code in response to events

- examples of events
  - file uploads
  - http requests
  - database changes

- properties
  - serverless
  - event-driven
  - short-lived functions: each invocation runs in an isolated environment with max duration of 15 minutes
  - stateless: each function runs independently

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

## Amazon FSx

> fully managed service that allows you to launch high-performance file systems in the cloud

- amazon FSx for Windows File Server: for windoes-based applications that require SMB protocol and Active Directory integration
- amazon FSx for Lustre: used for machine learning, analytics, video processing, financial modeling
  - high performance computing
  - scalable file storage
  - lustre = linux + cluster

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

> fully managed service that makes it easy for developers to create, publish, maintain, monitor, and secure APIs at any scale

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

## aurora

> proprietary relational database engine that is part of rds

- fully managed
- highly available
- costs more than rds (20% more), but is more efficient
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

## aws config

> assess, audit and evaluate the configurations of your aws resources

- records configurations and changes over time

- aws config rules: check if your aws resources comply with specific desired configurations or policies

- use cases
  - enforce encryption on RDS databases
  - ensure IAM policies don't allow wildcard actions
  - check whether all resources are in approved regions
  - make sure CloudTrail is enabled and logging correctly

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

## CloudFormation

> fully managed service that allows you to create and manage AWS resources using templates

- template: configuration file that defines infrastructure
- similar to terraform
- uses json or yaml
- easily generate diagrams of your templates
- no need to define ordering and orchestration
- use existing templates on the web
- support for almost all aws resources

### workflow

1. you write a template
1. upload template to CloudFormation
1. CloudFormation provisions and manages the resources

## CloudFront

> content delivery network (CDN) service that delivers data, videos, applications and APIs to customers globally with low latency and high transfer speeds

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
      - any http backend you want

## CloudWatch

> monitoring and observability service that provides data and insights

- monitor your applications
- respond to system-wide performance changes
- optimize resource utilization
- get a unified view of operational health

### CloudWatch Alarm

> monitor metric or math expression and perform one or more actions based on it over a specific time periond

> trigger notifications for any metric

- alarm actions
  - auto scaling: increase/decrease ec2 instances "desired" count
  - ec2 actions: stop, terminate, reboot or recover ec2 instance
  - [sns](#sns-simple-notification-service) notifications: send notification into sns topic
- can choose period on which to evaluate the alarm
- statistic types: %, max, min, sum, sampleCount, etc
- alarm states
  - OK
  - INSUFICIENT_DATA
  - ALARM

### CloudWatch Logs

> logging service that collects, monitors and stores log data from various aws resources, applications and services in real time

- log groups: container for logs from similar sources
- log streams

## CodeArtifact

> TODO

## CodeBuild

> fully managed build service that compiles source code, runs tests, and produces software packages that are ready to deploy

- similar to github actions and gitlab ci
- use cases
  - compile source code
  - run tests
  - produce packages that are ready to be deployed (by CodeDeploy, for example)
- benefits

## CodeCommit

> fully managed source control service that makes it easy for teams to host secure and highly scalable private Git repositories

- aws' github competitor
- source control service
- fully managed

## CodeDeploy

> fully managed deployment service that automates software deployments to a variety of compute services such as EC2, Lambda and on-premises servers

- works with
  - ec2 instance
  - on-premise servers
- servers/instances must be provisioned and configured ahead of time with CodeDeploy agent

## CodePipeline

> continuous integration and continuous delivery (CI/CD) service that automates the build, test and deploy phases of your release process

- fully managed
- similar to github actions and gitlab ci
- compatible with:
  - CodeCommit
  - CodeBuild
  - CodeDeploy
  - elastic Beanstalk
  - CloudFormation
  - github
  - 3rd party services
  - custom plugins

## CodeStar

## Cognito

> fully managed identity service that allows you to add user sign-up, sign-in and access control to your apps quickly and easily

- similar to auth0 and firebase
- identity for web/mobile application's users
- don't create an IAM user for the clients of your application, use Cognito
- instead of giving your app's users aws iam accounts (which are meant for admins and systems), you use cognito to manage their identities securely
- also capable of signing in with google/facebook/twitter accounts

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

## EKR (Elastic Kubernetes Service)

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
- 3 architecture models
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

## Kinesis

> fully managed service that allows you to collect, process, and analyze real-time streaming data

## KMS (Key Management Service)

> managed service that allows you to create and control the encryption keys used to encrypt your data

## Lightsail

> deploy and manage applications/websites with pre-configured cloud resources

## neptune

> fully managed graph database

- high availability
- replication across 3 AZ
- great for knowledge graphs

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

## Redshift

> fully managed data warehouse service that allows you to run complex queries on large datasets

- based on postgresql, but it's not used for OLTP
- used for OLAP (analytics and data warehouse)
- 10x better performance in comparison to other data ware houses
- load data every hour, not every second
- has sql interface for queries
- can be integrated with aws tools, such as quicksight, tableau

## Route 53

> scalable domain name system (DNS) web service that translates domain names into IP addresses

- hosted zone: logical container within aws route 53 that holds DNS records for domains
  - DNS records: maps domain name to a specific resource (e.g. IP address, mail server)
    - types:
      - `A` (Address): points a domain/subdomain directly to a public IPv4 address
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

## SNS (Simple Notification Service)

> fully managed messaging service that allows you to send notifications to a large number of recipients

## SQS (Simple Queue Service)

> message queuing service that allows you to decouple and scale microservices, distributed systems and serverless applications

- fully managed
- low latency
- messages are kept for 14 days
- messages can be
  - xml
  - json

- producer service: produces requests to sqs
- consumer service: consumers requests from sqs
- decoupling: when both producers and consumers scale independently from each other

## Step Functions

> fully managed service that allows you to coordinate multiple AWS services into serverless workflows

## storage gateway

> hybrid cloud storage service that enables on-premises applications to seamlessly use aws cloud storage

- bridge between on-premise data and cloud data in S3
- use cases
  - back up and archive on-premises data to aws
  - provide low-latency access to frequently used data via local caching
  - migrate data to aws over time
  - extend on-premises storage without buying new hardware

## STS (Security Token Service)

> create temporary, short-term credentials to access your aws resources with limited privileges

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

### data transfer

- inbound (data going into aws services): free
- within the same Availability Zone (AZ): free
- from aws to the internet (first 100GB/month): free

## tutorials

- [setup kubernetes in ec2 instance tutorial](https://varunmanik1.medium.com/setting-up-a-kubernetes-cluster-on-aws-ec2-with-ubuntu-22-04-lts-and-kubeadm-5c54930a4659)

## networking costs per GB

> [!TIP]
> use private IP to save money and better network performance.
> use same AZ for maximum savings (at the cost of high availability)

TODO

## billing tools

- pricing calculator: estimates costs in the cloud

- tracking costs in the cloud
  - billing dashboard
  - cost allocation tags
  - cost and usage reports
  - cost explorer
- monitoring agains costs plans
  - billing alarms

TODO

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
- cloudformation to deploy stacks acrosss accounts and region
- trusted advisor to get insights, choose a support plan adapted to your needs
  - support plan: subscription for trusted advisor features
    - basic/free, developer, business, enterprise
- send service logs and access logs to s3 or cloudwatch logs
  - if your account is compromised: change the root password, delete and rotate all passwords/keys, contact the aws support
- cloudtrail to record api calls made within your account
- if your account is compromised: change the root password, delete and rotate all passwords/keys, contact AWS support

---

- serverless: server doesn't requires provisioning and scaling
  - e.g. aws lambda, azure functions, google cloud functions
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
- hybrid service: service that can operate both in the cloud and on-premises
- vpc flow logs: feature that captures information about the ip traffic going to and from network interfaces
  - network interface: virtual network card that connects ec2 instances and other aws resources to a vpc network
  - what the vpc flow logs capture
    - source ip and destination ip
    - source port and destination port
    - protocol (tcp, udp)
    - traffic acceptance (accept or reject)
    - bytes transferred

