package model

type movieDomain struct {
	id          int
	name        string
	genre       string
	description string
	urlImg      string
}

func (md *movieDomain) GetName() string {
	return md.name
}

func (md *movieDomain) GetGenre() string {
	return md.genre
}

func (md *movieDomain) GetDescription() string {
	return md.description
}

func (md *movieDomain) GetUrlImg() string {
	return md.urlImg
}

func (md *movieDomain) GetID() int {
	return md.id
}

func (md *movieDomain) SetID(id int) {
	md.id = id
}
