
# data warehouse

> centralized system that stores large volumes of structured, historical data from multiple sources

- optimized for querying and analysis
- not for daily operations
- non volatile: data is rarely updated or deleted
- data warehouse is not OLAP itself, but enables OLAP
  - a data warehouse stores the data in a way that is optimized for OLAP-type queries
    - e.g. multidimensional analysis, slicing and dicing, etc
  - OLAP systems query data from the data warehouse for analysis

- use cases:
  - [business intelligence](./business_intelligence.md)
  - storing structured historical data
  - dashboards
  - KPI (*) tracking
  - optimized for complex queries
  - integration of heterogeneous data

- data is typically loaded using ETL (Extract, Transform, Load) processes

- data granularity: level of detail of the data stored
  - high granularity = detailed data
    - e.g. every individual transaction, every second, every customer interaction
  - low granularity = summarized or aggregated data
    - e.g. total sales per month, average revenue per region

| Aspect | **ETL** (Extract, Transform, Load) | **ELT** (Extract, Load, Transform) |
|------|----------------------------------|---------------------------------|
| **Order** | Extract data → Transform it → Load into target | Extract data → Load into target → Transform inside target |
| **Where Transformation Happens** | In an intermediate system (like a data processing server or ETL tool) | Inside the target system (usually a powerful data warehouse) |
| **Typical Target** | Traditional databases, data marts | Modern cloud data warehouses (BigQuery, Snowflake, Redshift) |
| **Data Volume** | Better for smaller or medium-sized datasets | Better for massive datasets (petabytes) |
| **Performance** | Limited by the ETL server capacity | Takes advantage of the scalability of cloud warehouses |
| **Flexibility** | High control over transformation before loading | More flexibility post-load because raw data is preserved |
| **Examples of Use** | Banking, healthcare (where data must be cleaned before storing) | Big data analytics, machine learning, real-time analysis |

- ETL: You clean and organize data before loading it into the database
- ELT: You dump raw data first into the database, and transform it there, using its power

## ETL (Extract, Transform, Load)

data source => (extract) => staging area (transform) => (load) => repository

- Extract: collect data from multiple sources
- Transform: clean, validate and format data
- Load: store transformed data into the warehouse

- transformation happens outside the warehouse
- generally slower than ETL

- staging area: temporary storage where data is placed after extraction and before loading
  - is isolated and ephemeral
- batch: group of data processed together at the same time
  - carries transformed data from staging area and loads them in the repository

### transformation techniques

- discretization: process of transforming continuous variables into discrete variables
- standardization (also called z-score normalization): converts variables in a normal standard distribution
  - transform values of a variable so that they have a mean of zero and a standard deviation of one
- removal of outliers
- dimensionality reduction: decrease the number of variables or dimensions in a dataset

### architectures

| Feature                 | Lambda Architecture                         | Kappa Architecture                      |
|-------------------------|---------------------------------------------|-----------------------------------------|
| **Main Goal**           | Handle batch + real-time data together      | Simplify real-time-only data processing |
| **Components**          | Batch Layer + Speed Layer + Serving Layer   | Stream Processing only                  |
| **Batch Processing?**   | Yes                                         | No (everything is treated as streaming) |
| **Complexity**          | High (two systems to maintain)              | Lower (one unified system)              |
| **Typical use cases**   | Historical data + real-time analytics       | Real-time analytics (fresh data focus)  |

lambda architecture:

```
Input → Batch Layer → Serving Layer → Query
       ↘ Speed Layer ↗
```

kappa architecture:

```
Input → Stream Processor → Serving Layer → Query
```


## ELT (Extract, Load, Transform)

- Extract: pull data from multiple sources
- Load: store data into the warehouse
- Transform: clean, validate and format data inside the warehouse

- transformation happens inside the warehouse (using its resources)
- generally faster than ETL

## fact table

> stores quantitative data

- contains numeric values
- has foreign keys that reference dimension tables
- represents events or transactions
- often very large, with millions of rows

- e.g. Sales, BudgetAllocation

## dimension table

> stores descriptive attributes related to business dimensions

- contains textual data or categorical values like names, descriptions, types
- provides context to the numeric data in the fact table
- generally smaller in size compared to fact tables

- e.g. Customer, Product, Time, Employee, Date

## schema types

| Schema Type            | Fact Tables | Dimensions | Normalization | Complexity | Use Case                    |
|------------------------|-------------|------------|----------------|------------|-----------------------------|
| Star Schema            | 1           | Denormalized | No             | Low        | Fast querying, simple DW    |
| Snowflake Schema       | 1           | Normalized   | Yes            | Medium     | Hierarchies, space savings  |
| Fact Constellation     | Many        | Shared       | Mixed          | High       | Complex, enterprise DW      |

### star schema

- central fact table
- connected to multiple dimension tables
- dimension tables are denormalized

### snowflake schema

- extension of star schema
- central fact table
- dimension tables are normalized into multiple related tables
- more joins, but reduces data redundancy

### fact constellation schema (galaxy schema)

- multiple fact tables share dimension tables
- supports complex applications like combining sales and shipments

## types of data repositories

data warehouse vs data lake vs data mesh:

| Feature            | Data Warehouse                | Data Lake                     | Data Mesh                      |
|--------------------|-------------------------------|-------------------------------|--------------------------------|
| Architecture       | Centralized                   | Centralized                   | Decentralized (Domain-driven)  |
| Data Types         | Structured                    | All types                     | All types                      |
| Schema             | Schema-on-write               | Schema-on-read                | Depends (team decision)        |
| Main Use           | Reporting, BI                 | Raw storage, ML, Big Data     | Scalable, domain-specific data |
| Example            | Snowflake, BigQuery           | Amazon S3, Azure Data Lake    | Netflix architecture           |

analogy:

- data warehouse = a clean, curated museum (only exhibits are allowed after strict vetting)
- data lake = a massive, chaotic storage room (everything is dumped, can find treasures, but also junk)
- data mesh = many museums managed independently by different curators following some shared rules

---

- data warehouse: centralized system that stores large volumes of structured, historical data from multiple sources
- data lake: centralized storage system that holds large amounts of raw data in its original format

- KPI (Key Performance Indicator): measures how effectively an organization is achieving key business objectives

