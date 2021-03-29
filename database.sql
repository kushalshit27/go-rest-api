CREATE TABLE posts (
	id serial PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT DEFAULT '',
	created_on TIMESTAMP  DEFAULT NULL,
	status BOOLEAN NOT NULL DEFAULT 'true'
);


