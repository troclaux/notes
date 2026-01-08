
# azure

> cloud computing platform that provides a wide range of services and tools for building and managing applications and infrastructure in the cloud

- resource group: logical container in azure
  - holds related resources for an application, workload or project (e.g. VMs, databases, storage accounts, virtual networks, etc)

## shared responsibility model

## Azure Active Directory (AD) or Entra ID

> identity and access management service that helps you manage user identities, authentication and access to applications and services in Azure

- AD DS (Active Directory Domain Services): on-premises directory service
  - domain controller: server that runs AD DS
- azure ad/entra id: the cloud version, equivalent to aws iam
  - doesn't include the organizational unit (OU) class
  - doesn't use kerberos authentication, uses http and https protocols like SAML, ws-federation and openid connect for authentication
  - uses oauth for authorization

- service principal: application/automation identity in a tenant that can be assigned roles to access resources
  - created when you register an app (app registration = global template; service principal = tenant-local instance)
  - authenticate with a client secret or certificate; prefer certificates for longer-lived automation
  - assign Azure RBAC roles at the right scope (resource, resource group, subscription) with least privilege
  - managed identity is a Microsoft-managed service principal bound to a resource; use it instead of manual credentials when available

- also known as a directory
- provides authentication and authorization
- provides SSO (login once and access multiple apps)
- common authentication protocols:
  - SAML 2.0: XML assertions for SSO between an identity provider (Entra ID) and a service provider; popular with enterprise SaaS
  - WS-Federation: older Microsoft token-based protocol for legacy apps; similar use case to SAML but less common today
  - OpenID Connect (OIDC): identity layer on OAuth 2.0; returns ID tokens (JWTs) to prove who the user is for modern web/mobile apps
  - OAuth 2.0: authorization framework that issues access tokens for API access; pair with OIDC when you need authentication

### users, roles and identity basics

- tenant: isolated directory for an org; holds users, groups, app registrations and service principals
- security principals: what you assign roles to
  - users: `member` (internal) or `guest` (B2B from another tenant)
  - groups: bundle users for easier role assignment
  - service principals: app/automation identities created from app registrations
  - managed identities: azure-managed service principals
    - system-assigned ties lifecycle to one resource; deleted when the resource is removed
    - user-assigned is reusable across resources
- azure RBAC controls access to resources: `role assignment = principal + role definition + scope`
  - scopes: management group > subscription > resource group > resource; prefer least privilege
  - role definitions can be built-in (e.g., Owner, Contributor, Reader) or custom
- directory roles (e.g., Global Administrator) control Entra/Azure AD features; RBAC roles control resource access
- PIM (Privileged Identity Management) can make high-privilege roles just-in-time with approvals and time limits

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
## Azure CLI

> CLI tool that allows users to manage and automate Azure resources and services from a terminal or script, supporting a wide range of operations like creating, managing and deleting resources across Azure's cloud platform

```bash
az <resource group> <subcomand> [parameters]
```

- `az login`: log in
- `az ... --help`: see available options
- `az vm create ...`: create virtual machines

## Azure Virtual Network (VNet)

> network service that enables secure, isolated and private communication between Azure resources, including VMs, services and on-premises networks

- 

## Azure Virtual Machines

> compute service that allows you to create and manage virtualized instances of computers running on Azure's infrastructure

### sizes

- general purpose
  - A, B, D-series
- compute optimized: more cpu, less ram
  - F-series
- memory optimized: lots of ram per cpu
  - E, M-series
- storage optimized: fast I/O
  - L-series
- gpu optimized: AI, ML, rendering
  - NC, ND, NV-series
- high performance compute: very powerful CPU, low latency
  - H-series

## Azure Functions

> serverless compute service that lets you run event-driven code without having to manage infrastructure

## Azure App Service

> fully managed platform for building, deploying, and scaling web apps and APIs using various programming languages

## Azure Blob Storage

> scalable object storage service for unstructured data like text, images, videos, and backups

- BLOB: Binary Large OBject
  - blobs: the actual objects/files
    - types of blobs
      - block blob: for text, images, videos, backups (most common)
        - suppots large files (up to ~190 TB)
      - append blob: optimized for logs and data that must be appended
      - page blob: random-access storage, used for virtual disks in azure virtual machines
- containers: life folders inside the storage account

### storage tiers (cost optimization)

- hot tier: optimized for frequently accessed data
- cool tier: cheaper, optimized for infrequently accessed data
- archive tier: very cheap, for rarely accessed data (retrieval may take hours)

## Azure Disk Storage

> block-level storage that provides persistent storage for Azure Virtual Machines

## Azure File Storage

> managed file storage service that allows you to share files across multiple VMs, accessible via the SMB protocol

## Azure SQL Database

> fully managed relational database service based on SQL Server, offering automatic scaling, high availability, and security features

## Azure Cosmos DB

