[sql](./sql.md)
[databases](./databases.md)

# PostgreSQL

> open-source relational database management system

[common PostgreSQL mistakes](https://wiki.postgresql.org/wiki/Don%27t_Do_This)

## data types

- numeric types
  - `numeric`: stores numbers
- boolean types
  - `boolean`
- character types
  - `CHAR(10)`: fixed-length character string
  - `VARCHAR(200)`: variable-length character string
  - `TEXT`: variable unlimited-length string
- date/time types
  - `DATE`: date (no time)
  - `TIME`: time (no date)
  - `TIMESTAMP`: date and time
  - `INTERVAL`: time interval
- binary types
  - `BYTEA`: binary data
- other types
  - `UUID`: universally unique identifier
  - `JSON`: JSON data
  - `JSONB`: binary JSON data
  - `ARRAY`: array of other types
  - `ENUM`: enumerated list of values

## setup

[install PostgreSQL](https://docs.fedoraproject.org/en-US/quick-docs/postgresql)

to use postgresql interactive terminal:

```bash
sudo -u postgres psql
```

- `-u postgres`: run as the postgres user
- `psql`: open the PostgreSQL prompt

to create user and database:

```
postgres=# CREATE USER lenny WITH PASSWORD 'leonard';
postgres=# CREATE USER lenny WITH PASSWORD 'leonard';
postgres=# CREATE DATABASE my_project OWNER lenny;
```

- `postgres`: the database I'm currently connected to

postgresql interactive terminal commands:

- list databases: `\l`
- list all commands: `\?`
- connect to a database: `\c dbname`
- describe a table: `\d tablename`
- describe/list tables: `\dt`
- describe/list users: `\du`
- exit postgres prompt: `\q`
- create new database: `CREATE DATABASE your_database_name;`

> [!IMPORTANT]
> use snake_case in postgresql, avoid camelCase

## connect to database from terminal

### using `psql` command-line arguments (recommended method)

```bash
psql -h localhost -U username -d database_name 5432
```

- doesn't expose the password in the command history

### using connection string

> string that contains all the information needed to connect to a database

- format depends on the database
- postgresql format:
  - `protocol://username:password@localhost:5432/dbname`
    - `protocol`: can be `postgresql` or `postgres`
    - `username:password`: the credentials
    - `localhost`: the host
    - `5432`: the port that postgresql is running on
    - `dbname`: the database name
- can be used in `.env` as credentials for the application
  - format: `DATABASE_URL=postgresql://<DB_USER>:<DB_PASSWORD>@<RDS_ENDPOINT>:5432/<DB_NAME>`
  - e.g. `DATABASE_URL=postgresql://admin:mysecurepassword@your-db-instance.cle8abcd1234.sa-east-1.rds.amazonaws.com:5432/mydatabase`

| **Component** | **Value** |
|--------------|----------|
| Protocol | `postgresql://` |
| User | `<DB_USER>` |
| Password | `<DB_PASSWORD>` |
| Host (RDS Endpoint) | `<RDS_ENDPOINT>` |
| Port | `5432` |
| Database Name | `<DB_NAME>` |

## postgresql sql keywords

- `SERIAL`: auto-incrementing integer

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);
```

- `UUID`: Universally Unique IDentifier
  - ensures each value is globally unique
- `PRIMARY KEY`: makes it the primary key for the table

- `RETURNING`: returns the values of the columns listed after the keyword

```sql
INSERT INTO employees (name, age, is_manager, salary)
VALUES ('John Doe', 30, true, 50000)
RETURNING *;
```

## error troubleshooting

### can't connect with connection string

error message after running `psql <connection string>`:

```
psql: error: connection to server at "localhost" (::1), port 5432 failed: FATAL: Ident authentication failed for user "postgres"
```

- the problem occurs because by default, postgresql is configured to use "ident" authentication for local connections. this means:
  - postgresql checks if your linux system username matches the postgresql username you're trying to use
  - if you try to connect as a "postgres" user, but you are not logged in as a "postgres" linux user, you will get this error

- the connection string uses password authentication, but the default pg_hba.conf ignores the password and tries to use ident instead
- the solution is to execute a command as `postgres` user to change the authentication method in the pg_hba.conf file

```bash
sudo -u postgres vi /var/lib/pgsql/data/pg_hba.conf
```

change this line from `pg_hba.conf`, replace `ident` with `md5`:

```
# IPv4 local connections:
host    all             all             127.0.0.1/32            ident
# IPv6 local connections:
host    all             all             ::1/128                 ident
```
