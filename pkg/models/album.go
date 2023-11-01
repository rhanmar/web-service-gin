package models

type Album struct {
	ID     int     `json:"id" gorm:"autoIncrement;primaryKey;not null"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
