# Data Engineering

## The Medallion Architecture (bronze => silver => gold)

| Layer | Description | Example |
|-------|-------------|---------|
| **Bronze** | Raw data, no transformations | Bloomberg parquet with nulls, duplicates and misspelled columns → `raw` |
| **Silver** | Cleaned and standardized | Nulls handled, duplicates removed, columns renamed → `clean` |
| **Gold** | Business-ready aggregations | Daily returns per ticker, risk summary per fund, portfolio NAV series |

Each layer exists so you can reprocess downstream without re-fetching source data.

## Schema Handling

Knowing what columns you expect and enforcing that contract at ingestion.

> If a required column like `id` is missing, the pipeline should **alert and abort** — not silently produce nulls downstream.

- Define the expected schema explicitly
- Validate against it on arrival
- Fail loudly on required column absence

## Schema Evolution

Handling upstream schema changes gracefully without crashing.

| Scenario | Wrong approach | Right approach |
|----------|---------------|----------------|
| Bloomberg adds `PX_OPEN_ADJ` | Crash on unknown column | Log it, ignore it, continue |
| Bloomberg renames `PX_LAST` → `PX_CLOSE` | Silently produce nulls | Detect missing expected column, alert |

## Data Integrity Rules

Invariants your data must satisfy, checked after **each transformation** — not just at the end.

| Rule | Example |
|------|---------|
| No nulls on key fields | `id_qtq` must never be null in `market-data-v2` |
| Value range checks | `return` outside `[-1, 2]` for a single day is suspicious |
| Valid domain values | `date` must be a valid trading day |

## Merge & Transform Safety

- **Never** select `df["some_column"]` without checking it exists first — use `.get()` or `"some_column" in df.columns`
- **Before merging** on `id_qtq`, verify both sides have the key and neither has unexpected duplicates
- **Always assert cardinality** after merges — a silent many-to-many explodes your row count

> If a merge is many-to-many when you expected one-to-one, your row count explodes silently.

## Output Stability

Downstream consumers depend on a fixed output schema.

- The schema must **never change** without a versioned migration
- All expected columns must always exist — fill with `NULL` if empty, never drop them
- Example: if your gold layer feeds a risk calculator, it must always emit `["date", "id_qtq", "log_return", "financial_volume"]` even on days with missing data

> Dropping a column because it's empty that day will crash the downstream reader with a `KeyError`.

## Idempotency & Safety

A pipeline is idempotent if running it **twice** produces the same result as running it **once**.

- **Prefer upsert over overwrite** — re-running a date range updates existing rows, not duplicates them
- **Never delete source data** before writing the result
- If the write fails halfway, you must be able to rerun from scratch without data loss

## Observability

Every pipeline run must emit structured logs covering:

| What to log | Example message |
|-------------|-----------------|
| Missing required columns | `"Expected id_qtq in source, not found — aborting."` |
| Missing optional columns | `"Column px_open absent — filling with null."` |
| Unknown new columns | `"Unknown column PX_OPEN_ADJ received — logging and ignoring."` |
| Row counts at each step | `"Input: 12,450 → dedup: 12,448 → join: 12,300"` |
| Dropped records with reason | `"Dropped 148 rows: unknown tickers"` |

> Logs must be good enough to answer "why was this ticker missing on March 10th?" six months later.

## Core Principles

- Don't trust data: always validate at ingestion
    - sources lie, break formats and send duplicates
- Don't lose data: land raw first, transform later
    - never overwrite source
- Don't break on change: guard against schema drift so a new upstream column doesn't take down your pipeline
- Always be able to explain the data later: logs and audit trails are not optional
