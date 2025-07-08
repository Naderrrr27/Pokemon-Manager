package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Pokemon struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level string `json:"level"`
}

var Pokemons []Pokemon

func getPokemons(context *gin.Context) {
	context.JSON(http.StatusOK, Pokemons)
}

func main() {
	router := gin.Default()

	//Get/pokemons
	router.GET("/pokemons", getPokemons)

	router.Run("localhost:8080")
}
