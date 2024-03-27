package repository

var (
	sqlFindUserByID = `SELECT id, name, email 
					   FROM users 
					   WHERE id = $1 `
	sqlCreateUser = `INSERT INTO users (email, password, name) 
						values ($1,$2,$3); returning id`
	sqlDeleteUser = `DELETE FROM users 
					 WHERE id = $1`
)
