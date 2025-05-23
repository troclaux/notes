[databases](./databases.md)
[postgresql](./postgresql.md)
[MySQL](./mysql.md)
[MongoDB](./mongodb.md)

# SQL (Structured Query Language)

> standard language for accessing and manipulating databases

> [!IMPORTANT]
> although SQL is an ANSI/ISO standard, there are different versions of the SQL language
> most of the SQL database programs also have their own proprietary extensions in addition to the SQL standard
> each database uses its own variation of SQL

- basic operations:
  - execute queries against a database
  - retrieve data from a database
  - insert records in a database
  - update records in a database
  - delete records from a database
  - create new databases
  - create new tables in a database
  - create stored procedures in a database
  - create views in a database
  - set permissions on tables, procedures, and views

order of operations in SQL query:

```
FROM => JOIN => WHERE => GROUP BY => HAVING => SELECT => ORDER BY => LIMIT
```

- `SELECT`: retrieve data from database
  - `*`: select all
- `FROM`: select the target(s) table(s)
  - list the tables from where data is fetched
  - able to create an alias for a table: `FROM users u`
- `WHERE`: filter records based on conditions
  - `BETWEEN`: selects values within given range (values can be numbers, text or dates)
  - `LIKE`: filter records that matches the string
    - wildcards: used for pattern matching
      - `%`: zero or more chars
        - `SELECT * FROM products WHERE product_name LIKE 'banana%';`
      - `_`: single char
        - `SELECT * FROM products WHERE product_name LIKE '__oot';`
- `GROUP BY`: group records based on one or more columns
  - joins rows of a column that have same value
  - you must pass HOW to group the other values in the remaining columns
  - `SELECT column1, COUNT(*) FROM table_name GROUP BY column1;`
- `HAVING`: similar to the `WHERE` clause, but
  - operates on groups after they've been grouped, rather than rows before they've been grouped
  - IMPORTANT: you can't use `HAVING` before `GROUP BY`
- `ORDER BY`: order records based on one or more columns
- `LIMIT`: limit the number of rows returned
  - `LIMIT 3;`: show only the first 3 results
- `OFFSET`: skip a number of rows before starting to return results
  - `OFFSET 5;`: skip first 5 results and show the rest

- `INSERT`: add new rows to a table
  - data types: syntax is different depending on the DB (mysql, postgresql, etc)
    - numeric
      - data types that exist on mysql and postgresql
        - `NUMERIC(precision, scale)`
          - precision: total number of digits that can be stored
          - scale: maximum number of digits to the right of the decimal point
        - INTEGER
        - DECIMAL
        - FLOAT
        - SMALLINT
        - BIGINT
        - DOUBLE PRECISION
      - only mysql
        - DOUBLE
      - only postgresql
        - REAL
    - character string: CHAR, VARCHAR
    - binary: BINARY(32)
    - boolean: BOOLEAN
    - date and time: `DATE`, `TIME`, `DATETIME`, `TIMESTAMP`
- `UPDATE`: modify existing data in a table
- `DELETE`: remove data from a table

- `CREATE TABLE`: create new table in the database
  - `FOREIGN KEY`: creates a column with the values of a column in another table
    - `CASCADE`: defines what happens to dependent rows in the foreign key table when the referenced row is updated or deleted
      - `ON DELETE`: when a record in the primary table is deleted, any records in the foreign key table that reference the deleted primary key will also be deleted
      - `ON UPDATE`: changes on primary table also apply to child tables
- `ALTER table`: modify existing table in the database
- `DROP table`: delete a table and its data from the database
- `TRUNCATE table`: delete all rows from a table without deleting the table itself
- `RENAME`: `ALTER TABLE users RENAME TO customers;`
  - postgresql syntax

> [!WARNING]
> when creating a table, the last attribute should not have a comma after it, just like in the example below

```sql
CREATE TABLE Customers (
    CustomerID INT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(100) UNIQUE
);
```

- `CREATE DATABASE`: create new database
- `DROP mydatabase`: delete a database and all its contents

- set operations: combine results from two `SELECT` statements (they must have the same number of columns and compatible types)
  - `UNION`: combines and remove duplicates
  - `INTERSECT`: returns only rows that appear in both result sets
  - `EXCEPT`(postgresql) or `MINUS`(oracle): returns rows from the first set that aren't in the second

```sql
SELECT name FROM students
UNION
SELECT name FROM teachers;
```

- restrictions: enforce rules on tables/columns to ensure data integrity
  - table restrictions
    - column restrictions
      - `NULL`
        - you can't compare `NULL` with equal sign (e.g. `WHERE NOME=NULL`)
        - use `IS` keywork (e.g. `WHERE NOME IS NULL`)
      - `NOT NULL`
      - `UNIQUE`
      - `FOREIGN KEY`
      - `CHECK` enforces a condition on values in a column
  - assertions: apply conditions across the whole table

