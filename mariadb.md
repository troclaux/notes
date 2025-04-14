
# MariaDB

> open-source relational database management system

- can replace MySQL, but is maintained separately from MySQL

## cli commands

- login: `sudo mariadb -u root -p`
  - `-u root`: login as root user
  - `-p`: prompt for password

- dump a database: `mysqldump -u root -p mydb > mydb_backup.sql`
- restore a database: `mysql -u root -p mydb < mydb_backup.sql`
