package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"example/hello/morestrings"

	"github.com/google/go-cmp/cmp"

	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/greetings"
	"golang.org/x/example/stringutil"
	"rsc.io/quote"

	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

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

	fmt.Println("Start Server NOW..!")

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/user/:id", getUserByid)

	router.Run("localhost:8080")
}

type User struct {
	ID       string `json:"id"`
	XingMing string `json:"xingMing"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func openDB() *sql.DB {
	fmt.Println("begin open db.....")

	var db *sql.DB
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
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

	return db
}

func getUserByid(c *gin.Context) {
	// var e error
	userId, e := strconv.Atoi(c.Param("id"))
	if e != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not valid userid"})
		return
	}

	var u User

	db := openDB()
	row := db.QueryRow("SELECT id,xingMing,userName,`password` FROM ty_user WHERE id = ?", userId)
	err := row.Scan(&u.ID, &u.XingMing, &u.UserName, &u.Password)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, u)
}

func getUsers(c *gin.Context) {
	db := openDB()

	var users []User

	rows, err := db.Query("SELECT id, xingMing,userName,`password` FROM ty_user limit 20")

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no users"})
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.XingMing, &u.UserName, &u.Password)
		if err != nil {
			fmt.Println(err)
			// return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		// return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	c.IndentedJSON(http.StatusOK, users)
	return
}
