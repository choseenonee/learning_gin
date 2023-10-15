package model

type Contact struct {
	Id          int
	ContactType string
	Name        string
	Number      string
	Email       string
}

type UpdateContactInput struct {
	ContactType string
	Name        string
	Number      string
	Email       string
}
