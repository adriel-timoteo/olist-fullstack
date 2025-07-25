-- CREATE DATABASE
CREATE DATABASE olist_db;

-- Connect to the new database
\connect olist_db

-- USERS
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR NOT NULL UNIQUE,
    password_hash VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

-- MARKETING FUNNEL
CREATE TABLE marketing_funnel (
    mql_id UUID PRIMARY KEY,
    first_contact_date TIMESTAMP NOT NULL,
    landing_page_id VARCHAR(50),
    origin VARCHAR(50),
    lead_type VARCHAR(50),
    lead_behaviour_profile VARCHAR(50),
    declared_product_category VARCHAR(100),
    business_segment VARCHAR(50),
    won_date TIMESTAMP
);


-- CUSTOMERS
CREATE TABLE customers (
    customer_id UUID PRIMARY KEY,
    customer_unique_id UUID NOT NULL UNIQUE,
    customer_zip_code_prefix INT NOT NULL,
    customer_city VARCHAR(100) NOT NULL,
    customer_state VARCHAR(10) NOT NULL,
    mql_id UUID,
    FOREIGN KEY (mql_id) REFERENCES marketing_funnel(mql_id)
);

-- ORDERS
CREATE TABLE orders (
    order_id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    order_status VARCHAR(20) NOT NULL,
    order_purchase_timestamp TIMESTAMP NOT NULL,
    order_approved_at TIMESTAMP,
    order_delivered_carrier_date TIMESTAMP,
    order_delivered_customer_date TIMESTAMP,
    order_estimated_delivery_date TIMESTAMP NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id)
);

-- SELLERS
CREATE TABLE sellers (
    seller_id UUID PRIMARY KEY,
    seller_zip_code_prefix INT NOT NULL,
    seller_city VARCHAR(100) NOT NULL,
    seller_state VARCHAR(10) NOT NULL
);

-- PRODUCTS
CREATE TABLE products (
    product_id UUID PRIMARY KEY,
    product_category_name VARCHAR(100),
    product_weight_g INT,
    product_length_cm INT,
    product_height_cm INT,
    product_width_cm INT
);

-- ORDER ITEMS
CREATE TABLE order_items (
    order_id UUID NOT NULL,
    order_item_id INT NOT NULL,
    product_id UUID NOT NULL,
    seller_id UUID NOT NULL,
    shipping_limit_date TIMESTAMP NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    freight_value NUMERIC(10,2) NOT NULL,
    PRIMARY KEY (order_id, order_item_id),
    FOREIGN KEY (order_id) REFERENCES orders(order_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    FOREIGN KEY (seller_id) REFERENCES sellers(seller_id)
);

-- PAYMENTS
CREATE TABLE payments (
    order_id UUID NOT NULL,
    payment_sequential INT NOT NULL,
    payment_type VARCHAR(50) NOT NULL,
    payment_installments INT NOT NULL,
    payment_value NUMERIC(10,2) NOT NULL,
    PRIMARY KEY (order_id, payment_sequential),
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

-- REVIEWS
CREATE TABLE reviews (
    review_id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    review_score INT NOT NULL,
    review_comment_title TEXT,
    review_comment_message TEXT,
    review_creation_date TIMESTAMP NOT NULL,
    review_answer_timestamp TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

-- GEOLOCATION
CREATE TABLE geolocation (
    geolocation_id SERIAL PRIMARY KEY,
    zip_code_prefix INT NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(10) NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
);
