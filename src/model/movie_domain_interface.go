package model

type MovieDomainInterface interface {
	GetName() string
	GetGenre() string
	GetDescription() string
	GetUrlImg() string
	GetID() int
	SetID(id int)
	TransformNameToUpperCase()
	TransformDescriptionToUpperCase()
	TransformGenreToUpperCase()
}

func NewMovieDomain(name, genre, description, urlImg string) MovieDomainInterface {
	return &movieDomain{
		id:          0,
		name:        name,
		genre:       genre,
		description: description,
		urlImg:      urlImg,
	}
}
