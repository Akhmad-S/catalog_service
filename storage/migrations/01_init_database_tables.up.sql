CREATE TABLE category(
    id	CHAR(36) PRIMARY KEY,
    category_title VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE product(
    id	CHAR(36) PRIMARY KEY,
    category_id CHAR(36),
    title VARCHAR(30) NOT NULL,
    descrip VARCHAR(250),
    price VARCHAR(10),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

ALTER TABLE product ADD CONSTRAINT fk_product_category FOREIGN KEY (category_id) REFERENCES category (id);
