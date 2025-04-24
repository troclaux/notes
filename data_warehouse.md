
# data warehouse

> centralized system that stores large volumes of structured, historical data from multiple sources

- optimized for querying and analysis
- data warehouse is not OLAP itself, but enables OLAP
  - OLAP systems query data from the data warehouse for analysis
- not for daily operations

- use cases:
  - storing structured historical data
  - optimized for complex queries
  - integration of heterogeneous data

- data is typically loaded using ETL (Extract, Transform, Load) processes

- data granularity: level of detail of the data stored
  - high granularity = detailed data
    - e.g. every individual transaction, every second, every customer interaction
  - low granularity = summarized or aggregated data
    - e.g. total sales per month, average revenue per region

## ETL (Extract, Transform, Load)

- Extract: pull data from multiple sources
- Transform: clean, validate and format data
- Load: store transformed data into the warehouse

- transformation happens outside the warehouse
- generally slower than ETL

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
