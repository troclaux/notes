[databases](./databases.md)
[postgresql](./postgresql.md)

[ER model cardinality explained with examples](https://www.gleek.io/blog/er-model-cardinality)

# SQL

order of operations in SQL query:

```
FROM => WHERE => GROUP BY => HAVING => SELECT => ORDER BY => LIMIT
```

- `SELECT`: retrieve data from database
  - `*`: select all
- `FROM`: select the target(s) table(s)
  - list the tables from where data is fetched
- `WHERE`: filter records based on conditions
  - `BETWEEN`: selects values within given range (values can be numbers, text or dates)
  - `LIKE`: filter records that matches the string
    - wildcards: used for pattern matching, like regex
      - `%`: zero or more chars
        - `SELECT * FROM products WHERE product_name LIKE 'banana%';`
      - `_`: single char
        - `SELECT * FROM products WHERE product_name LIKE '__oot';`
      - `[]`: any single char within the brackets
      - `^`: any char NOT in the brackets
        - `[^abc]`: matches any char that is not `a` or `b` or `c`
      - `-`: any single char within specified range
        - `[A-Z]` matches any uppercase letter
        - `[0-9]` matches any digit
        - `[a-zA-Z]` to match any letter regardless of case
        - `[a-zA-Z0-9]` to match any alphanumeric character
- `GROUP BY`: group records based on one or more columns
  - joins rows with same attribute (the column chosen)
  - `SELECT column1, COUNT(*) FROM table_name GROUP BY column1;`
- `HAVING`: similar to the `WHERE` clause, but
  - operates on groups after they've been grouped, rather than rows before they've been grouped
  - IMPORTANT: you can't use `HAVING` before `GROUP BY`
- `ORDER BY`: order records based on one or more columns
  - limit: restrict the number of matches

- `INSERT`: add new rows to a table
  - data types: syntax is different depending on the DB (mysql, postgresql, etc)
    - numeric
      - data types that exist on mysql and postgresql
        - SMALLINT
        - BIGINT
        - INTEGER
      - only mysql
        - DECIMAL
        - FLOAT
        - DOUBLE
      - only postgresql
        - NUMERIC(precision, scale) precision: total number of digits that can be stored
          - scale: maximum number of digits to the right of the decimal point
        - REAL
        - DOUBLE PRECISION
    - character string: CHAR, VARCHAR
    - binary: binary
    - boolean: BOOLEAN
    - data and time: data, datetime
- `UPDATE`: modify existing data in a table
- `DELETE`: remove data from a table


- `CREATE TABLE`: create new table in the database
  - `FOREIGN KEY`: creates a column with the values of the column of another table
    - `CASCADE`: defines what happens to foreign keys when the reference is changed
      - `ON DELETE`: when a record in the primary table is deleted, any records in the foreign key table that reference the deleted primary key will also be deleted
      - ON UPDATE: changes on primary table also apply to child tables
- `ALTER table`: modify existing table in the database
- `DROP table`: delete a table and its data from the database
- `TRUNCATE table`: delete all rows from a table without deleting the table itself

- `CREATE database`: create new database
- `DROP database`: delete a database and all its contents

set operations:
- `UNION`
- `INTERSECT`
- `MINUS` or `EXCEPT`

TODO

restrictions:
- table restrictions
  - column restrictions
    - `NULL`
      - you can't compare `NULL` with equal sign (e.g. `WHERE NOME=NULL`)
      - you should use `IS` keywork (e.g. `WHERE NOME IS NULL`)
    - `NOT NULL`
    - `UNIQUE`
    - `FOREIGN KEY`
    - `CHECK` table restrictions
- assertions
- domain restrictions
  - check

- constraints: enforce rules on table's data
  - `PRIMARY KEY`: value can't be NULL and has to be UNIQUE
  - `FOREIGN KEY`: links the value of a column in table1 to the value of another column in table2
    - `CASCADE`: explained before
  - `UNIQUE`: all values in the column are distinct
  - `CHECK`: all values in the column satisfy a condition
    - `CONSTRAINT chk_stock CHECK (stock_quantity >= 0)`
  - `DEFAULT`: defines the initial value of a column when inserting a new row
    - `balance DECIMAL(10, 2) DEFAULT 0.00`
  - `NOT NULL`: a column cannot have NULL value
    - `email VARCHAR(100) NOT NULL`

- `ALIAS`: temporary name for column
  - `select firstname as "first name", lastname as "last name" from employees;`
- `DISTINCT`: removes duplicate rows from the results of a query
  - `SELECT DISTINCT City FROM Customers;`

- aggregate functions: return only 1 result
  - `SUM()`: return the sum of all values in the column
  - `AVG()`: return the average value of the column
  - `MIN()`: return minimum value
  - `MAX()`: return maximum value
  - `COUNT(column_or_table)`: count number of rows in result set with non-NULL values in specified column or table

> [!IMPORTANT]
> aggregate functions CAN be used inside: `SELECT`, `HAVING`, `ORDER BY`
> aggregate functions CANNOT be used inside: `FROM`, `WHERE`

## Subqueries

- subqueries queries:
  - `IN`: filter results from another sql query
    - This operator checks if a value matches any value in a list and is suitable for one-to-many comparisons
  - `EXISTS`
  - `ALL`
  - `SOME`
  - `ANY`

Use `IN` when the subquery may return one or more values. 

```sql
SELECT id, song_name, artist_id
FROM songs
WHERE artist_id IN (
    SELECT id
    FROM artists
    WHERE artist_name LIKE 'Rick%'
);
```

Use `=` when the subquery is expected to return a single value. If the subquery returns more than one value, the query will fail with an error. This operator is used for one-to-one comparisons

```sql
SELECT *
FROM transactions
WHERE user_id = (
    SELECT id
    FROM users
    WHERE name = 'David'
);
```
## Joins

| types of join   | join conditions    |
|--------------- | --------------- |
| INNER JOIN   | NATURAL   |
| LEFT OUTER JOIN   | ON PREDICATE   |
| RIGHT OUTER JOIN   | USING (A1, A2, ...)   |
| FULL OUTER JOIN   | --- |

- natural join: automatically joins matching values
  - simpler join operation
- inner join: intersection
- left join: left table
- right join: right table
- full join: L U R

### INNER JOIN

```sql
SELECT *
FROM employees
INNER JOIN departments
ON employees.department_id = departments.id;
```

```sql
SELECT students.name, classes.name
FROM students
INNER JOIN classes on classes.class_id = students.class_id;
```

- table_name.column_name
- The `ON` clause specifies the condition to join the tables
    - If the columns on the `ON` clause have the same name, the column won't appear twice in the result
- When joining tables, if there are columns with same name, they will appear twice in the result

### LEFT OUTER JOIN = LEFT JOIN

```sql
SELECT e.name, d.name
FROM employees e
LEFT JOIN departments d
ON e.department_id = d.id;
```

### RIGHT OUTER JOIN (not supported by SQLite)

TODO

### FULL OUTER JOIN = FULL JOIN (not supported by SQLite)

TODO

## View

> alias for a SQL query that can be treated as a table

```sql
CREATE VIEW it_employees AS SELECT id, first_name, last_name, salary
FROM employees
WHERE department = 'IT';
```

## Performance

### Index

- creates a binary tree
- faster to look up values in a column
  - o(log n)
- primary keys are indexed by default
- it's fairly common to name an index after the column it's created on with a suffix of `_idx`
- you shouldn't index too many columns
  - indexes take up space
  - create performance overhead when inserting or updating data
  - each time you insert a record, that record needs to be added to many trees

```sql
CREATE INDEX email_idx ON users(email);
```

### Multi-column indexes

- speed up look ups that depend on multiple columns
- only add multi-column indexes if you're doing frequent lookups on a specific combination of columns
- A multi-column index is sorted by the first column first, the second column next, and so forth
- first column lookup in multi-column index has same performance as a single-column index
- 2nd or 3rd columns lookups in multi-column index have bad performance compared to single-column lookups

```sql
CREATE INDEX first_name_last_name_age_idx
ON users (first_name, last_name, age);
```

## Examples

```sql
CREATE TABLE employees(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER,
    is_manager BOOLEAN,
    salary INTEGER
);
```

```sql
SELECT employee_name, salary
FROM employees
WHERE salary BETWEEN 30000 and 60000;
```


```sql
SELECT product_name, quantity
FROM products
WHERE quantity NOT BETWEEN 20 and 100;
```

```sql
SELECT *
FROM Orders
WHERE OrderDate BETWEEN '01/07/1996' AND '31/07/1996';
```

```sql
SELECT name, price, quantity
FROM products
ORDER BY price DESC;
```

```sql
ALTER TABLE employees
RENAME TO contractors;
```

```sql
ALTER TABLE employees
ADD birthdate DATE;
```

```sql
ALTER TABLE contractors
RENAME COLUMN salary TO invoice;
```

```sql
SELECT id, name
WHERE id = 1
FROM users;
```

```sql
CREATE TABLE departments (
    id INTEGER PRIMARY KEY,
    department_name TEXT NOT NULL
);
```

```sql
CREATE TABLE employees (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    department_id INTEGER,
    CONSTRAINT fk_departments
    FOREIGN KEY (department_id)
    REFERENCES departments(id)
);
```

```sql
SELECT COUNT(*) FROM tableName;
```

```sql
SELECT COUNT(columnName) FROM tableName WHERE condition;
```

```sql
SELECT SUM(columnName) FROM tableName;
```

```sql
SELECT SUM(columnName) FROM tableName WHERE condition;
```

```sql
SELECT max(price)
FROM products;
```

```sql
SELECT product_name, min(price)
FROM products;
```


```sql
SELECT AVG(columnName) FROM tableName;
```

```sql
SELECT AVG(columnName) FROM tableName WHERE condition;
```

```sql
SELECT album_id, count(song_id)
FROM songs
GROUP BY album_id;
```

```sql
SELECT album_id, count(id) as count
FROM songs
GROUP BY album_id
HAVING count > 5;
```

```sql
CREATE TABLE Customers (
    CustomerID INT PRIMARY KEY,
    Name VARCHAR(100)
);

CREATE TABLE Orders (
    OrderID INT PRIMARY KEY,
    CustomerID INT,
    Product VARCHAR(100),
    FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID) ON DELETE CASCADE
);
```

```sql
TRUNCATE TABLE Employees;
```

```sql
CREATE TABLE Departments (
    DeptID INT PRIMARY KEY,
    DeptName VARCHAR(100)
);

CREATE TABLE Employees (
    EmpID INT PRIMARY KEY,
    EmpName VARCHAR(100),
    DeptID INT,
    FOREIGN KEY (DeptID) REFERENCES Departments(DeptID) ON UPDATE CASCADE
);
```

example combining several constraints:

```sql
CREATE TABLE employees (
    employee_id INT PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE,
    department_id INT,
    manager_id INT,
    salary DECIMAL(10, 2) CHECK (salary > 0),
    hire_date DATE DEFAULT CURRENT_DATE,
    CONSTRAINT fk_department
        FOREIGN KEY (department_id) REFERENCES departments (department_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,
    CONSTRAINT fk_manager
        FOREIGN KEY (manager_id) REFERENCES employees (employee_id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);
```

```sql
SELECT product_name, shipment_status
FROM products
WHERE shipment_status IN ('shipped', 'preparing', 'out of stock');
```

```sql
SELECT EmployeeName
FROM Employees
WHERE EmployeeID IN (SELECT EmployeeID FROM Orders WHERE OrderDate > '2020-01-01');
```

```sql
SELECT email
FROM employees
INTERSECT
SELECT email
FROM contractors;
```

```sql
SELECT email
FROM employees
UNION
SELECT email
FROM contractors;
```

```sql
SELECT email
FROM employees
MINUS
SELECT email
FROM contractors;
```

---

## theory

- grau de tabela: nº of columns in table
- cardinality: nº of rows in table

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

### relational algebra

- Selection (σ)
  - example: σ(age > 30)(employees)
  - ```SELECT * FROM employees WHERE age > 30; ```
  - OBS: σ => Sigma => Select * => sElEct => whErE
- Projection (π)
  - example: π(name, age)(employees)
  - ```SELECT name, age FROM employees;```
  - OBS: π => PI => PIck => Projection
- Union (∪)
  - example: employees ∪ contractors
  - ```SELECT email FROM employees UNION SELECT email FROM contractors;```
- Intersection (∩)
  - example: employees ∩ contractors
  - ```SELECT email FROM employees INTERSECT SELECT email FROM contractors;```
- Difference (−)
  - example: employees − contractors
  - ```SELECT email FROM employees EXCEPT SELECT email FROM contractors;```
- Rename (ρ)
  - example: ρ(name/employee_name)(employees)
  - ```SELECT name AS employee_name FROM employees;```
  - OBS: ρ => Rho => Rename
- Cartesian Product (×)
  - example: employees × departments
- Join (⨝ or ٭)
  - example: employees ⨝ employees.department_id = departments.department_id departments


### Normal Forms and Data Normalization

```
BCNF ⊂ 3NF ⊂ 2NF ⊂ 1NF
```

> [!IMPORTANT]
> primary key in SQL != primary key in data normalization/normal forms
> in data normalization, primary key is the collections that uniquely identifies a row
> in SQL, primary key is a single column that uniquely identifies a row

#### First Normal Form (1NF)

- it must have a unique primary key
- a cell can't have a nested table as its value
  - a cell can't have multiple values
    - may not even be possible

> [!TIP]
> 1NF = 1 value per cell

#### Second Normal Form (2NF)

- table must be in 1NF
- all non-key attributes must be fully dependent on the entire primary key
  - non-key attributes: columns that are not part of the primary key

> [!IMPORTANT]
> tables can have composite primary keys (multiple columns that collectively functions as a single primary key

#### Third Normal Form (3NF)

- table must be in 2NF
- all columns that aren't part of the primary key are dependent solely on the primary key
- columns can't depend on other non-key attributes

#### Boyce-Codd Normal Form (BCNF)

- a column that's part of a primary key can't be entirely dependent on a column that's not part of that primary key

### Rules of thumb for database design

1. Every table should always have a unique identifier (primary key)
2. 90% of the time, that unique identifier will be a single column named id
3. Avoid duplicate data
4. Avoid storing data that is completely dependent on other data. Instead, compute it on the fly when you need it
5. Keep your schema as simple as you can. Optimize for a normalized database first. Only denormalize for speed's sake when you start to run into performance problems
