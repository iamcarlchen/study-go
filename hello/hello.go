package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/greetings"
	"golang.org/x/example/stringutil"
	"rsc.io/quote"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// fmt.Println("hello")

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println(quote.Go())
	message, err := greetings.Hellos([]string{"Gladys", "Samantha", "Darrin"})

	fmt.Println(message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stringutil.Reverse("Hello"))

	fmt.Println("Hello End.")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
