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

```
ALTER TABLE employees
RENAME TO contractors;
```

```
ALTER TABLE contractors
RENAME COLUMN salary TO invoice;
```

## Basic query

```
SELECT id, name
WHERE id = 1
FROM users;
```
