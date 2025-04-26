CREATE TABLE tasks (
	id SERIAL PRIMARY KEY,
	task TEXT NOT NULL,
	is_done BOOLEAN DEFAULT false,
	user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);
