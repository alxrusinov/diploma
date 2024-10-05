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
        accrual FLOAT (2),
        uploaded_at TEXT,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS balance (
        id SERIAL PRIMARY KEY,
        user_id INT,
        balance FLOAT (2),
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );
