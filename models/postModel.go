package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body string
}

// migrate the db class model => this creates tables
