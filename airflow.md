[python](./python.md)

# Airflow

> platform to programmatically author, schedule and monitor workflows

- airflow is an orchestrator, not a data processor
- workflows are defined as code
- DAGs (Directed Acyclic Graphs): the workflow
  - schedule: cron (`"0 2 * * *"` or presets `"@daily"`)
  - catchup
  - tags
- task: unit of work (python, bash, sql, spark, etc)
- operators: template for a predefined Task
  - `BashOperator`: executes a bash command
  - `PythonOperator`: calls an arbitrary Python function
  - `EmailOperator`: sends an email

example of minimal DAG

```python
# dags/example_etl.py
from pendulum import datetime
from airflow.decorators import dag, task

@dag(
    start_date=datetime(2024, 1, 1),   # first logical interval
    schedule="@daily",                 # each run covers previous day
    catchup=False,                     # avoid giant backfills while developing
    tags=["etl", "example"],
)
def example_etl():
    @task
    def extract():
        # Fetch rows from somewhere; return SMALL data only
        return [{"id": 1, "value": 3}, {"id": 2, "value": 5}]

    @task
    def transform(rows: list[dict]):
        return [r | {"value": r["value"] * 2} for r in rows]

    @task
    def load(rows: list[dict]):
        # Write to DB/data lake; keep this idempotent!
        print(f"Loaded {len(rows)} rows")

    load(transform(extract()))

example_etl()
```
