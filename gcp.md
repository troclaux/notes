
# GCP (Google Cloud Provider)

## google cloud SKD

> collection of tools that can interact with Google Cloud APIs/services and manage Google resources

- Google Cloud's SDK (Software Development Kit) is composed by:
  - `gcloud`
  - `gsutil`: CLI tools for managing google cloud storage
  - libraries
  - APIs
  - documentation

- used for:
  - provisioning and managing infrastructure
  - deploying applications
  - storage and data management
  - CI/CD pipelines

### gcloud CLI

> CLI tool to interact with GCP

## IAM (Identity and Access Management)

> security framework that controls access to GCP resources by defining who (principals) can do what (roles) on which resources

- helps manage and secure resource access through fine-grained permissions

- principals: who makes requests to access Google Cloud resources
  - e.g. google accounts, service accounts, google groups, cloud identity
    - service account: non-human account designed to facilitate automated interactions between applications/systems/services
      - represents the identity of applications/services themselves
- roles: set of permissions that are granted to principals

## cloud run

> serverless container hosting service

- purpose: run instances of docker images
- scales automatically
- automatically configures:
  - load balancing
  - DNS
  - HTTPS

- 2 types of applications
  - service: listens and responds web requests
  - job: task that runs to completion

## cloud build API

> fully managed (*) CI/CD service

- automates build, test and deployment of software applications

## artifact registry

> fully managed service that stores and manages artifacts (e.g. docker images)

- artifact: package(s) that contains software
  - e.g. docker images, java artifacts and python packages
- registry: centralized repository that provides a single source of truth for packages

## Google Compute Engine

> fully managed service that provides virtual machines (VMs) for running workloads

- enables running virtual machines on Google's infrastructure
- supports multiple operating systems and machine configurations
- provides high performance and reliability for compute-intensive tasks

## Google Cloud Functions

> serverless compute service for running event-driven code

- executes code in response to events
- automatically scales based on demand
- pay only for compute time used
- supports multiple programming languages

## Google App Engine

> fully managed platform for building and deploying web applications and APIs

- handles infrastructure management automatically
- supports multiple programming languages
- provides automatic scaling and load balancing
- includes built-in services and APIs

## Google Cloud Storage

> fully managed object storage service for storing and accessing data

- provides durable and highly available storage
- supports different storage classes for various use cases
- enables global access and data sharing
- includes versioning and lifecycle management

## Google Persistent Disk

> block storage service for use with Google Cloud VMs

- provides durable and high-performance storage
- supports multiple disk types (SSD, HDD)
- enables snapshots and backups
- allows dynamic resizing

## Google Cloud Filestore

> fully managed file storage service for applications

- provides high-performance file storage
- supports NFS protocol
- enables shared file access
- ideal for enterprise applications

## Google Cloud SQL

> fully managed relational database service

- supports MySQL, PostgreSQL, and SQL Server
- handles maintenance and updates automatically
- provides high availability and backup options
- includes security and encryption features

## Google Cloud Firestore

> fully managed NoSQL document database service

- provides real-time data synchronization
- supports offline data access
- automatically scales with usage
- includes strong consistency guarantees

## Google Cloud Memorystore

> fully managed in-memory data store service

- provides Redis and Memcached compatibility
- enables high-performance caching
- supports automatic scaling
- includes high availability options

## Google Cloud Virtual Network (VPC)

> managed networking service for cloud resources

- enables private network configuration
- supports global connectivity
- provides security and access control
- includes routing and firewall features

## Google Cloud DNS

> managed DNS service for domain and subdomain management

- provides high availability and low latency
- supports public and private zones
- includes automatic scaling
- enables programmatic management

## Google Cloud CDN

> content delivery network for fast content distribution

- accelerates content delivery globally
- reduces latency and bandwidth costs
- provides SSL/TLS security
- includes caching and optimization features

## Google Cloud IAM

> identity and access management service

- controls access to cloud resources
- provides fine-grained permissions
- supports role-based access control
- enables audit logging

## Google Cloud KMS

> key management service for cryptographic keys

- manages encryption keys securely
- supports key rotation and versioning
- enables hardware security module protection
- provides audit logging

## Google Cloud API Gateway

> fully managed service for building and managing APIs

- provides API security and monitoring
- supports multiple API formats
- enables traffic management
- includes authentication and authorization

## Google Cloud Pub/Sub

> fully managed messaging service

- enables real-time message delivery
- supports publish/subscribe patterns
- provides at-least-once delivery
- includes global message routing

## Google Cloud Composer

> fully managed workflow orchestration service

- based on Apache Airflow
- enables complex workflow automation
  - e.g. data pipelines, ETL processes
- provides built-in monitoring
- includes version control integration

## Google Cloud Build

> fully managed CI/CD platform

- automates build and deployment processes
- supports multiple programming languages
- integrates with source repositories
- provides build automation tools

## Google Cloud Source Repositories

> fully managed Git repository service

- provides private Git repositories
- integrates with Cloud Build
- supports source code management
- includes code search features

## Google Cloud Deployment Manager

> infrastructure deployment service

- enables infrastructure as code
- supports template-based deployment
- provides version control
- includes dependency management

## Google Kubernetes Engine (GKE)

> managed Kubernetes service

- automates container orchestration
- provides cluster management
- enables automatic scaling
- includes security features

## Google Cloud Run

> serverless container platform

- runs stateless containers
- automatically scales to zero
- supports HTTP workloads
- includes built-in security features

## Google BigQuery

> fully managed data warehouse service

- enables large-scale data analytics
- provides serverless architecture
- supports SQL queries
- includes machine learning features

## Google Cloud Dataproc

> managed service for running Apache Spark and Hadoop

- enables big data processing
- supports multiple cluster types
- provides automatic scaling
- includes built-in monitoring

## Google Cloud Dataflow

> fully managed data processing service

- enables stream and batch processing
- supports multiple programming languages
- provides automatic scaling
- includes built-in templates

## Google Pub/Sub

> fully managed messaging service

- enables asynchronous messaging
- supports global message delivery
- provides at-least-once delivery
- includes message filtering features

---

- authenticated account vs project, whats the difference between both?
- use cloud service account to create a service account with necessary permissions

- authenticated account: user or service account that is logged in and has access to GCP services
  - types:
    - google user account: used by people (e.g., developers, administrators)
    - service account: used for automated, non-human access to GCP resources (e.g., CI/CD pipeline or an application running in Compute Engine)
- google cloud project: organizatinal unit in GCP used to manage resources and services
  - every resource or action in GCP belongs to a specific project
  - helps logically organize and isolate GCP resources (e.g., permissions, billing, APIs)
- GCP Console: web-based user interface that allows GCP resources management
- fully managed service: cloud provider manages the service's infrastructure and resources
- client library: enables developers to integrate with a service (e.g. Google Cloud) without handling raw HTTP requests, authentication or other low-level operations directly
  - simplifies API integration
  - handles authentication
