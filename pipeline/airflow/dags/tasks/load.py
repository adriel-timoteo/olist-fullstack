import os
import pandas as pd
import logging
from config import DATASETS, LOCAL_TMP_DIR
from olist_plugin import PostgresHelper
from utils.transformations import clean_nulls

def load_batch(dataset_name, **kwargs):
    logger = logging.getLogger(f"load_{dataset_name}")
    logger.info("Starting load for dataset: %s", dataset_name)

    ds = next(d for d in DATASETS if d["name"] == dataset_name)
    transformed_file = kwargs['ti'].xcom_pull(key='transformed_file', task_ids=f"transform_{dataset_name}")
    remaining_file = os.path.join(LOCAL_TMP_DIR, f"remaining_{dataset_name}.csv")
    done_flag = os.path.join(LOCAL_TMP_DIR, f"done_{dataset_name}.flag")

    if transformed_file in [None, "SKIPPED"] or not os.path.exists(transformed_file):
        logger.info("No transformed file to load for dataset: %s", dataset_name)
        return

    df = pd.read_csv(transformed_file, dtype=str)
    if df.empty:
        os.remove(transformed_file)
        logger.info("No valid rows to load for dataset: %s", dataset_name)
        return

    pk_cols = ds.get("pk", [])
    if not pk_cols:
        logger.warning("No primary key defined for dataset: %s. Skipping deduplication.", dataset_name)

    pg = PostgresHelper()

    if pk_cols:
        # Build WHERE clause for composite PK
        pk_select = f"SELECT {', '.join(pk_cols)} FROM {ds['table']}"
        with pg.conn.cursor() as cur:
            cur.execute(pk_select)
            existing_keys = set(tuple(map(str, row)) for row in cur.fetchall())

        before_count = len(df)
        df["__pk_tuple__"] = df[pk_cols].astype(str).agg(tuple, axis=1)
        df = df[~df["__pk_tuple__"].isin(existing_keys)].drop(columns=["__pk_tuple__"])
        after_count = len(df)

        logger.info("Deduplicated %d rows using primary key(s): %s", before_count - after_count, pk_cols)

    if df.empty:
        os.remove(transformed_file)
        logger.info("All rows already exist in table '%s'. Nothing to insert.", ds["table"])
        return

    rows = df[ds["columns"]].map(clean_nulls).values.tolist()
    columns_str = ', '.join(ds["columns"])
    placeholders = ', '.join(['%s'] * len(ds["columns"]))
    query = f"INSERT INTO {ds['table']} ({columns_str}) VALUES ({placeholders})"

    try:
        pg.batch_insert(query, rows)
        logger.info("Inserted %d new rows into table '%s'.", len(rows), ds["table"])

        if os.path.exists(remaining_file):
            remaining_df = pd.read_csv(remaining_file, dtype=str)

            if pk_cols:
                # Drop rows with same PK from remaining_file
                remaining_df["__pk_tuple__"] = remaining_df[pk_cols].astype(str).agg(tuple, axis=1)
                df["__pk_tuple__"] = df[pk_cols].astype(str).agg(tuple, axis=1)
                remaining_df = remaining_df[~remaining_df["__pk_tuple__"].isin(df["__pk_tuple__"])]
                remaining_df = remaining_df.drop(columns=["__pk_tuple__"])

            if remaining_df.empty:
                os.remove(remaining_file)
                with open(done_flag, 'w') as f:
                    f.write("done")
                logger.info("All data loaded. Marked dataset '%s' as done.", dataset_name)
            else:
                remaining_df.to_csv(remaining_file, index=False)
                logger.info("Updated remaining file with %d rows left.", len(remaining_df))

        os.remove(transformed_file)

    except Exception as e:
        logger.error("Failed to insert rows into %s: %s", ds["table"], str(e))
        raise
