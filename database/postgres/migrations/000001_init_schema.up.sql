-- Users Table
CREATE TABLE users (
    user_id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(16),
    role INT,
    password VARCHAR(80)
);