- constraints: enforce rules on table's data
  - `PRIMARY KEY`: value can't be NULL and has to be UNIQUE
    - the primary key can be one column or a set of columns (e.g. `PRIMARY KEY (user_id, order_id)`)
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
  - `SELECT DISTINCT city FROM customers;`

- aggregate functions: return only 1 result
  - `SUM()`: return the sum of all values in the column
  - `AVG()`: return the average value of the column
  - `MIN()`: return minimum value
  - `MAX()`: return maximum value
  - `COUNT(column_or_table)`: count number of rows in result set with non-NULL values in specified column or table

> [!IMPORTANT]
> aggregate functions CAN be used inside: `SELECT`, `HAVING`, `ORDER BY`
> aggregate functions CANNOT be used inside: `FROM`, `WHERE`

- `AS`: rename a table or column temporarily in a query
  - column aliasing: `SELECT name AS full_name FROM users;`
  - table aliasing: `SELECT u.name FROM users AS u;`
- `ON`: tells sql which columns to compare to join rows from two tables
  - defines how to join rows by matching column values
  - use cases:
    - `JOIN`
    - `CREATE TRIGGER`
    - `FOREIGN KEY`
    - `GRANT`/`REVOKE`
    - `CREATE INDEX`

```sql
SELECT *
FROM orders o
JOIN customers c ON o.customer_id = c.id;
```

## named constraints

> rules that allow data to be entered into a table only if it meets the predefined conditions

- examples
  - `NOT NULL`: ensures a column cannot have NULL value
  - `UNIQUE`: ensures all values in a column are different
    - can also be used to ensure a combination of columns is unique, also called "composite unique constraint"
  - `PRIMARY KEY`: uniquely identifies each row in a table
  - `FOREIGN KEY`: ensures referential integrity between tables
  - `DEFAULT`: sets a default value for a column
  - `CHECK`: ensures all values in a column satisfy certain conditions

- adding a constraint: `ALTER TABLE users ADD CONSTRAINT email_unique UNIQUE (email);`
  - another example: `ALTER TABLE users ADD CONSTRAINT primary_key PRIMARY KEY (city, name, id);`
- renaming a constraint: `ALTER TABLE users RENAME CONSTRAINT email_unique TO unique_email_address;`
- removing a constraint: `ALTER TABLE users DROP CONSTRAINT email_unique;`
- naming a constraint: `ALTER TABLE users ADD CONSTRAINT fk_city FOREIGN KEY (city_id) REFERENCES cities(id);`

```sql
CREATE TABLE workout_exercises (
  id UUID PRIMARY KEY,
  workout_id UUID NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
  exercise_id UUID NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
  sets INTEGER,
  reps INTEGER,
  order_number INTEGER,
  CONSTRAINT unique_workout_exercise UNIQUE (workout_id, exercise_id)
);
```

> [!IMPORTANT]
> try to name constraints when creating them

## subqueries

- `IN`: filter results from another sql query
  - checks if a value exists in a list or subquery result
  - use when the subquery may return multiple values
- `EXISTS`
- `ALL`
- `SOME`
- `ANY`

```sql
SELECT id, song_name, artist_id
FROM songs
WHERE artist_id IN (
    SELECT id
    FROM artists
    WHERE artist_name LIKE 'Rick%'
);
```

- use `=` when the subquery is expected to return a single value
  - if the subquery returns more than one value, the query will fail with an error
  - this operator is used for one-to-one comparisons

```sql
SELECT *
FROM transactions
WHERE user_id = (
    SELECT id
    FROM users
    WHERE name = 'David'
);
```
## join

> combines data from multiple tables into a single result set

- combines related data
- reduces data redundancy
- simplify queries

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

> returns only rows where there's a match in both tables

```sql
SELECT *
FROM orders o
INNER JOIN customers c ON o.customer_id = c.id;
```

```sql
SELECT students.name, classes.name
FROM students
INNER JOIN classes on classes.class_id = students.class_id;
```

- `table_name.column_name`
- the `ON` clause specifies the condition to join the tables
  - if the columns on the `ON` clause have the same name, the column won't appear twice in the result
- when joining tables, if there are columns with same name, they will appear twice in the result

### LEFT OUTER JOIN = LEFT JOIN

> returns all rows from the left table, plus matching rows from the right table (NULLs if no match)

