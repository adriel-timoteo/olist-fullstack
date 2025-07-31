-- -- CREATE DATABASE
-- CREATE DATABASE olist_db;

-- -- Connect to the new database
-- \connect olist_db

-- CATEGORIES
CREATE TABLE category_translations (
    product_category_id SERIAL PRIMARY KEY,
    product_category_name TEXT NOT NULL UNIQUE,
    product_category_name_english TEXT
);

-- GEOLOCATION
CREATE TABLE geolocations (
    geolocation_zip_code_prefix TEXT NOT NULL UNIQUE,
    geolocation_city TEXT NOT NULL,
    geolocation_state TEXT NOT NULL,
    geolocation_lat FLOAT NOT NULL,
    geolocation_lng FLOAT NOT NULL,
    geolocation_id SERIAL PRIMARY KEY
);

-- USERS
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- MARKETING FUNNEL
CREATE TABLE marketing_funnel (
    mql_id UUID PRIMARY KEY,
    first_contact_date TIMESTAMP DEFAULT NOW(),
    landing_page_id TEXT,
    origin TEXT,
    lead_type TEXT,
    lead_behaviour_profile TEXT,
    declared_product_category TEXT,
    business_segment TEXT,
    won_date TIMESTAMP
);


-- CUSTOMERS
CREATE TABLE customers (
    customer_id UUID PRIMARY KEY,
    customer_unique_id UUID NOT NULL,
    customer_zip_code_prefix TEXT NOT NULL,
    customer_city TEXT NOT NULL,
    customer_state TEXT NOT NULL
);

-- ORDERS
CREATE TABLE orders (
    order_id UUID PRIMARY KEY,
    customer_id UUID NOT NULL,
    order_status TEXT NOT NULL,
    order_purchase_timestamp TIMESTAMP DEFAULT NOW(),
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
    seller_city TEXT NOT NULL,
    seller_state TEXT NOT NULL
);

-- PRODUCTS
CREATE TABLE products (
    product_id UUID PRIMARY KEY,
    product_category_name TEXT,
    product_weight_g FLOAT,
    product_length_cm FLOAT,
    product_height_cm FLOAT,
    product_width_cm FLOAT,
    FOREIGN KEY (product_category_name) REFERENCES category_translations(product_category_name)
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
    payment_type TEXT NOT NULL,
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
    review_creation_date TIMESTAMP DEFAULT NOW(),
    review_answer_timestamp TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

-- SEED TRANSLATIONS
\copy category_translations(product_category_name, product_category_name_english) FROM 'D:/Programming/olist-fullstack/db/dataset/ecommerce/product_category_name_translation.csv' DELIMITER ',' CSV HEADER;

INSERT INTO category_translations (product_category_name, product_category_name_english)
VALUES
('pc_gamer', 'gaming_pc'),
('portateis_cozinha_e_preparadores_de_alimentos', 'kitchen_appliances_and_food_processors');