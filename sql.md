# SQL

## Create table

```sql
CREATE TABLE employees(
    id INTEGER,
    name TEXT,
    age INTEGER,
    is_manager BOOLEAN,
    salary INTEGER
);
```

## Alter table

```sql
ALTER TABLE employees
RENAME TO contractors;
```

```sql
ALTER TABLE contractors
RENAME COLUMN salary TO invoice;
```

## Basic query

```sql
SELECT id, name
WHERE id = 1
FROM users;
```

## Primary and Foreign keys

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

## Functions

### COUNT()

```sql
SELECT COUNT(*) FROM tableName;
```

```sql
SELECT COUNT(columnName) FROM tableName WHERE condition;
```

### SUM()

```sql
SELECT SUM(columnName) FROM tableName;
```

```sql
SELECT SUM(columnName) FROM tableName WHERE condition;
```

## Aggregating functions

### MAX() MIN()

```sql
SELECT max(price)
FROM products;
```

```sql
SELECT product_name, min(price)
FROM products;
```

### AVG()

```sql
SELECT AVG(columnName) FROM tableName;
```

```sql
SELECT AVG(columnName) FROM tableName WHERE condition;
```

### GROUP BY

```sql
SELECT album_id, count(song_id)
FROM songs
GROUP BY album_id;
```

### HAVING

The `HAVING` clause is similar to the `WHERE` clause, but it operates on groups after they've been grouped, rather than rows before they've been grouped

## Subqueries

Use `IN` when the subquery may return one or more values. This operator checks if a value matches any value in a list and is suitable for one-to-many comparisons

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

```sql
SELECT album_id, count(id) as count
FROM songs
GROUP BY album_id
HAVING count > 5;
```

## Normal forms

- BCNF contains 3NF, which contains 2NF, which contains 1NF

### First normal form (1NF)

- It must have a unique primary key
A cell can't have a nested table as its value (depending on the database you're using, this may not even be possible)

### Second normal form (2NF)

- All columns that are not part of the primary key are dependent on the entire primary key, and not just one of the columns in the primary key

### Third normal form (3NF)

- All columns that aren't part of the primary are dependent solely on the primary key

### Boyce-Codd normal form (BCNF)

- A column that's part of a primary key can not be entirely dependent on a column that's not part of that primary key

## Rules of thumb for database design

1. Every table should always have a unique identifier (primary key)
2. 90% of the time, that unique identifier will be a single column named id
3. Avoid duplicate data
4. Avoid storing data that is completely dependent on other data. Instead, compute it on the fly when you need it
5. Keep your schema as simple as you can. Optimize for a normalized database first. Only denormalize for speed's sake when you start to run into performance problems


## Joins

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

### LEFT JOIN

```sql
SELECT e.name, d.name
FROM employees e
LEFT JOIN departments d
ON e.department_id = d.id;
```

### FULL JOIN (not supported by SQLite)



## NULL values

```sql
SELECT name
FROM users
WHERE first_name IS NULL;
```

```sql
SELECT name
FROM users
WHERE first_name IS NOT NULL;
```

## BETWEEN

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

## BETWEEN

```sql
SELECT DISTINCT previous_company
FROM employees;
```

## IN

```sql
SELECT product_name, shipment_status
FROM products
WHERE shipment_status IN ('shipped', 'preparing', 'out of stock');
```

## LIKE

% wildcard operator matches zero or more characters

```sql
SELECT *
FROM products
WHERE product_name LIKE 'banana%';
```

_ wildcard operator matches a single character

```sql
SELECT *
FROM products
WHERE product_name LIKE '__oot';
```

## LIMIT

```sql
SELECT *
FROM products
WHERE product_name LIKE '%berry%'
LIMIT 50;
```

## ORDER BY

```sql
SELECT name, price, quantity
FROM products
ORDER BY price DESC;
```

### Left join

```sql
SELECT e.name, d.name
FROM employees e
LEFT JOIN departments d
ON e.department_id = d.id;
```

## Performance

### Index

- Creates a binary tree
- Faster to look up values in a column
    - O(log n)
- Primary keys are indexed by default
- It's fairly common to name an index after the column it's created on with a suffix of `_idx`
- You shouldn't index too many columns
    - Indexes take up space
    - Create performance overhead when inserting or updating data
    - Each time you insert a record, that record needs to be added to many trees

```sql
CREATE INDEX email_idx ON users(email);
```

### Multi-column indexes

- Speed up look ups that depend on multiple columns
- Only add multi-column indexes if you're doing frequent lookups on a specific combination of columns

A multi-column index is sorted by the first column first, the second column next, and so forth. A lookup on only the first column in a multi-column index gets almost all of the performance improvements that it would get from its own single-column index. However, lookups on only the second or third column will have very degraded performance

```sql
CREATE INDEX first_name_last_name_age_idx
ON users (first_name, last_name, age);
```
