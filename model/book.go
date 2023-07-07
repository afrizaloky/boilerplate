package model

type Book struct {
	ID     int    `gorm:"primaryKey,autoIncrement" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}
