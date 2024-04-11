package repository

var (
	SqlFindMovieByID   = `SELECT id, name, genre, description, url_img FROM movies where id = $1`
	SqlFindMovieByName = `SELECT id, name, genre, description, url_img FROM movies where name = $1`
	SqlUpdateMovie     = `UPDATE movies SET name = $1, genre = $2, description = $3, url_img = $4 WHERE id = $5 RETURNING id`
)
