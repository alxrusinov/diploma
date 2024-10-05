package store

const (
	selectUserByLoginQuery = `SELECT id FROM users WHERE login = $1;`
	insertOrderQuery       = `INSERT INTO orders (user_id, number, process, accrual, uploaded_at)
	VALUES ($1, $2, $3, $4, $5);`
	selectBalanceQuery = `SELECT balance FROM balance WHERE user_id = $1;`
	updateBalanceQuery = `UPDATE balance SET balance = $1 WHERE user_id = $2`
	insertUserQuery    = `INSERT INTO users (login, password, token)
				VALUES ($1, $2, $3)
				RETURNING id;`

	insertBalanceQuery = `INSERT INTO balance (user_id, balance)
				VALUES ($1, $2);`
	selectUserByLoiginPasswordQuery = `SELECT id FROM users WHERE login = $1 and password = $2;`
	selectOrdersQuery               = `SELECT number, process, accrual, uploaded_at FROM orders WHERE user_id = $1 ORDER BY TO_TIMESTAMP(uploaded_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') DESC;`
	updateUserTokenQuery            = `UPDATE users SET token = $1 WHERE id = $2;`
	checkOrderQuery                 = `SELECT user_id FROM orders WHERE number = $1;`
	getOrderQuery                   = `SELECT number, process, accrual, uploaded_at FROM orders WHERE number = $1 and user_id = $2;`
	setWithdrawnQuery               = `INSERT INTO withdrawls (user_id, order_number, sum, processed_at)
	VALUES ($1, $2, $3, $4);`
	getWithdrawlsQuery = `SELECT order_number, sum, processed_at FROM withdrawls WHERE user_id = $1 ORDER BY TO_TIMESTAMP(processed_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') DESC`
)
