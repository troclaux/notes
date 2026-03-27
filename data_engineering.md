# data engineering

## the medallion architecture

```
bronze => silver => gold
```

| Layer | Description | Example |
|-------|-------------|---------|
| **Bronze** | Raw data, no transformations | Raw parquet with nulls, duplicates and misspelled columns → `raw` |
| **Silver** | Cleaned and standardized | Nulls handled, duplicates removed, columns renamed → `clean` |
| **Gold** | Business-ready aggregations | Daily returns per ticker, risk summary per fund, portfolio NAV series |

Each layer exists so you can reprocess downstream without re-fetching source data

### Bronze: land raw, no changes

```python
df = pd.read_csv("source_prices.csv")
df.to_parquet("bronze/prices.parquet")
```

### Silver: clean and standardize

```python
df = pd.read_parquet("bronze/prices.parquet")
df = df.drop_duplicates(subset=["id", "date"])
df = df.dropna(subset=["id"])
df = df.rename(columns={"Px_Last": "px_close", "Dt": "date"})
df["date"] = pd.to_datetime(df["date"])
df.to_parquet("silver/prices.parquet")
```

### Gold: aggregate for business use

```python
df = pd.read_parquet("silver/prices.parquet")
df = df.sort_values(["id", "date"])
df["log_return"] = df.groupby("id")["px_close"].transform(lambda x: np.log(x / x.shift(1)))
daily = df.groupby("date")["log_return"].mean().reset_index()
daily.to_parquet("gold/daily_avg_return.parquet")
```

## Schema Handling

Knowing what columns you expect and enforcing that contract at ingestion

> If a required column like `id` is missing, the pipeline should **alert and abort** — not silently produce nulls downstream

- define the expected schema explicitly
- validate against it on arrival
- fail loudly on required column absence

## schema evolution

handling upstream schema changes gracefully without crashing

| Scenario | Wrong approach | Right approach |
|----------|---------------|----------------|
| Source adds an unknown column | Crash on unknown column | Log it, ignore it, continue |
| Source renames an expected column | Silently produce nulls | Detect missing expected column, alert |

## data integrity rules

Invariants your data must satisfy, checked after **each transformation** — not just at the end.

| Rule | Example |
|------|---------|
| No nulls on key fields | `id` must never be null in the output dataset |
| Value range checks | `return` outside `[-1, 2]` for a single day is suspicious |
| Valid domain values | `date` must be a valid trading day |

## merge and transform safety

- **Never** select `df["some_column"]` without checking it exists first — use `.get()` or `"some_column" in df.columns`
- **Before merging** on a key column, verify both sides have the key and neither has unexpected duplicates
- **Always assert cardinality** after merges — a silent many-to-many explodes your row count

> If a merge is many-to-many when you expected one-to-one, your row count explodes silently

## output stability

Downstream consumers depend on a fixed output schema.

- The schema must **never change** without a versioned migration
- All expected columns must always exist — fill with `NULL` if empty, never drop them
- Example: if your gold layer feeds a risk calculator, it must always emit `["date", "id", "log_return", "financial_volume"]` even on days with missing data

> Dropping a column because it's empty that day will crash the downstream reader with a `KeyError`.

## idempotency and safety

A pipeline is idempotent if running it **twice** produces the same result as running it **once**.

- **Prefer upsert over overwrite** — re-running a date range updates existing rows, not duplicates them
- **Never delete source data** before writing the result
- If the write fails halfway, you must be able to rerun from scratch without data loss

### delete + insert vs upsert

**Delete + insert** removes all existing rows for a partition, then writes the new ones. Simple but risky: if the insert crashes, the deleted data is gone

```python
db.execute("DELETE FROM prices WHERE date = '2026-01-02'")
db.execute_many("INSERT INTO prices ...", new_rows)  # if this crashes, data is gone
```

**Upsert** updates a row if it exists, inserts it if it doesn't, so existing data is never removed first

```python
# INSERT ... ON CONFLICT (id, date) DO UPDATE SET px_close = EXCLUDED.px_close
```

| | Delete + Insert | Upsert |
|---|---|---|
| Atomicity | Two operations — gap between them | Single operation per row |
| Partial failure | Can lose data if insert crashes | Existing rows untouched on failure |
| Idempotency | Yes, if wrapped in a transaction | Yes, by design |
| Performance | Faster for full partition rewrites | Slower due to conflict checks |

Use **delete + insert inside a transaction** when rewriting a full partition and your DB supports atomic transactions. Use **upsert** when merging incremental updates where data gaps are not acceptable.

## observability

Every pipeline run must emit structured logs covering:

| What to log | Example message |
|-------------|-----------------|
| Missing required columns | `"Expected id_qtq in source, not found — aborting."` |
| Missing optional columns | `"Column px_open absent — filling with null."` |
| Unknown new columns | `"Unknown column received — logging and ignoring."` |
| Row counts at each step | `"Input: 12,450 → dedup: 12,448 → join: 12,300"` |
| Dropped records with reason | `"Dropped 148 rows: unknown tickers"` |

> Logs must be good enough to answer "why was this ticker missing on March 10th?" six months later

## core principles

- don't trust data: always validate at ingestion
    - sources lie, break formats and send duplicates
- don't lose data: land raw first, transform later
    - never overwrite source
- don't break on change: guard against schema drift so a new upstream column doesn't take down your pipeline
- always be able to explain the data later: logs and audit trails are not optional
