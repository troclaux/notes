
# ORM (Object-Relational Mapping)

> programming technique that allows developers to interact with a database using an object-oriented paradigm

- examples:
  - prisma (recommended)
    - typescript
    - supports postgresql, mysql, mongodb, sqlite
    - simpler
  - TypeORM
    - typescript
  - sequelize
    - javascript
    - supports postgresql, mysql, sqlite, mongodb

## prisma

1. install prisma

```bash
npm install @prisma/client @prisma/cli --save-dev
```

2. initialize prisma

```bash
npx prisma init
```

- this command creates a `.env` file and a `prisma/schema.prisma` file

3. define database schema

modify `prisma/schema.prisma` with your entities

4. run migrations

```bash
npx prisma migrate dev --name init
```

5. sync prisma with the database

```bash
npx prisma generate
```
