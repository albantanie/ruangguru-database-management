-- TODO: answer here
CREATE TABLE persons (
    id INTEGER PRIMARY KEY,
    NIK VARCHAR(255) NOT NULL UNIQUE,
    fullname VARCHAR(255) NOT NULL,
    gender VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL,
    is_married BOOLEAN,
    height FLOAT,
    weight FLOAT,
    address TEXT
);