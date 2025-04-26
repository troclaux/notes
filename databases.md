[sql](./sql.md)
[postgresql](./postgresql.md)
[MySQL](./mysql.md)
[MongoDB](./mongodb.md)

# databases

> collection of data that is organized so that it can be easily accessed, managed and updated

- data: raw facts and figures without context
  - example: 98, 102, 95
  - no meaning on its own
- information: data that has been processed and given context to be meaningful
  - example: the temperatures this week were 98°F, 102°F, and 97°F
  - information = data + context
- knowledge: insights and understanding gained from analyzing information
  - example: those temperatures indicate an extreme heatwave; precautions should be taken
  - knowledge = information + meaning
- wisdom: application of knowledge
  - example: due to the drop in temperature, we should alert farmers to prepare for potential crop damage
  - wisdom = knowledge + application

- types of databases:
  - SQL databases:
    - relational
    - analytical
    - data is stored in tables (rows and columns)
    - enforces schema and [ACID](./databases.md#acid) properties
    - examples: PostgreSQL, MySQL, oracle, microsoft sql server
  - NoSQL databases:
    - document-oriented:
      - store data as documents (often json or bson)
      - flexible schema
      - examples: MongoDB, CouchDB
    - key-value stores:
      - data is stored as key-value pairs
      - very fast ookups
      - examples: redis, dynamodb
    - column family stores:
      - stores data in columns instead of rows
      - optimized for reading/writing large volumes of data
      - examples: apache cassandra, hbase
    - graph databases:
      - store data as nodes and edges (like a graph)
      - great for representing relationships
      - examples: neo4j, arangodb

- record = row = tuple
  - single data entry in a table
- field = column = attribute
  - defines specific type of data
- table: holds data about an entity (e.g. users, products, orders)

- degree of a table: nº of columns in table
- cardinality: nº of rows in table

- failover: automatic switching from a primary database to a replica in case of failure
  - guarantees redundancy
- failback: return to original database once it's fixed again

## CRUD

> four basic operations used in persistent storage

- Create: add new data
- Read: retrieve existing data
- Update: modify existing data
- Delete: remove data

## ACID

> set of properties of database transactions intended to guarantee data validity despite errors

- transactions: a sequence of operations that are treated as a single unit of work
  - e.g. transferring money from one account to another

- Atomicity: transactions are all or nothing, if any part of it fails, the entire transaction is rolled back
- Consistency: database always moves from one valid state to another
- Isolation: transactions are independent of each other
- Durability: once a transaction is committed, it is permanently saved

## ER Diagram

> visual representation of data

[ER model cardinality explained with examples](https://www.gleek.io/blog/er-model-cardinality)
[ER diagrams](https://www.lucidchart.com/pages/er-diagrams)

- strong entity (rectangle/square): can be uniquely identified by its own attributes without relying on other entity
- weak entity: depends on a strong entity to be identified
  - rectangle that contains another rectangle or balloon with dotted underline
  - its primary key is in another entity
    - cannot be uniquely identified by its own attributes
- associative entity: M:N relationships that have attributes
- relationship (diamond): describes how entities are associated with each other
  - id/identifying relationship: diamond that contains another diamond
    - relationship that defines the primary key of a weak entity
  - recursive relationship: occurs when an entity has multiple functions with itself
- attribute (empty dot): property or characteristic of an entity
  - identifying attribute: black dot or balloon with underlined attribute
  - multivalued attribute: can have multiple values (e.g. phone number)
    - balloon that contains another balloon
  - derived attribute: attribute that can be defined from another attributes
    - dotted baloon
    - e.g. age is derived attribute from birth_date
  - composite attribute: can be fragmented into sub-attributes (e.g. address => block, street, etc)
- inheritance (triangle)
  - total or partial
    - total (t): all instances of superclass must be instances of at least one subclass
    - partial (p): some instances of superclass may not belong to any subclasses
  - overlapped or disjoint
    - disjoint (x): entities in superclass can only be instance of one subclass
    - overlapping (c): entities in superclass can instances of multiples subclasses simultaneously

### crow's foot notation

- relationships are bidirected
- there are symbols on both sides that represent min:max cardinality

- `[ student ] >|--has--|| [ school ]`
  - student relates with one and exactly one school
  - school relates to one or many students

- `--------0|` = 0:1
- `--------0<` = 0:n
- `--------||` = 1:1 (one and exactly one)
- `--------|<` = 1:n

### chen notation

- `[ student ] 1----<has>---1 [ seat ]`
  - one student can have maximum of 1 seat
  - one seat belongs to maximum of 1 student

### barker notation

- `#`: primary key
- `*` or `·`: obligatory attribute
- `o`: optional attribute

```
+----------------+          +----------------+          +----------------+
|   Customer     |          |     Order      |          |    Product     |
|----------------|          |----------------|          |----------------|
|*CustomerID     |     1,M  |*OrderID        |          |*ProductID      |
| Name           |----------| OrderDate      |     1,M  | ProductName    |
| Email          |          |#CustomerID     |----------| Price          |
+----------------+          +----------------+          +----------------+
                              0,M  |                         0,M   |
                                   |                               |
                                   |                               |
                                   +-------------------------------+
                                                1,M
                                           +------------+
                                           |  OrderItem |
                                           |------------|
                                           |*OrderItemID|
                                           |#OrderID    |
                                           |#ProductID  |
                                           | Quantity   |
                                           +------------+
```

## Class diagram

- cardinality symbols
  - how to read class diagrams:
    - Entity1 ----------------------> 0..* Entity2
    - Entity1 can have 0..* relations with Entity2
      - the multiplicity is closer to the entity it applies to
  - multiplicity
    - 1: one to one
    - *: many
    - 0..1: zero or one
    - 0..*: zero or many
    - 1..*: one or many
    - n..m: n to m

## theory

- DBMS (Database Management System): software that allows users and programs to interact with the database
  - handles data storage, retrieval, updates and security
  - e.g. MySQL, PostgreSQL, sqlite, MongoDB
- DB (DataBase): organized collection of data that can be easily accessed, managed and updated
- DBS (Database System) = DBMS + DB

- SQL sublanguages
  - DDL (Data Definition Language): object definition
    - create
    - alter
    - drop
    - truncate
  - DML (Data Manipulation Language): data manipulation
    - delete
    - update
    - select
    - DQL (Data Query Language)
      - select
  - DTL (Data Transaction Language): transactions control
    - commit
    - rollback
    - savepoint
  - DCL (Data Control Language): safety and access control
    - grant
    - revoke

- abstraction: level of details/complexity with which data is represented
  - fewer detail == higher abstraction
  - more detail == lower abstraction

### conceptual model (high-level abstraction)

- captures what the data is, not how it will be implemented
- made for stakeholders, business analysts and data architects

- describes entities (e.g. Customer, Order), attributes (e.g. name, address) and relationships (e.g. Customer places Order)
- no concern for data types
- Entity-Relationship Diagrams (ERDs)

### logical model (middle-level abstraction)

- maps the conceptual model to a logical structure that a database system can understand
- made for database designers and developers

- includes tables, columns, primary keys, foreign keys, normalization
- DBMS-agnostic (not specific to MySQL, PostgreSQL, etc)
- includes constraints, data types (abstract) and relationships

example of logical data model:

| Table       | Columns                                            |
|-------------|--------------------------------------------------|
| Customers   | CustomerID (INT, PK), Name (VARCHAR), Email (VARCHAR) |
| Orders      | OrderID (INT, PK), CustomerID (INT, FK), OrderDate (DATE) |
| Products    | ProductID (INT, PK), ProductName (VARCHAR), Price (DECIMAL) |
| OrderItems  | ItemID (INT, PK), OrderID (INT, FK), ProductID (INT, FK), Quantity (INT) |

### physical model (low-level abstraction)

- specifies how data is actually stored and accessed in the DBMS
- made for database administrators and system architects

- includes storage formats, indexes, partitioning, access paths
- DBMS-specific (PostgreSQL, Oracle, MySQL, etc.)
- may include things like tablespaces, file paths and buffer sizes

example of Physical Data Model for a Relational DBMS:

```sql
CREATE TABLE Customers (
    CustomerID INT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(100) UNIQUE
);

CREATE TABLE Orders (
    OrderID INT PRIMARY KEY,
    CustomerID INT,
    OrderDate DATE,
    FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID)
);

CREATE INDEX idx_email ON Customers(Email);
```

## database schema

[what is a database schema](https://www.ibm.com/topics/database-schema)

> defines how data is organized within a relational database

- types of database schema
  - conceptual: high-level view using ER diagrams, without technical details
  - logical: database-neutral tables, columns, relationships with data types and constraints
  - physical: specific DBMS implementation with indexes, storage, and access paths

## migrations

> process of transforming a database from one state to another

- good migrations are small, incremental and reversible
- can be automated in continuous deployment environments

- up migration: change database to new state
- down migration: revert database to previous state

## metadata

> data about data

- descriptive metadata: used to identify and discover data
  - what it includes: title, author, date created, keywords, description
  - "what is this data about?"
- structural metadata: describes how data is organized or related internally
  - what it includes: tables of contents, page numbers, chapters, data models, xml schema, file formats
  - "how is this data structured or arranged?"
- administrative metadata: helps with management and processing of data
  - what it includes: file type, creation date, access permissions, file size, owner, storage info
  - "how can we store, manage and protect this data?"

## relational algebra

- Selection (σ): `σ <conditions> (<table>)`
  - example: σ age > 30 (employees)
  - `SELECT * FROM employees WHERE age > 30;`
  - OBS: σ => Sigma => Select * => sElEct => whErE
- Projection (π): `π <columns> (<conditions>) (<table>)`
  - example: π name, age (employees)
  - `SELECT name, age FROM employees;`
  - OBS: π => PI => PIck => Projection
- Union (∪)
  - example: employees ∪ contractors
  - `SELECT email FROM employees UNION SELECT email FROM contractors;`
- Intersection (∩)
  - example: employees ∩ contractors
  - `SELECT email FROM employees INTERSECT SELECT email FROM contractors;`
- Difference (−)
  - example: employees − contractors
  - `SELECT email FROM employees EXCEPT SELECT email FROM contractors;`
- Rename (ρ)
  - example: ρ(name/employee_name)(employees)
  - `SELECT name AS employee_name FROM employees;`
  - OBS: ρ => Rho => Rename
- Cartesian Product (×)
  - example: employees × departments
- Join (⨝  or ٭)
  - example: employees ⨝  employees.department_id = departments.department_id departments

## normal forms and data normalization

> rules used in relational database design to eliminate redundancy and ensure data integrity

- split large, complex tables into smaller, simpler ones
- while also defining relationships between tables to improve data consistency, storage efficiency and query performance

- normalization: the process of applying the normal forms rules
- Unnormalized Form (UNF or 0NF): table that is not even in 1NF

```
5NF ⊂ 4NF ⊂ BCNF ⊂ 3NF ⊂ 2NF ⊂ 1NF ⊂ 0NF
```

> [!IMPORTANT]
> primary key in SQL != primary key in data normalization/normal forms
> in data normalization, primary key is the collections that uniquely identifies a row
> in SQL, primary key is a single column that uniquely identifies a row

#### First Normal Form (1NF)

- cells must be atomic (no repeating groups or lists/sets in a column)
- all entries in a column are of the same type

> [!TIP]
> 1NF = 1 value per cell

#### Second Normal Form (2NF)

- must be in 1NF
- no partial dependency: every non-key column depends on the whole primary key
  - non-key column: columns that are not part of the primary key

> [!IMPORTANT]
> tables can have composite primary keys (multiple columns that collectively functions as a single primary key)

#### Third Normal Form (3NF)

- must be in 2NF
- no transitive dependencies (non-key columns shouldn't depend on other non-key columns)

#### Boyce-Codd Normal Form (BCNF or 3.5NF)

- must be in 3NF
- even candidate keys shouldn't depend on other non-key columns
- candidate key: any column (or combination of columns) that can uniquely identify each row in a table, without redundancy
  - must be unique and minimal (no extra columns, removing any part breaks uniqueness)

#### Fourth Normal Form (4NF)

- must be in BCNF
- avoid storing multiple independent sets of data in one table
- remove multivalued dependencies

#### Fifth Normal Form (5NF)

- must be in 4NF
- eliminate redundancy from complex joins

### rules of thumb for database design

1. every table should always have a unique identifier (primary key)
2. 90% of the time, that unique identifier will be a single column named id
3. avoid duplicate data
4. avoid storing data that is completely dependent on other data. instead, compute it on the fly when you need it
5. keep your schema as simple as you can. optimize for a normalized database first. only denormalize for speed's sake when you start to run into performance problems

---

- data consistency: data in the database is accurate, correct and reflects the real-world entities it represents (without contradictions)
- database connection pool: mechanism that allows multiple applications/services to share a set of database connections
  - why do we create connection pools? to avoid the overhead of opening and closing a connection for each request
  - maintains a pool of open connections, instead of opening and closing a connection for each request
    - creating a new database connection is expensive
      - authentication, initialization, etc
- pagination: process of dividing a large dataset into smaller pages and providing navigation controls to allow users to browse data in smaller, manageable sections

- dump a database = export entire structure and/or data of a database into a file, usually in sql format
