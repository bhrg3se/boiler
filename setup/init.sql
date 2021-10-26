
CREATE TABLE users (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at timestamp,
    updated_at timestamp
);

