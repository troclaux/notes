[sql](./sql.md)
[postgresql](./postgresql.md)
[MySQL](./mysql.md)
[MongoDB](./mongodb.md)

# databases

> collection of data that is organized so that it can be easily accessed, managed, and updated

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
  - relational databases
    - uses sql
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

- failover: automatic switching from a primary database to a replica in case of failure
  - guarantees redundancy
- fallback: return to original database once it's fixed again

## ER Diagram

[ER model cardinality explained with examples](https://www.gleek.io/blog/er-model-cardinality)

- strong entity: rectangle/square
  - can be uniquely identified by its own attributes without relying on other entity
- weak entity: rectangle that contains another rectangle or balloon with dotted underline
  - cannot be uniquely identified by its own attributes
  - depends on a strong entity to be identified
- relationship: diamond
  - id relationship: diamond that contains another diamond
  - recursive relationship: occurs when an entity has multiple functions with itself
- attribute: empty dot
  - identifying attribute: black dot or balloon with underlined attribute
  - multivalued attributes: balloon that contains another balloon
  - derived attributes: dotted baloon
    - attribute that can be defined from another attributes
- inheritance: triangle
  - total or partial
    - total: all instances of superclass must be instances of subclass
    - partial: some instances of superclass must be instances of subclass
  - overlapped or disjoint
    - disjoint: entities in superclass can only be instance of one subclass
    - overlapped: entities in superclass can instances of multiples subclasses simultaneously

### crow's notation

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

- DBS: Database System
- DBMS: Database Management System
  - e.g. MySQL, PostgreSQL, sqlite
- DB: Database
- DBS = DBMS + DB

- categories of data models:
  - conceptual (high level)
    - Entity-Relationship (ER) diagram
    - typical users: business analysts, designers
    - records what data is stored, but not how
  - logical (representative)
    - logical data model
    - typical users: developers, DBAs
  - physical (low level)
    - physical data model
    - typical users: database administrators

example of logical data model:

| Table       | Columns                                            |
|-------------|--------------------------------------------------|
| Customers   | CustomerID (INT, PK), Name (VARCHAR), Email (VARCHAR) |
| Orders      | OrderID (INT, PK), CustomerID (INT, FK), OrderDate (DATE) |
| Products    | ProductID (INT, PK), ProductName (VARCHAR), Price (DECIMAL) |
| OrderItems  | ItemID (INT, PK), OrderID (INT, FK), ProductID (INT, FK), Quantity (INT) |

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

## ACID

> properties of database transactions

- transactions: a sequence of operations that are treated as a single unit of work
  - e.g. transferring money from one account to another

- Atomicity: all or nothing
- Consistency: data is always in a valid state
- Isolation: transactions are independent of each other
- Durability: changes are permanent

---

- database connection pool: mechanism that allows multiple applications/services to share a set of database connections
  - why do we create connection pools? to avoid the overhead of opening and closing a connection for each request
  - maintains a pool of open connections, instead of opening and closing a connection for each request
    - creating a new database connection is expensive
      - authentication, initialization, etc
- pagination: process of dividing a large dataset into smaller pages and providing navigation controls to allow users to browse data in smaller, manageable sections
