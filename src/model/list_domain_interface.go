package model

type ListDomainInterface interface {
	GetID() int
	SetID(id int)
	GetTitle() string
	SetTitle(title string)
	GetDescription() string
	SetDescription(desc string)
	GetUserID() int
	SetUserID(userId int)
}

func NewListDomain(title, desc string, userID int) ListDomainInterface {
	return &listDomain{
		title: title, description: desc, userID: userID,
	}
}
