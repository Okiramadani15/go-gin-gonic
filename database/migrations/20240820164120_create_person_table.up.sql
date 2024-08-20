CREATE TABLE person(
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(255) NOT NULL,
    address   VARCHAR(255) NOT NULL,
    email     VARCHAR(255) NOT NULL,
    birthdate TIMESTAMP NOT NULL
);

-- CREATE TABLE:
-- migrate create -ext sql -dir database/migrations create_person_table


-- RUNNING MIGRATION: 
-- migrate -database postgres://postgres:learngo123@localhost:5432/postgres?sslmode=disable -path database/migrations up