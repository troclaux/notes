[sql](./sql.md)
[databases](./databases.md)

# MongoDB

> NoSQL document-based database

- stores data in JSON-like documents
  - e.g. `{ "name": "John", "age": 30 }`
- schema-less, each document can have a different structure
- scales horizontally
- default port is 27017
- `mongosh`: official CLI used to interact with MongoDB databases
  - replaced the older shell (`mongo`) since MongoDB 5.0+
  - used to:
    - connect to MongoDB server
    - run database commands using javascript
    - query and modify data
    - manage databases

| MongoDB | relational DB |
|---------|-----|
| Documents | Rows |
| Collections | Tables |
| Fields | Columns |
| Schema-less | Predefined schema |
| Horizontal scaling | Mostly vertical |
| Good for nested data | Requires JOINs |

document example in MongoDB:

```javascript
{
  "name": "Alice",
  "age": 25,
  "email": "alice@example.com"
}
```

- `"name"`, `"age"` and `"email"` are fields
- this whole object is a **document**
- it belongs to a **collection**

## CLI commands

- connect to mongodb: `mongosh`
- connect to mongodb with connection string: `mongosh "mongodb://<host>:<port>/<database>"`
- show all database: `show dbs`
- show current database: `db`
- use or create a database: `use mydb`

- list running databases and their stats: `db.stats()`
- drop/delete a database: `db.dropDatabase()`

- show all collections: `show collections`
- create a collection manually: `db.createCollection("myCollection")`
- drop/delete a collection: `db.myCollection.drop()`

- create/insert a document: `db.users.insertOne({ name: "Alice", age: 30 })`
  - this command automatically create the `users` collection if it doesn't exist
- find a document: `db.users.find({ name: "Alice" })`
- delete a document: `db.users.deleteOne({ name: "Alice" })`
- delete many documents: `db.myCollection.deleteMany({ age: { $lt: 30 } })`
- find all documents: `db.myCollection.find()`
  - pretty-print results: `db.myCollection.find().pretty()`
  - limit number of results: `db.myCollection.find().limit()`
  - sort results: `db.myCollection.find().sort()`
- find first match of a document: `db.myCollection.findOne({ name: "Alice" })`
- find document with filters: `db.myCollection.find({ age: { $gt: 25 } })   // age > 25`
- update document: `db.myCollection.updateOne( { name: "Alice" }, { $set: { age: 31 } })`
- count documents: `db.myCollection.countDocuments()`

insert many documents:

```javascript
db.myCollection.insertMany([
  { name: "Bob", age: 30 },
  { name: "Charlie", age: 28 }
])
```

update many documents:

```javascript
db.myCollection.updateMany(
  { age: { $gt: 25 } },
  { $set: { status: "senior" } }
)
```

- export/dump a database: `mongodump --db nome_do_banco --out ~/backups/mongodb`
- import/restore a database: `mongorestore --db nome_do_banco_novo /caminho/do/backup/nome_do_banco_original`

