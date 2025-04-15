
# MongoDB

> NoSQL document-based database

- stores data in JSON-like documents
  - e.g. `{ "name": "John", "age": 30 }`
- schema-less, each document can have a different structure
- scales horizontally

| MongoDB | relational DB |
|---------|-----|
| Documents | Rows |
| Collections | Tables |
| Schema-less | Predefined schema |
| Horizontal scaling | Mostly vertical |
| Good for nested data | Requires JOINs |

## CLI commands

- connect o mongodb: `mongosh`
- show all database: `show dbs`
- list running databases and their stats: `db.stats()`
- use or create a database: `use mydb`
- create or insert a document: `db.users.insertOne({ name: "Alice", age: 30 })`
  - this command automatically create the `users` collection
- find a document: `db.users.find({ name: "Alice" })`
- show collections: `show collections`
- delete a document: `db.users.deleteOne({ name: "Alice" })`

- queries:
  - find all users: `db.users.find()`
  - find first match of a user: `db.users.findOne({ name: "Alice" })`
  - find user with filters: `db.users.find({ age: { $gt: 25 } })   // age > 25`
  - update user: `db.users.updateOne( { name: "Alice" }, { $set: { age: 31 } })`

- export/dump a database: `mongodump --db nome_do_banco --out ~/backups/mongodb`
- import/restore a database: `mongorestore --db nome_do_banco_novo /caminho/do/backup/nome_do_banco_original`


