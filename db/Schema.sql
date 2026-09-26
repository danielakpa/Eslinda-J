-- Cakes available in the shop
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    image TEXT,
    category TEXT,
    has_variants BOOLEAN DEFAULT FALSE
);

-- Size/weight options for a cake, if it has any (e.g. 0.5kg, 1kg)
CREATE TABLE product_variants (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id),
    label TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

-- Customer orders
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_name TEXT NOT NULL,
    phone TEXT NOT NULL,
    address TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    total_amount NUMERIC(10, 2) NOT NULL,
    status TEXT DEFAULT 'pending',
    paystack_ref TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Individual cakes inside one order
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id),
    product_id INTEGER REFERENCES products(id),
    variant_id INTEGER REFERENCES product_variants(id),
    quantity INTEGER NOT NULL
);

-- People who filled the "Learn/Apply" form
CREATE TABLE class_interest (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    contact TEXT NOT NULL,
    notes TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);