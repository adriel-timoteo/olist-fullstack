import os
import pandas as pd
import logging
from config import DATASETS, LOCAL_TMP_DIR
from utils.transformations import normalize_uuid
from olist_plugin import PostgresHelper
from datetime import datetime

def filter_valid_foreign_keys(dataset_name, **kwargs):
    logger = logging.getLogger(f"transform_{dataset_name}")
    logger.info("Starting FK validation for dataset: %s", dataset_name)

    ds = next(d for d in DATASETS if d["name"] == dataset_name)
    batch_file = kwargs['ti'].xcom_pull(key='batch_file', task_ids=f"extract_{dataset_name}")

    if batch_file in [None, "SKIPPED"] or not os.path.exists(batch_file):
        logger.info("Skipping FK validation for dataset '%s' due to no batch file.", dataset_name)
        kwargs['ti'].xcom_push(key='transformed_file', value='SKIPPED')
        return

    df = pd.read_csv(batch_file, dtype=str)
    logger.info("Loaded %d rows for transformation.", len(df))

    if "fk" in ds:
        for fk_col, parent_table in ds["fk"].items():
            query = f"SELECT {fk_col} FROM {parent_table}"
            pg = PostgresHelper()
            with pg.conn.cursor() as cur:
                cur.execute(query)
                valid_ids = {normalize_uuid(row[0]) for row in cur.fetchall() if normalize_uuid(row[0])}

            logger.info("Fetched %d valid UUIDs from parent table '%s'.", len(valid_ids), parent_table)
            df[fk_col] = df[fk_col].apply(normalize_uuid)
            before_count = len(df)
            df = df[df[fk_col].isin(valid_ids)]
            logger.info("FK Validation complete for '%s': %d rows removed, %d remain.",
                        fk_col, before_count - len(df), len(df))
    else:
        logger.info("No FK constraints for dataset '%s'.", dataset_name)

    transformed_file = os.path.join(LOCAL_TMP_DIR, f"transformed_{dataset_name}.csv")
    df.to_csv(transformed_file, index=False)
    kwargs['ti'].xcom_push(key='transformed_file', value=transformed_file)

def transform_order_dates(dataset_name, **kwargs):
    logger = logging.getLogger(f"transform_dates_{dataset_name}")
    logger.info("Starting custom date transformation for dataset: %s", dataset_name)

    if dataset_name != "orders":
        raise ValueError("transform_order_dates is only applicable to the 'orders' dataset.")

    transformed_file = kwargs['ti'].xcom_pull(key='transformed_file', task_ids=f"transform_{dataset_name}")
    if not transformed_file or not os.path.exists(transformed_file):
        logger.warning("No transformed file found for dataset: %s", dataset_name)
        kwargs['ti'].xcom_push(key='date_transformed_file', value=None)
        return

    df = pd.read_csv(transformed_file)
    now = datetime.now()

    # Parse relevant columns as datetime
    datetime_cols = [
        "order_purchase_timestamp",
        "order_approved_at",
        "order_delivered_carrier_date",
        "order_delivered_customer_date",
        "order_estimated_delivery_date"
    ]
    for col in datetime_cols:
        df[col] = pd.to_datetime(df[col], errors="coerce")

    def transform_row(row):
        # Determine last non-null update
        update_cols = [
            "order_delivered_customer_date",
            "order_delivered_carrier_date",
            "order_approved_at",
            "order_purchase_timestamp"
        ]

        # Find the latest non-null update
        last_valid_idx = next(
            (i for i, col in enumerate(update_cols) if pd.notnull(row[col])),
            None
        )
        if last_valid_idx is None:
            return row  # no update timestamps at all, skip
        
        original_approved = row["order_approved_at"]
        original_estimated = row["order_estimated_delivery_date"]

        last_col = update_cols[last_valid_idx]
        last_original = row[last_col]
        row[last_col] = now

        # Adjust previous updates
        for col in update_cols[last_valid_idx+1:]:
            original = row[col]
            if pd.notnull(original):
                delta = last_original - original
                row[col] = now - delta

        # Nullify future updates (after last update)
        for col in update_cols[:last_valid_idx]:
            if pd.notnull(row[col]):
                row[col] = pd.NaT

        # Recalculate estimated delivery
        approved = row["order_approved_at"]
        
        if pd.notnull(original_approved) and pd.notnull(original_estimated):
            delta = original_estimated - original_approved
            row["order_estimated_delivery_date"] = approved + delta

        return row

    df = df.apply(transform_row, axis=1)

    df.to_csv(transformed_file, index=False)
    logger.info("Saved date-transformed data to %s with %d rows", transformed_file, len(df))
    kwargs['ti'].xcom_push(key='transformed_file', value=transformed_file)
