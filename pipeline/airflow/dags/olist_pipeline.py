from airflow import DAG
from airflow.operators.python import PythonOperator
from datetime import datetime, timedelta
from tasks.extract import extract_batch, extract_batch_by_fk
from tasks.transform import filter_valid_foreign_keys, transform_order_dates
from tasks.load import load_batch
from config import DATASETS

default_args = {
    'owner': 'airflow',
    'start_date': datetime(2023, 1, 1),
    'retries': 1,
    'retry_delay': timedelta(minutes=1),
}

dag = DAG(
    'olist_pipeline',
    default_args=default_args,
    description='Batched ETL with FK constraint handling and incremental cleanup',
    # schedule='*/5 * * * *',
    catchup=False
)

task_refs = {}

for ds in DATASETS:
    name = ds["name"]
    fk_map = ds.get("fk", {})
    upstream = ds.get("upstream", None)

    if fk_map and upstream:
        # FK-based extraction
        extract = PythonOperator(
            task_id=f"extract_{name}_by_fk",
            python_callable=extract_batch_by_fk,
            op_kwargs={
                "dataset_name": name,
                "fk_column": list(fk_map.keys())[0],  # Currently only supporting 1 FK field per dataset
                "upstream_dataset_name": upstream,
            },
            dag=dag,
        )
    else:
        # Standard extraction
        extract = PythonOperator(
            task_id=f"extract_{name}",
            python_callable=extract_batch,
            op_kwargs={"dataset_name": name},
            dag=dag,
        )

    # Common transform
    transform = PythonOperator(
        task_id=f"transform_{name}",
        python_callable=filter_valid_foreign_keys,
        op_kwargs={"dataset_name": name},
        dag=dag,
    )

    # Load
    load = PythonOperator(
        task_id=f"load_{name}",
        python_callable=load_batch,
        op_kwargs={"dataset_name": name},
        dag=dag,
    )

    # Orders gets special treatment (extra transform)
    if name == "orders":
        date_transform = PythonOperator(
            task_id="transform_order_dates",
            python_callable=transform_order_dates,
            op_kwargs={"dataset_name": name},
            dag=dag,
        )
        extract >> transform >> date_transform >> load
    else:
        extract >> transform >> load

    task_refs[name] = {
        "extract": extract,
        "load": load,
    }

# Set FK-based extract dependencies to run *after* upstream LOAD
for ds in DATASETS:
    if "fk" in ds:
        for fk_field, parent_table in ds["fk"].items():
            upstream_load = task_refs.get(parent_table, {}).get("load")
            current_extract = task_refs[ds["name"]]["extract"]
            if upstream_load and current_extract:
                upstream_load >> current_extract