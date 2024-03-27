package repository

var (
	sqlFindUserByID = `SELECT id, name, email 
					   FROM users 
					   WHERE id = $1 `

	sqlCreateUser = `INSERT INTO users (email, password, name) 
						values ($1,$2,$3); returning id`

	sqlDeleteUser = `DELETE FROM users 
					 WHERE id = $1`

	sqlUpdateUser = `UPDATE users 
					 SET name = $1, email = $2, password = $3;
					 WHERE id = $4
					 RETURNING id`
)
