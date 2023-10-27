package model

type Hotel struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	LocationId  int    `json:"location_id" binding:"required"`
	Number      string `json:"number" binding:"required"`
	WorkerId    int    `json:"worker_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type HotelWithContact struct {
	Id          int     `json:"title"`
	Name        string  `json:"name"`
	LocationId  int     `json:"location_id"`
	Number      string  `json:"number"`
	Worker      Contact `json:"worker"`
	Description string  `json:"description"`
}
