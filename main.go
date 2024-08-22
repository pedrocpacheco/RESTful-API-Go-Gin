package main

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Amado Amigo", Artist: "Guga Xavier", Price: 38.40},
	{ID: "3", Title: "Fogo em Meus Olhos", Artist: "Alessandro Villas Boas", Price: 18.40},
}
