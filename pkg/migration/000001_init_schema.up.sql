CREATE TABLE datafiles (
    id SERIAL PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    data VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	
);