create table IF NOT EXISTS todos (
	id serial PRIMARY KEY,
	text VARCHAR(100) not null,
	created_at TIMESTAMP not NULL,
	updated_at TIMESTAMP
);