[python](./python.md)

# Airflow

> platform to programmatically author, schedule and monitor workflows

- airflow is an orchestrator, not a data processor
- workflows are defined as code
- DAGs (Directed Acyclic Graphs): the workflow
  - schedule: cron (`"0 2 * * *"` or presets `"@daily"`)
    - set `schedule=None` for manual/externally triggered DAGs; if you omit it, Airflow defaults to a schedule (`@daily`) and will create runs automatically
  - catchup
  - tags
- task: unit of work (python, bash, sql, spark, etc)
    - nodes: tasks
        - task: unit of work (python, bash, sql, spark, etc)
    - edges: dependencies
    - schedule: cron (`"0 2 * * *"` or presets `"@daily"`)
    - catchup
    - tags
    - properties
        - each DAG run covers a fixed time interval (e.g. 1 day)
        - each task instance processes data for that interval
        - tasks should be idempotent (re-running a task for the same interval should not change the result)
        - tasks can pass small amounts of data (XComs) but not large datasets
        - tasks can be retried on failure
        - tasks can be run in parallel
        - tasks can be distributed across multiple workers
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

passing params into tasks

```python
import pendulum, datetime as dt, pytz
from airflow import DAG
from airflow.models.param import Param
from airflow.operators.python import PythonOperator

with DAG(
    dag_id="bbg_dag",
    schedule=None,  # manual/external trigger only; omit => @daily kicks in
    start_date=pendulum.now("America/Sao_Paulo").subtract(days=2),
    catchup=False,
    params={
        "query_date": Param(
            default=str(dt.datetime.now(pytz.timezone("America/Sao_Paulo")).date()),
            type=["string"],
            format="date",
        ),
        "overwrite": Param(default=False, type="boolean"),
        "ref_params_names": Param(
            default=[
                "param_bond_pricing_request",
                "param_crypto_pricing_request",
                "param_index_pricing_request",
            ],
            type="array",
        ),
    },
    render_template_as_native_obj=True,
) as dag:
    load_prices = PythonOperator(
        task_id="load_prices",
        python_callable=raw_bloomberg_equity_pricing,
        op_kwargs={
            "query_date": "{{ params.query_date }}",
            "overwrite": "{{ params.overwrite }}",
            "ref_params_names": "{{ params.ref_params_names }}",
        },
    )
```

- define DAG-level `params` (optionally with `Param` for types/defaults) and reference them inside tasks with Jinja (`{{ params.<key> }}`) via `op_kwargs` or templates.
