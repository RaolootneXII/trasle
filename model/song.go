package model

import "github.com/jinzhu/gorm"

// Song - is a song
type Song struct {
	gorm.Model
	Title      string `json:"title" gorm:"not null"`
	Lyrics     string `json:"lyrics" gorm:"not null"`
	Author     string `json:"author" gorm:"not null"`
	AlbumImage string `json:"image" gorm:"not null"`
	SongNumber int    `json:"song_number" gorm:"not null"`
	SongLink   string `json:"song_link" gorm:"not null"`
	Rating     int    `json:"rating" gorm:"not null"`
}
