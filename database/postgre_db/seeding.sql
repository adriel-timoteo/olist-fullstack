-- SEED DATA FOR TESTING

INSERT INTO marketing_funnel (mql_id, first_contact_date, landing_page_id, origin, lead_type, lead_behaviour_profile, declared_product_category, business_segment, won_date)
VALUES
  (gen_random_uuid(), '2025-01-05 10:00:00', 'LP001', 'Google Ads', 'B2C', 'High', 'Electronics', 'Retail', '2025-02-10 15:00:00'),
  (gen_random_uuid(), '2025-01-15 11:30:00', 'LP002', 'Facebook', 'B2C', 'Medium', 'Fashion', 'Retail', NULL),
  (gen_random_uuid(), '2025-02-10 09:00:00', 'LP003', 'Email', 'B2B', 'High', 'Home', 'Wholesale', '2025-02-20 14:00:00'),
  (gen_random_uuid(), '2025-02-18 16:00:00', 'LP004', 'Instagram', 'B2C', 'Low', 'Sports', 'Retail', NULL),
  (gen_random_uuid(), '2025-03-01 08:30:00', 'LP005', 'Organic', 'B2B', 'Medium', 'Books', 'Retail', NULL),
  (gen_random_uuid(), '2025-03-10 17:30:00', 'LP006', 'Google Ads', 'B2C', 'High', 'Toys', 'Retail', '2025-03-20 12:00:00'),
  (gen_random_uuid(), '2025-04-05 14:00:00', 'LP007', 'Facebook', 'B2C', 'Low', 'Electronics', 'Retail', NULL),
  (gen_random_uuid(), '2025-04-15 13:00:00', 'LP008', 'Email', 'B2B', 'High', 'Garden', 'Wholesale', '2025-04-25 10:00:00'),
  (gen_random_uuid(), '2025-05-01 09:15:00', 'LP009', 'Instagram', 'B2C', 'Medium', 'Health', 'Retail', NULL),
  (gen_random_uuid(), '2025-05-20 15:00:00', 'LP010', 'Organic', 'B2C', 'High', 'Fashion', 'Retail', '2025-05-28 09:00:00');

INSERT INTO customers (customer_id, customer_unique_id, customer_zip_code_prefix, customer_city, customer_state, mql_id)
SELECT gen_random_uuid(), gen_random_uuid(), (10000 + i), 'City'||i, 'ST', mql_id
FROM generate_series(1, 10) i
JOIN (SELECT mql_id FROM marketing_funnel LIMIT 10) mf ON true;

INSERT INTO sellers (seller_id, seller_zip_code_prefix, seller_city, seller_state)
SELECT gen_random_uuid(), (20000 + i), 'SellerCity'||i, 'SS'
FROM generate_series(1, 10) i;

INSERT INTO products (product_id, product_category_name, product_weight_g, product_length_cm, product_height_cm, product_width_cm)
SELECT gen_random_uuid(), 'Category'||i, (100 + i*10), (10 + i), (5 + i), (3 + i)
FROM generate_series(1, 10) i;

INSERT INTO orders (order_id, customer_id, order_status, order_purchase_timestamp, order_approved_at, order_delivered_carrier_date, order_delivered_customer_date, order_estimated_delivery_date)
SELECT gen_random_uuid(), c.customer_id, 'delivered',
  NOW() - (i || ' days')::interval,
  NOW() - (i || ' days')::interval + INTERVAL '2 hours',
  NOW() - (i || ' days')::interval + INTERVAL '1 day',
  NOW() - (i || ' days')::interval + INTERVAL '3 days',
  NOW() + INTERVAL '7 days'
FROM generate_series(1, 10) i
JOIN customers c ON c.customer_id = (SELECT customer_id FROM customers OFFSET i-1 LIMIT 1);

INSERT INTO order_items (order_id, order_item_id, product_id, seller_id, shipping_limit_date, price, freight_value)
SELECT o.order_id, 1, p.product_id, s.seller_id,
  NOW() + INTERVAL '2 days',
  (50 + i * 5), (5 + i)
FROM generate_series(1, 10) i
JOIN orders o ON o.order_id = (SELECT order_id FROM orders OFFSET i-1 LIMIT 1)
JOIN products p ON p.product_id = (SELECT product_id FROM products OFFSET i-1 LIMIT 1)
JOIN sellers s ON s.seller_id = (SELECT seller_id FROM sellers OFFSET i-1 LIMIT 1);

INSERT INTO payments (order_id, payment_sequential, payment_type, payment_installments, payment_value)
SELECT o.order_id, 1, 'credit_card', (i % 3 + 1), (100 + i*10)
FROM generate_series(1, 10) i
JOIN orders o ON o.order_id = (SELECT order_id FROM orders OFFSET i-1 LIMIT 1);

INSERT INTO reviews (review_id, order_id, review_score, review_comment_title, review_comment_message, review_creation_date, review_answer_timestamp)
SELECT gen_random_uuid(), o.order_id, (i % 5 + 1), 'Title'||i, 'Message'||i,
  NOW() - (i || ' days')::interval, NOW() - (i || ' days')::interval + INTERVAL '1 hour'
FROM generate_series(1, 10) i
JOIN orders o ON o.order_id = (SELECT order_id FROM orders OFFSET i-1 LIMIT 1);

INSERT INTO geolocation (zip_code_prefix, city, state, latitude, longitude)
SELECT (30000 + i), 'GeoCity'||i, 'GS',
  ( -6.0 + i * 0.01), (106.0 + i * 0.01)
FROM generate_series(1, 10) i;
