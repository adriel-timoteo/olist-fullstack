import psycopg2
from psycopg2.extras import execute_batch
import os
from config import DB_HOST, DB_NAME, DB_USER, DB_PASSWORD

class PostgresHelper:
    def __init__(self):
        self.conn = psycopg2.connect(
            host=DB_HOST,
            dbname=DB_NAME,
            user=DB_USER,
            password=DB_PASSWORD,
        )
        self.conn.autocommit = True

    def batch_insert(self, query: str, rows: list, page_size: int = 1000):
        """
        Execute a batch insert using psycopg2's execute_batch.

        Args:
            query (str): SQL insert query with placeholders (%s).
            rows (list): List of tuples containing row data.
            page_size (int): Number of records per batch.
        """
        if not rows:
            print("No rows to insert.")
            return
        with self.conn.cursor() as cursor:
            execute_batch(cursor, query, rows, page_size=page_size)
