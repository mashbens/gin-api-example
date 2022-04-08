package models

type Book struct {
	ID    int
	Title string
	Price int
}

type BookInput struct {
	ID    int    `json:"id" param:"id" `
	Title string `json:"title" binding:"required" `
	Price int    `json:"price" binding:"required,number" `
}