- has the same number of rows as the left table
- adds matching data from the right table (or `NULL` if there's no match)

```sql
SELECT e.name, d.name
FROM employees e
LEFT JOIN departments d
ON e.department_id = d.id;
```

### RIGHT OUTER JOIN (not supported by SQLite)

> opposite of left join

- has the same number of rows as the right table
- adds matching data from the left table (or `NULL` if there's no match)

- not supported in SQLite

example: returns all departments, even if no employee is assigned

```sql
SELECT employees.name, departments.name
FROM employees
RIGHT JOIN departments ON employees.dept_id = departments.id;
```

### FULL OUTER JOIN = FULL JOIN (not supported by SQLite)

> returns all rows from both tables, with NULLs where there's no match on either side

```
number of rows in result = rows that matched + rows only in LEFT table with no match in RIGHT + rows only in RIGHT table with no match in LEFT
```

- not supported in SQLite and MySQL

example: returns all employees and all departments. If there's no match, it still shows the row with `NULL` values

```sql
SELECT employees.name, departments.name
FROM employees
FULL OUTER JOIN departments ON employees.dept_id = departments.id;
```

## trigger

- events that can trigger: `SELECT`, `UPDATE`, `DELETE`
  - can happen `BEFORE` or `AFTER` event

```sql
CREATE TRIGGER log_insert
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION log_user_insert();
```

## view

> alias for a SQL query that can be treated as a table

```sql
CREATE VIEW it_employees AS SELECT id, first_name, last_name, salary
FROM employees
WHERE department = 'IT';
```

## performance

### index

> performance optimization structure that allows the database to retrieve rows more quickly

- creates a binary tree
- faster to look up values in a column
  - O(log n)
- slower write operations
- primary keys are indexed by default
- it's fairly common to name an index after the column it's created on with a suffix of `_idx`
- you shouldn't index too many columns
  - indexes take up space
  - create performance overhead when inserting or updating data
  - each time you insert a record, that record needs to be added to many trees

```sql
CREATE INDEX email_idx ON users(email);
```

- remove index query
  - postgresql: `DROP INDEX email_idx;`
  - mysql: `ALTER TABLE your_table DROP INDEX index_name;`

### multi-column indexes

- speed up look ups that depend on multiple columns
- only add multi-column indexes if you're doing frequent lookups on a specific combination of columns
- A multi-column index is sorted by the first column first, the second column next, and so forth
- first column lookup in multi-column index has same performance as a single-column index
- 2nd or 3rd columns lookups in multi-column index have bad performance compared to single-column lookups

```sql
CREATE INDEX first_name_last_name_age_idx
ON users (first_name, last_name, age);
```

## transaction control

> manage changes made by a transaction and endure database integrity

- use transaction control (`BEGIN`, `COMMIT`, `ROLLBACK`) when:
  - you’re doing multiple related updates that must succeed or fail together
  - you’re updating multiple tables and want consistency
  - you want to recover from errors during the process (e.g., using SAVEPOINT)

- `BEGIN`/`START TRANSACTION`: starts a new transaction
  - use case: when you want multiple SQL statements to be executed as a single atomic unit
    - atomic unit: either all of the sql statements succeed or none of them do

- `COMMIT`: makes all changes made in the current transaction **permanent**
  - use case: after a set of successful operations, use `COMMIT` to save them to the database

```sql
BEGIN;

INSERT INTO accounts (user_id, balance) VALUES (1, 1000);
UPDATE accounts SET balance = balance - 200 WHERE user_id = 1;

COMMIT;
```

- `ROLLBACK`: undo changes made in the current transaction
  - if something goes wrong, cancel all operations since the last `BEGIN`

```sql
BEGIN;

UPDATE accounts SET balance = balance - 200 WHERE user_id = 1;

-- Something goes wrong
ROLLBACK;
```

- `SAVEPOINT`: sets a named point within a transaction that you can roll back to without rolling back the entire transaction
  - use case: useful in large transactions where you may want to undo only part of it

```sql
BEGIN;

INSERT INTO accounts (user_id, balance) VALUES (2, 500);
SAVEPOINT after_insert;

UPDATE accounts SET balance = balance / 0 WHERE user_id = 2; -- This will cause an error

ROLLBACK TO after_insert; -- Undo only the update, keep the insert

COMMIT; -- Save the insert
```

## queries

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
SELECT id, name
FROM users
WHERE id = 1;
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

delete all records in users table:

```sql
DELETE FROM users;
```

```sql
DELETE FROM users
WHERE email = 'example@email.com';
```

> [!IMPORTANT]
> use `=` when comparing non-NULL values
> use `IS NULL` or `IS NOT NULL` only when checking for NULLs

```sql
DELETE FROM users
WHERE email IS NULL;
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
INSERT INTO employees (name, age, is_manager, salary)
VALUES ('John Doe', 30, true, 50000);
```

```sql
ALTER TABLE contractors
RENAME COLUMN salary TO invoice;
```

```sql
UPDATE users
SET name = 'John Doe', salary = 20000
WHERE id = 5;
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
    FOREIGN KEY (department_id) REFERENCES departments(id)
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
