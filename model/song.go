package model

import "github.com/jinzhu/gorm"

// Song - is a song
type Song struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Lyrics string `json:"lyrics" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Rating int    `json:"rating" gorm:"not null"`
}
