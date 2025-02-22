
# cloud computing

> delivery of computing services over the internet

- service: a software that provides functionality and performs a task or set of tasks for your system
  - examples:
    - apache
    - nginx
    - postgresql
    - auth0
    - prometheus
- infrastructure: the physical or virtual resources that support the operation of a system
  - e.g. servers, storage, network, databases, etc

- serverless: servers aren't directly managed
  - e.g. AWS Lambda
- private cloud: cloud infrastructure that is operated solely for a single organization
- public cloud: cloud infrastructure that is available to the general public

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

- scalability vs elasticity vs agility
  - scalability: ability to accommodate larger load by adding resources
    - scaling up: increasing resources of a single instance
    - scaling out: adding more instances
  - elasticity: ability to automatically scale resources based on demand
  - agility: how quickly resources get provisioned to developers

- vertical scalability: increasing/decreasing the size of an instance
- Horizontal scalability: increasing/decreasing the number of instances/systems for your application

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
  - block: a sequence of bytes or bits having a fixed length
    - each block has a unique id, no additional metadata is stored
- object storage: data is stored as objects
  - e.g. AWS S3
  - object: a collection of data that contains both data and metadata
  - ideal for non-structured data


## Compute

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure Virtual Machines              | AWS EC2 (Elastic Compute Cloud)    | Google Compute Engine              |
| Azure Functions                     | AWS Lambda                         | Google Cloud Functions             |
| Azure App Service                   | AWS Elastic Beanstalk              | Google App Engine                  |

## Storage

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure Blob Storage                  | AWS S3 (Simple Storage Service)    | Google Cloud Storage               |
| Azure Disk Storage                  | AWS EBS (Elastic Block Store)      | Google Persistent Disk             |
| Azure File Storage                  | AWS EFS (Elastic File System)      | Google Cloud Filestore             |

## Database

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure SQL Database                  | AWS RDS (Relational Database Service) | Google Cloud SQL                   |
| Azure Cosmos DB                     | AWS DynamoDB                       | Google Cloud Firestore             |
| Azure Cache for Redis               | AWS ElastiCache                    | Google Cloud Memorystore           |

## Content Delivery and Network

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure Virtual Network Azure DNS     | AWS VPC (Virtual Private Cloud)    | Google Cloud Virtual Network (VPC) |
| Azure Content                       | AWS Route 53                       | Google Cloud DNS                   |
| Delivery Network (CDN)              | AWS CloudFront                     | Google Cloud CDN                   |

## Security and Identity Management

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure Active Directory              | AWS IAM (Identity and Access Management) | Google Cloud Identity and Access Management (IAM) |
| Azure Key Vault                     | AWS Cognito                        | Google Cloud Identity Platform    |
| ---                                 | AWS Key Management Service (KMS)   | Google Cloud Key Management Service (KMS) |

## Applications and Services

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure API Management                | AWS API Gateway                    | Google Cloud API Gateway           |
| Azure Service Bus                   | AWS SQS (Simple Queue Service)     | Google Cloud Pub/Sub               |
| Azure Logic Apps                    | AWS Step Functions                 | Google Cloud Composer              |

## Development, CI/CD and DevOps

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure DevOps Services               | AWS CodeCommit & AWS CodePipeline  | Google Cloud Build                 |
| Azure Pipelines                     | AWS CodeBuild                      | Google Cloud Source Repositories   |
| Azure Repos                         | AWS CodeDeploy                     | Google Cloud Deployment Manager    |

## Container Orchestration

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure Kubernetes Service (AKS)      | AWS ECS (Elastic Container Service) | Google Kubernetes Engine (GKE)     |
| Azure Container Instances (ACI)     | AWS ECR (Elastic Container Registry) | Google Cloud Run                  |
| Azure Container Registry            | AWS EKS (Elastic Kubernetes Service) | Google Cloud Container Registry    |

## Data Analysis and BigData

| **Azure**                          | **AWS**                            | **GCP**                            |
|-------------------------------------|------------------------------------|------------------------------------|
| Azure Synapse Analytics             | AWS Redshift                       | Google BigQuery                    |
| Azure Data Lake Storage             | AWS Athena                         | Google Cloud Dataproc              |
| Azure HDInsight                     | AWS EMR (Elastic MapReduce)        | Google Cloud Dataflow              |
| Azure Stream Analytics              | AWS Kinesis                        | Google Pub/Sub                     |
| Azure Databricks                    |                                    |                                    |

---

- OLAP
  - used for platforms with intense read operations
  - used for analyzing large quantities of data from multiple perspectives or dimension
- OLTP
  - used for transactional processing of data in real-time
  - used for processing business transactions such as order processing, inventory management, etc
