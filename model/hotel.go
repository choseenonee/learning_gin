package model

type Hotel struct {
	Id          int
	Name        string
	LocationId  int
	Number      string
	WorkerId    int
	Description string
}

type UpdateHotelInput struct {
	Name        string
	LocationId  int
	Number      string
	WorkerId    int
	Description string
}
