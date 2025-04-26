CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	email TEXT NOT NULL,
	password TEXT NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);
