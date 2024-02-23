package model

type CocktailDetails struct {
	Name        string
	Recipe      string
	Ingredients []IngredientDetails
}

type IngredientDetails struct {
	Name   string
	Amount int
	Unit   int
}