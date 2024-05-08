CREATE TABLE cars (
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(20) NOT NULL,
    mark VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    year INT
);
