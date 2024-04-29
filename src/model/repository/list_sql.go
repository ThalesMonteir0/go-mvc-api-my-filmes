package repository

var (
	findListsByUserIDSQL = `Select * From lists where user_id = $1`
)
