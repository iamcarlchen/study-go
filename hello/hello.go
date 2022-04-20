package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/greetings"
	"golang.org/x/example/stringutil"
	"rsc.io/quote"

	"database/sql"
	"github.com/go-sql-driver/mysql"
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

var db *sql.DB

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


	fmt.Println("begin data access logic")
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "qazWSXedc",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "db_ty",
	}
	// Get a database handle.
	var errDB error
	db, errDB = sql.Open("mysql", cfg.FormatDSN())
	if errDB != nil {
		log.Fatal(errDB)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")


	fmt.Println("Start Server NOW..!")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
