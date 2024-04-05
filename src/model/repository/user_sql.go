package repository

var (
	sqlFindUserByID = `SELECT id, name, email 
					   FROM users 
					   WHERE id = $1 `

	sqlCreateUser = `INSERT INTO users (email, password, name, created_at) 
						values ($1,$2,$3,$4) RETURNING ID`

	sqlDeleteUser = `DELETE FROM users 
					 WHERE id = $1`

	sqlUpdateUser = `UPDATE users 
					 SET name = $1, email = $2, updated_at = $3
					 WHERE id = $4
					 RETURNING id`

	sqlFindUserByEmail = `SELECT name, email, password
							FROM users
							WHERE email = $1`

	sqlFindUserByEmailAndPassword = `SELECT id, name, email, password
							FROM users
							WHERE email = $1 and password = $2`
)
