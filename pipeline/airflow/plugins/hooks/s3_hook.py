import boto3
import os
from config import AWS_KEY, AWS_SECRET, S3_BUCKET, S3_REGION

class S3Helper:
    def __init__(self):
        self.client = boto3.client(
            "s3",
            region_name=S3_REGION,
            aws_access_key_id=AWS_KEY,
            aws_secret_access_key=AWS_SECRET,
        )
        self.bucket = S3_BUCKET

    def download_file(self, s3_key, local_path):
        os.makedirs(os.path.dirname(local_path), exist_ok=True)
        self.client.download_file(self.bucket, s3_key, local_path)
