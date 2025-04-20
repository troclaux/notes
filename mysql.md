
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
