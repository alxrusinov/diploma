CREATE TABLE
    IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        login TEXT,
        password TEXT,
        token TEXT
    );

CREATE TABLE
    IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        user_id INT,
        number TEXT,
        process TEXT,
        accrual INT,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS balance (
        id SERIAL PRIMARY KEY,
        user_id INT,
        balance FLOAT,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );
