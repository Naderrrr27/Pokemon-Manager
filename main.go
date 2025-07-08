package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID    int    `json:"id"`    // Unique ID assigned by the server
	Name  string `json:"name"`  // Name of the Pokemon
	Type  string `json:"type"`  // Type (e.g., Fire, Water)
	Level int    `json:"level"` // Level (must be > 0)
}

var uniqueId int = 1
var pokemons = make(map[int]Pokemon)

// Just converting map values to a slice for JSON response
func getPokemons(context *gin.Context) {

	returnedPokemons := make([]Pokemon, 0, len(pokemons))

	for _, pokemon := range pokemons {
		returnedPokemons = append(returnedPokemons, pokemon)
	}

	context.JSON(http.StatusOK, returnedPokemons)

}

// Validate and add a new Pokemon with a unique ID
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

	pokemons[newPokemon.ID] = newPokemon

	context.JSON(http.StatusCreated, newPokemon)

}

// Validate and update Pokemon info by ID if it exists
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

	if _, exists := pokemons[id]; exists {
		updatedPokemon.ID = id
		pokemons[id] = updatedPokemon

		context.JSON(http.StatusOK, pokemons[id])
		return
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})

}

// Remove Pokemon from the map if the ID exists
func deletePokemon(context *gin.Context) {

	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, exists := pokemons[id]; exists == true {
		delete(pokemons, id)
		context.JSON(http.StatusOK, gin.H{"Message": "Pokemon released successfully"})
		return
	}

	context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found in API"})

}

// Fetch basic Pokemon info from external PokeAPI by name
func getPokemonInfo(context *gin.Context) {

	name := context.Param("name")
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strings.ToLower(name))

	if err != nil || response.StatusCode != http.StatusOK {
		context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found in API"})
		return
	}

	defer response.Body.Close()

	var apiResponse struct {
		Name   string `json:"name"`
		Height int    `json:"height"`
		Weight int    `json:"weight"`
		Types  []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
	}

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse API response"})
		return
	}

	var typeNames []string
	for _, t := range apiResponse.Types {
		typeNames = append(typeNames, t.Type.Name)
	}

	context.JSON(http.StatusOK, gin.H{
		"name":   apiResponse.Name,
		"height": apiResponse.Height,
		"weight": apiResponse.Weight,
		"types":  typeNames,
	})

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

	// GET/pokemon-info/id -> External API EndPoint
	router.GET("/pokemon-info/:name", getPokemonInfo)

	router.Run("localhost:8080")
}
