-- Users Table
CREATE TABLE users (
    user_id serial PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(16),
    role INT,
    password VARCHAR(80)
);

-- Tasks Table
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    status VARCHAR(50),
    manager_id INT REFERENCES users(user_id),
    assignee_id INT REFERENCES users(user_id)
);


