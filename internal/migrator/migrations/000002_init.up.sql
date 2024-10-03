CREATE TABLE
    IF NOT EXISTS withdrawals (
        id SERIAL PRIMARY KEY,
        user_id INT,
        order TEXT,
        sum INT,
        processed_at TEXT,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    )
