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
        accrual FLOAT (5),
        uploaded_at TEXT,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

CREATE TABLE
    IF NOT EXISTS balance (
        id SERIAL PRIMARY KEY,
        user_id INT,
        balance FLOAT (5),
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );
