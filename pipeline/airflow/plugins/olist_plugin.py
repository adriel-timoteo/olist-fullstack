from airflow.plugins_manager import AirflowPlugin

from hooks.s3_hook import S3Helper
from hooks.postgres_hook import PostgresHelper

class OlistPlugin(AirflowPlugin):
    name = "olist_plugin"
    hooks = [S3Helper, PostgresHelper]
