package model

type Cocktail struct {
	Name        string
	Recipe      string
	Ingredients []Ingredient
}

type Ingredient struct {
	Name   string
	Amount int
	Unit   int
}