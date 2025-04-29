CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    surname VARCHAR(30) NOT NULL,
    patronymic VARCHAR(30) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(6) NOT NULL,
    country VARCHAR(50) NOT NULL
);