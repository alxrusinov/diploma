CREATE TABLE
    IF NOT EXISTS withdrawls (
        id SERIAL PRIMARY KEY,
        user_id INT,
        order_number TEXT,
        sum INT,
        processed_at TEXT,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );
