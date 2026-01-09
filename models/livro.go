package models

type Livro struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"column:title; not null"`
	Author string `json:"author" gorm:"column:author; not null"`
	PublishedYear int `json:"published_year" gorm:"column:published_year; not null"`
}