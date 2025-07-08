package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level int    `json:"level"`
}

var uniqueId int = 1
var pokemons []Pokemon

func getPokemons(context *gin.Context) {
	context.JSON(http.StatusOK, pokemons)
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

	pokemons = append(pokemons, newPokemon)

	context.JSON(http.StatusCreated, newPokemon)

}

func updatePokemon(context *gin.Context) {

	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedPokemon Pokemon
	if err := context.BindJSON(&updatedPokemon); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if updatedPokemon.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Name must be non-empty"})
		return
	}

	if updatedPokemon.Type == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Type must be non-empty"})
		return
	}

	if updatedPokemon.Level <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Level must be > 0"})
		return
	}

	for i := range pokemons {
		if pokemons[i].ID == id {
			pokemons[i].Name = updatedPokemon.Name
			pokemons[i].Type = updatedPokemon.Type
			pokemons[i].Level = updatedPokemon.Level

			context.JSON(http.StatusOK, pokemons[i])
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})

}

func deletePokemon(context *gin.Context) {

	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i := range pokemons {
		if pokemons[i].ID == id {

			pokemons = append(pokemons[:i], pokemons[i+1:]...)
			context.JSON(http.StatusOK, gin.H{"Message": "Pokemon released successfully"})
			return
		}
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found in API"})

}

func main() {
	router := gin.Default()

	// GET/pokemons
	router.GET("/pokemons", getPokemons)

	// POST/pokemons using data.json
	router.POST("/pokemons", addPokemon)

	// PUT/pokemons/id
	router.PUT("/pokemons/:id", updatePokemon)

	// DELETE/pokemons/id
	router.DELETE("/pokemons/:id", deletePokemon)

	router.Run("localhost:8080")
}
