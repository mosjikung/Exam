CREATE TABLE IF NOT EXISTS products (
    id           SERIAL      PRIMARY KEY,
    product_code VARCHAR(19) NOT NULL,

    CONSTRAINT uni_products_product_code UNIQUE (product_code),
    CONSTRAINT chk_product_code_length   CHECK (LENGTH(product_code) = 19),
    CONSTRAINT chk_product_code_format   CHECK (
        product_code ~ '^[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}$'
    )
);

CREATE INDEX IF NOT EXISTS idx_products_code ON products (product_code);
