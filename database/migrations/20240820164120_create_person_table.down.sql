DROP TABLE IF EXISTS person;

-- migrate -database postgres://postgres:learngo123@localhost:5432/postgres?sslmode=disable -path database/migrations up