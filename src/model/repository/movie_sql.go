package repository

var (
	SqlFindMovieByID = `SELECT id, name, genre, description, url_img FROM movies where id = $1`
)
