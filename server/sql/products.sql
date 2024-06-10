CREATE TABLE IF NOT EXISTS products (
                                        product_id SERIAL PRIMARY KEY,
                                        product_name VARCHAR(255) NOT NULL,
                                        product_description TEXT,
                                        product_sku VARCHAR(50) NOT NULL,
                                        product_cost DECIMAL(10,2),
                                        added_date TIMESTAMP DEFAULT NOW(),
                                        updated_date TIMESTAMP DEFAULT NOW()
);
