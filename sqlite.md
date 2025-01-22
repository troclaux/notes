
[databases](/databases.md)
[sql](/sql.md)

# sqlite

> sql database

- lightweight: self-contained, single-file database that doesn't require separate server process
  - self-contained: stores entire database in a single-file on disk
  - single-file: stores the database in a single file on disk, rather than in a separate database server

## basics

- accessing a database with sqlite shell: `sqlite3 my_database.db`
  - if `my_database.db` doesn't exist, sqlite will create it automatically
- exit sqlite shell: `.quit`
- show tables: `.tables`
- show schema: `.schema users`

import data from csv:

```bash
.mode csv
.import file.csv users
```

export data to a csv:

```bash
.headers on
.mode csv
.output output.csv
SELECT * FROM users;
.output stdout
```
