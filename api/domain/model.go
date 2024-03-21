package domain

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

type CocktailByIngredient struct {
	IngredientName string
	CocktailName []CocktailName
}

type CocktailName struct {
	Name string
}