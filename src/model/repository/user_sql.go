package repository

var (
	sqlFindUserByID = `SELECT id, name, email FROM users WHERE id = $1 `
)
