import pandas as pd
import os
import logging
from config import DATASETS, LOCAL_TMP_DIR, BATCH_SIZE
from olist_plugin import S3Helper

def extract_batch(dataset_name, **kwargs):
    logger = logging.getLogger(f"extract_{dataset_name}")
    logger.info("Starting extraction for dataset: %s", dataset_name)

    done_flag = os.path.join(LOCAL_TMP_DIR, f"done_{dataset_name}.flag")
    if os.path.exists(done_flag):
        logger.info("Dataset '%s' already fully processed. Skipping extraction.", dataset_name)
        kwargs['ti'].xcom_push(key='batch_file', value='SKIPPED')
        return

    ds = next(d for d in DATASETS if d["name"] == dataset_name)
    remaining_file = os.path.join(LOCAL_TMP_DIR, f"remaining_{dataset_name}.csv")
    batch_file = os.path.join(LOCAL_TMP_DIR, f"batch_{dataset_name}.csv")

    if os.path.exists(remaining_file):
        logger.info("Found remaining file: %s", remaining_file)
        df = pd.read_csv(remaining_file)
    else:
        s3 = S3Helper()
        filename = os.path.basename(ds["s3_key"])
        local_path = os.path.join(LOCAL_TMP_DIR, filename)
        logger.info("Downloading from S3: %s -> %s", ds["s3_key"], local_path)
        s3.download_file(ds["s3_key"], local_path)
        df = pd.read_csv(local_path)
        df.to_csv(remaining_file, index=False)
        logger.info("Created initial remaining file: %s", remaining_file)

    if df.empty:
        logger.info("No data left for extraction in dataset: %s", dataset_name)
        # Write a "done" marker file
        with open(done_flag, 'w') as f:
            f.write("done")
        kwargs['ti'].xcom_push(key='batch_file', value='SKIPPED')
        return

    batch_df = df.head(BATCH_SIZE)
    batch_df.to_csv(batch_file, index=False)
    logger.info("Extracted batch of %d rows: %s", len(batch_df), batch_file)
    kwargs['ti'].xcom_push(key='batch_file', value=batch_file)
