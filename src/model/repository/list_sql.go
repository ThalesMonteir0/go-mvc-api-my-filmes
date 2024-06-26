package repository

var (
	findListsByUserIDSQL = `Select * From lists where user_id = $1`
	createListSQL        = `INSERT INTO public.lists(
	title, description, created_at, user_id)
	VALUES ($1, $2, $3, $4) RETURNING ID;`
	findListByIDSQL = `Select * from lists where id = $1`
	deleteListByID  = `DELETE FROM lists where id = $1`
)
