DROP TABLE IF EXISTS persons;

-- migrate -database postgres://postgres:learngo123@localhost:5432/postgres?sslmode=disable -path database/migrations up