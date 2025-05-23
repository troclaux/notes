[sql](./sql.md)
[databases](./databases.md)

# MySQL

> open-source relational database management system that uses SQL to manage and manipulate data

## CLI

- access monitor: `mysql -u root -p`
  - `-u`: specify username (e.g. `root`)
  - `-p`: prompt for password

after logging in:

- list all databases: `SHOW DATABASES;`
- select a database: `USE my_database;`
- create a new database: `CREATE DATABASE my_database;`
- delete a database: `DROP DATABASE my_database;`

- show a table: `SHOW TABLES;`
- view table structure: `DESCRIBE users;` or `SHOW COLUMNS FROM users;`

- exit mysql CLI: `exit;`

## architecture

- mysql instance: running mysql server
  - a single computer can run multiple instances of mysql, each one independent (with different settings and ports)
- schema: inside an instance, a schema is a logical containers that holds tables, views, procedures and triggers
  - a sinfle mysql instance can have many schemas
- storage engine: internal part of mysql that decides how data is actually stored on disk and accessed
  - InnoDB is the default storage engine

### InnoDB

> default and most powerful engine in MySQL

- ACID compliance

#### key components

- buffer pool: memory area to cache data and indexes (makes access faster)
- redo logs: track changes (INSERT, UPDATE, DELETE), used to restore data after a crash
- undo logs: used for rollback (undoing changes), also used for MVCC (seeing old versions of data)
- clustered index

## data types

### numeric

> [!IMPORTANT]
> use `DECIMAL` for money instead of `FLOAT` to avoid rounding errors

- `INTEGER`/`INT`: standard integer number of 4 bytes
- `MEDIUMINT`: integer number of 3 bytes
- `SMALLINT`: integer number of 2 bytes
- `TINYINT`: integer number of 1 bytes
- `BIGINT`: integer number of 8 bytes
- `DECIMAL`/`NUMERIC`: exact fixed-point number with specified precision and scale
- `FLOAT`: single-precision floating-point number (4 bytes)
- `DOUBLE`: double-precision floating-point number (8 bytes)
- `BIT`: stores bit values, can store 1 to 64 bits

### string

- `CHAR(n)`: fixed-length string of n chars (1-255 characters)
  - if n is higher that string that will be stored, the unused part is padded with spaces
    - e.g. storing "hi" in CHAR(5) attribute => "hi   "
- `VARCHAR(n)`: variable-length string (1-65,535 characters)
  - uses only the actual length of the string
- `BINARY`: fixed-length binary string
- `VARBINARY`: variable-length binary string
- `TEXT`: string with max length of 65,535 characters
- `BLOB`: stores binary objects
- `ENUM`: string object with only one value from a predefined list
- `SET`: string object that can have zero or more values from a predefined list
- `JSON`: native JSON data type for storing JSON documents

### date

> [!WARNING]
> always surround date types with single quotes
> avoid double quotes

- `DATE`: date in format 'YYYY-MM-DD', range 1000-01-01 to 9999-12-31
- `TIME`: time in format 'HH:MM:SS', range -838:59:59 to 838:59:59
- `DATETIME`: date and time, format 'YYYY-MM-DD HH:MM:SS'
- `TIMESTAMP`: timestamp, stored as seconds since Unix epoch (00:00:00 UTC on January 1, 1970)
- `YEAR`: year in 4-digit format '2025', range 1901 to 2155
