package model

type Contact struct {
	Id          int    `json:"id"`
	ContactType string `json:"contact_type" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Number      string `json:"number" binding:"required"`
	Email       string `json:"email" binding:"required"`
}
