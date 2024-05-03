CREATE TABLE orders (
    id varchar(36) NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz,
    user_id varchar(36) NOT NULL,
    payload text NOT NULL,
    is_success boolean NOT NULL,
    CONSTRAINT orders_id_pkey PRIMARY KEY (id)
);
