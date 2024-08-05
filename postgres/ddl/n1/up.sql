SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public' AND table_type = 'BASE TABLE';

SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'public';

-- -----------------------------------------------------------------------

CREATE TABLE products (
    product_id SERIAL UNIQUE, -- not primary key!
    product_name VARCHAR(100)
);

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    product_id INT,
    quantity INT,
    FOREIGN KEY (product_id) 
        REFERENCES products(product_id)
);

-- -----------------------------------------------------------------------