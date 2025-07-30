import os

AWS_KEY = os.getenv("AWS_ACCESS_KEY_ID")
AWS_SECRET = os.getenv("AWS_SECRET_ACCESS_KEY")
S3_BUCKET = "olist-raw-data-bucket"
S3_REGION = "ap-southeast-1"

DB_HOST = "olist-db.crw6o2g8mpd6.ap-southeast-1.rds.amazonaws.com"
DB_NAME = "olist_db"
DB_USER = "postgres"
DB_PASSWORD = "KcR9YQPwksBDXhZkVT9i"

BATCH_SIZE = 10000

LOCAL_TMP_DIR = "/tmp/olist"

DATASETS = [
    {
        "name": "customers",
        "s3_key": "ecommerce/olist_customers_dataset.csv",
        "table": "customers",
        "pk": ["customer_id"],
        "columns": [
            "customer_id", "customer_unique_id", "customer_zip_code_prefix",
            "customer_city", "customer_state"
        ]
    },
    {
        "name": "products",
        "s3_key": "ecommerce/olist_products_dataset.csv",
        "table": "products",
        "pk": ["product_id"],
        "columns": [
            "product_id", "product_category_name", "product_weight_g",
            "product_length_cm", "product_height_cm", "product_width_cm",
        ]
    },
    {
        "name": "sellers",
        "s3_key": "ecommerce/olist_sellers_dataset.csv",
        "table": "sellers",
        "pk": ["seller_id"],
        "columns": [
            "seller_id", "seller_zip_code_prefix", "seller_city", "seller_state",
        ]
    },
    {
        "name": "orders",
        "s3_key": "ecommerce/olist_orders_dataset.csv",
        "table": "orders",
        "pk": ["order_id"],
        "fk": {"customer_id": "customers"},
        "columns": [
            "order_id", "customer_id", "order_status", "order_purchase_timestamp",
            "order_approved_at", "order_delivered_carrier_date",
            "order_delivered_customer_date", "order_estimated_delivery_date"
        ]
    },
    {
        "name": "order_items",
        "s3_key": "ecommerce/olist_order_items_dataset.csv",
        "table": "order_items",
        "pk": ["order_id", "product_id"],
        "fk": {"order_id": "orders", "product_id": "products", "seller_id": "sellers"},
        "columns": [
            "order_id", "order_item_id", "product_id",
            "shipping_limit_date", "price", "freight_value",
        ]
    },
    {
        "name": "order_payments",
        "s3_key": "ecommerce/olist_order_payments_dataset.csv",
        "table": "order_payments",
        "pk": ["order_id", "payment_sequential"],
        "fk": {"order_id": "orders"},
        "columns": [
            "order_id", "payment_sequential", "payment_type",
            "payment_installments", "payment_value",
        ]
    },
    {
        "name": "order_reviews",
        "s3_key": "ecommerce/olist_order_reviews_dataset.csv",
        "table": "order_reviews",
        "pk": ["review_id"],
        "fk": {"order_id": "orders"},
        "columns": [
            "order_id", "review_id", "review_score",
            "review_comment_title", "review_comment_message",
            "review_creation_date", "review_answer_timestamp",
        ]
    },
]