CREATE TABLE persons(
     id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)


-- migrate -database postgres://postgres:learngo123@localhost:5432/postgres?sslmode=disable -path database/migrations up