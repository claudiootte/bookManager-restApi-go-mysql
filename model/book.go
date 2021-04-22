package model

import "gorm.io/gorm"

// Define-se o modelo do que irá se tratar na aplicação

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}
