
-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Insert fake users into the users table
INSERT INTO users (username, password_hash, created_at, updated_at) VALUES
('john_doe', '$2a$12$5O8hvM8rD9bc7K9zVZWuXeJovjJoHXOH.qkBLo0PsnD68/LOvIdU2', NOW(), NOW()), -- hashed password: "password123"
('jane_doe', '$2a$12$5O8hvM8rD9bc7K9zVZWuXeJovjJoHXOH.qkBLo0PsnD68/LOvIdU2', NOW(), NOW()), -- hashed password: "password123"
('admin', '$2a$12$5O8hvM8rD9bc7K9zVZWuXeJovjJoHXOH.qkBLo0PsnD68/LOvIdU2', NOW(), NOW());     -- hashed password: "password123"
