
# BI (Business Intelligence)

> technology, processes and practices used to collect, integrate, analyze and present business information

> transforms raw data into useful insights

- main objective: support better data-driven decision-making in organizations
  - make informed, strategic decisions
  - identify business trends and patterns
  - improve operational efficiency
  - monitor Key Performance Indicators (KPIs)
  - gain competitive advantage

- components:
  - data visualization
    - e.g. power bi (microsoft), tableau, google data studio
  - data collection
  - data storage
    - e.g. amazon redshift, google bigquery, snowflake
  - ETL tools
    - e.g. apache NiFi
  - OLAP tools
    - e.g. ibm cognos, microsoft analysis services (SSAS)
  - data analysis
  - decision support

- bi cycle
  - data collection
  - data processing and transformation
  - data storage
  - data analysis
  - decision-making

- BI vs data science vs data analytics
  - BI: descriptive (what happened?)
  - data analytics: diagnostic (why did it happen?)
  - data science: predictive/prescriptive (what will/could happen?)

raw data => tools that extract and transform data => storage tools <= analysis tools

## OLTP (OnLine Transaction Processing)

> OLTP systems are built to handle real-time transactional data

- handles day-to-day transactions
- used for transactional processing of data in real-time
- used for processing business transactions such as order processing, inventory management, etc
- normally uses postgresql, mysql, sql server

## OLAP (OnLine Analytical Processing)

> OLAP systems are built for analying large volumes of data

- handles analysis on large volumes of data and supports decision-making
- used for platforms with intense read operations
- used for analyzing large quantities of data from multiple perspectives or dimension
- normally uses amazon redshift, google bigquery, snowflake

### OLAP operations

- data cube: representation of multidimensional data
  - each dimension is an axis
  - each cell in the cube contains measured data (aka facts)
  - sub-cube: subset of the full data cube

- slice: selects single value for one dimension and creates a sub-cube by slicing along that dimension
- dice: selects sub-cube by choosing specific values for two or more dimensions
- drill down: increase detail + reduce granularity
  - e.g. yearly sales => quarterly => monthly => daily
- roll up: reduce detail + increase granularity
  - e.g. daily sales => monthly => quarterly => yearly
- pivot (rotate): reorients the multidimensional view to provide an alternative presentation of data
  - e.g. switch rows and columns

### OLAP architectures

- MOLAP
- ROLAP
- HOLAP

---

- KPI (Key Performance Indicator): measures how effectively an organization is achieving key business objectives
- data mining: process of automatically discovering useful patterns, trends, correlations or rules in large dataset

