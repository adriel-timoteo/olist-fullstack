from airflow import DAG
from airflow.operators.python import PythonOperator
from datetime import datetime, timedelta
from tasks.extract import extract_batch
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

task_map = {}

for ds in DATASETS:
    extract = PythonOperator(
        task_id=f"extract_{ds['name']}",
        python_callable=extract_batch,
        op_kwargs={'dataset_name': ds['name']},
        dag=dag,
    )

    transform = PythonOperator(
        task_id=f"transform_{ds['name']}",
        python_callable=filter_valid_foreign_keys,
        op_kwargs={'dataset_name': ds['name']},
        dag=dag,
    )

    if ds["name"] == "orders":
        date_transform = PythonOperator(
            task_id="transform_order_dates",
            python_callable=transform_order_dates,
            op_kwargs={'dataset_name': ds['name']},
            dag=dag,
        )

        load = PythonOperator(
            task_id=f"load_{ds['name']}",
            python_callable=load_batch,
            op_kwargs={'dataset_name': ds['name']},
            dag=dag,
        )

        extract >> transform >> date_transform >> load
    else:
        load = PythonOperator(
            task_id=f"load_{ds['name']}",
            python_callable=load_batch,
            op_kwargs={'dataset_name': ds['name']},
            dag=dag,
        )

        extract >> transform >> load

    task_map[ds['name']] = load

# Add FK dependencies
for ds in DATASETS:
    if "fk" in ds:
        for parent_table in ds["fk"].values():
            if parent_table in task_map:
                task_map[parent_table] >> task_map[ds['name']]
