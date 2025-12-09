
# sql server

> proprietary relational database management system by microsoft

## sql server vs postgresql

- `TOP` instead of `LIMIT`: `SELECT TOP 10 * FROM users;`
- `GETDATE()` instead of `NOW()`: `SELECT GETDATE();`
- `IDENTITY(1,1)` instead `SERIAL`
    - `IDENTITY(<starting value>, <increment value>)`

```sql
CREATE TABLE test (
    id INT IDENTITY(1,1) PRIMARY KEY,
    name NVARCHAR(100)
);
```

- no boolean, uses bit: `SELECT CAST(1 AS BIT), CAST(0 AS BIT);`

---

- Book: logical grouping of positions and trades
- StrategyOptimization:
- SignalField:
