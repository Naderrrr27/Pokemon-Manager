package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level int    `json:"level"`
}

var uniqueId int = 1
var Pokemons []Pokemon

func getPokemons(context *gin.Context) {
	context.JSON(http.StatusOK, Pokemons)
}

func addPokemon(context *gin.Context) {

	var newPokemon Pokemon

	if err := context.BindJSON(&newPokemon); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if newPokemon.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Name must be non-empty"})
		return
	}

	if newPokemon.Type == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Type must be non-empty"})
		return
	}

	if newPokemon.Level <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Level must be > 0"})
		return
	}

	newPokemon.ID = uniqueId
	uniqueId++

	Pokemons = append(Pokemons, newPokemon)

	context.JSON(http.StatusCreated, newPokemon)

}

func main() {
	router := gin.Default()

	//GET/pokemons
	router.GET("/pokemons", getPokemons)

	//POST/pokemons using data.json
	router.POST("/pokemons", addPokemon)

	router.Run("localhost:8080")
}
