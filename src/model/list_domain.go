package model

type listDomain struct {
	id          int
	title       string
	description string
	userID      int
}

func (ld *listDomain) GetID() int {
	return ld.id
}
func (ld *listDomain) SetID(id int) {
	ld.id = id
}
func (ld *listDomain) GetTitle() string {
	return ld.title
}
func (ld *listDomain) SetTitle(title string) {
	ld.title = title
}
func (ld *listDomain) GetDescription() string {
	return ld.description
}
func (ld *listDomain) SetDescription(desc string) {
	ld.description = desc
}
func (ld *listDomain) GetUserID() int {
	return ld.userID
}
func (ld *listDomain) SetUserID(userId int) {
	ld.userID = userId
}
