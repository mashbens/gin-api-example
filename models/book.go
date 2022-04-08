package models

type Book struct {
	ID    int    `json:"id" param:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

type BookInput struct {
	ID    int    `json:"id" param:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}
