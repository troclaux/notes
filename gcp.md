
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

## cloud build API

> fully managed* CI/CD service

- automates build, test and deployment of software applications

## artifact registry

> fully managed service that stores and manages artifacts (e.g. docker images)

- artifact: package(s) that contains software
  - e.g. docker images, java artifacts and python packages
- registry: centralized repository that provides a single source of truth for packages

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
