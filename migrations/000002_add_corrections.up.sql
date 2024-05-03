CREATE TABLE corrections (
    id varchar(36) NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz,
    payload text NOT NULL,
    is_error boolean NOT NULL,
    CONSTRAINT corrections_id_pkey PRIMARY KEY (id)
);