> globally distributed, multi-model NoSQL database service designed to provide low-latency, scalable access to data across multiple regions

## Azure Cache for Redis

> in-memory data store based on Redis, used to improve the performance of applications by caching frequently accessed data

## Azure DNS

> domain name system service that enables you to host and manage your DNS records for domain names

## Azure Content Delivery Network (CDN)

> distributed content delivery network service that accelerates the delivery of websites, media, and APIs by caching content in multiple locations worldwide

## Azure Key Vault

> secure service for managing secrets, keys, and certificates, allowing you to control and protect access to sensitive information

## Azure API Management

> service that enables developers to create, publish, and manage APIs with built-in security, monitoring, and analytics

## Azure Service Bus

> messaging service for building scalable and reliable applications by enabling communication between different systems or services through queues and topics

## Azure Logic Apps

> service that allows you to automate workflows, integrate apps, and process data across different systems using a visual designer

## Azure DevOps Services

> suite of development tools for managing the entire software development lifecycle, including planning, development, testing, and deployment

## Azure Pipelines

> continuous integration and delivery (CI/CD) service that automates the building, testing, and deployment of applications

- basics:
  - pipelines are defined in YAML and can be triggered by branches, PRs, schedules or manual runs
  - `stages` group related jobs, `jobs` run on an agent, and `steps` are individual tasks or scripts
  - agents come from hosted pools (e.g., `ubuntu-latest`) or self-hosted pools; each run gets a fresh workspace
  - variables let you centralize values (paths, tags, service connections) and can be secret-protected in variable groups or key vaults
  - service connections (e.g., to ACR) handle authentication without hardcoding credentials

- example pipeline (multi-image Docker build/push to ACR):
  - trigger on `main`; uses only the self repo as a resource
  - variables define the ACR service connection, registry name, image repositories, Dockerfile paths, build contexts and agent image
  - stage `Build` runs one job on `ubuntu-latest`
  - `PipAuthenticate@1` injects an authenticated `PIP_EXTRA_INDEX_URL` for the `prod_quantique` feed without hardcoding credentials in the YAML
  - inline `Bash@3` step writes that feed URL into `$(Agent.TempDirectory)/azure_token.txt`, scoped to the agent’s temp space
  - multiple `Docker@2` tasks:
    - build images for `commercial_app_dash`, `commercial_app_daily_ihfa`, `commercial_app_quarterly_ihfa`, `commercial_app_daily_cvm_returns`, and `commercial_app_daily_cdi_fund_data`
    - each build uses `--secret id=azuretoken,src=$(Agent.TempDirectory)/azure_token.txt` so the pip token is provided to BuildKit only during `RUN --mount=type=secret` steps and never baked into layers or logs
    - each build tags `latest` and targets the shared `quantm3airflowimages.azurecr.io` registry via the `azure-container-registry-conn` service connection
    - corresponding `Docker@2` push tasks publish the `latest` tag for each repository to ACR

## Azure Repos

> source code management service that allows teams to collaborate on code using Git or Team Foundation Version Control (TFVC)

## Azure Kubernetes Service (AKS)

> managed Kubernetes service that simplifies the deployment, management, and scaling of containerized applications

## Azure Container Instances

> serverless container service that allows you to run containers without managing virtual machines

## Azure Container Registry

> private registry for storing and managing Docker container images and Helm charts for use in Azure Kubernetes Service or other containerized environments

- `<registry-name>.azurecr.io/<repository>:<tag>`
- `myregistry.azurecr.io/webapp:1.0`

- registry: the root namespace for all your images
- repositories: logical collection of related images

## Azure Synapse Analytics

> integrated analytics service that combines big data and data warehousing to analyze large datasets and gain insights

## Azure Data Lake Storage

> highly scalable data storage service designed for big data analytics workloads, optimized for storing massive amounts of unstructured data

## Azure HDInsight

> fully managed cloud service that makes it easy to process and analyze large datasets using popular open-source frameworks like Hadoop, Spark, and Hive

## Azure Stream Analytics

> real-time data stream processing service that allows you to analyze and process data streams from various sources, such as IoT devices and social media feeds

## Azure Databricks

> Apache Spark-based analytics platform that simplifies big data processing and machine learning tasks by providing a collaborative environment for data engineers, data scientists, and analysts

## Cost Management

- inbound data (into azure) is always free
- outbound data (out of azure) is charged when data leaves to the public internet

- within azure
  - same region transfers: free
  - across regions: charged (e.g. vm in east us => storage in west us)
  - between services: some inter-service transfers are free if in same region, but cross-region always costs

### Azure Cost Management

> tools for analyzing usage, forecasting future costs and optimizing spending

- cost analysis: similar to aws' cost explorer
- budgets
- forecasting
- exports: automatically push usage and cost data to azure storage or power BI for custom reporting

### Azure Billing

> see invoices, payment methods, subscriptions, charges
