package store

const (
	selectUserByLoginQuery = `SELECT id FROM users WHERE login = $1`
	insertOrderQuery       = `INSERT INTO orders (user_id, number, process, accrual, uploaded_at)
	VALUES ($1, $2, $3, $4, $5)`
	selectBalanceQuery = `SELECT balance FROM balance WHERE user_id = $1`
	updateBalanceQuery = `UPDATE balance SET balance = $1 WHERE user_id = $2`
	insertUserQuery    = `INSERT INTO users (login, password, token)
				VALUES ($1, $2, $3)
				RETURNING id;`

	insertBalanceQuery = `INSERT INTO users (user_id, balance)
				VALUES ($1, $2);`
	selectUserByLoiginPasswordQuery = `SELECT id FROM users WHERE login = $1 and password = $2`
	selectOrdersQuery               = `SELECT number, process, accrual, uploaded_at FROM orders WHERE user_id = $1`
	updateUserTokenQuery            = `UPDATE users SET token = $1 WHERE login = $2`
)
