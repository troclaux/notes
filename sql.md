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

## Joins

### Inner join

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
- The ON clause specifies the condition to join the tables
    - If the columns on the **ON** clause have the same name, the column won't appear twice in the result
- When joining tables, if there are columns with same name, they will appear twice in the result

### Left join

```sql
SELECT e.name, d.name
FROM employees e
LEFT JOIN departments d
ON e.department_id = d.id;
```
