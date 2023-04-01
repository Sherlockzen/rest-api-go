package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {
	router := gin.Default()

 //  albums := []album{}

	// content, er := ioutil.ReadFile("dat.json")
	// if er != nil {
	// 	log.Fatal(er)
	// }

	// er = json.Unmarshal(content, &albums)

	router.GET("/albums", getAlbums)

	// router.Run("localhost:8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router.Run("0.0.0.0:" + port)
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {

  albums := []album{}

	content, er := ioutil.ReadFile("data.json")
	if er != nil {
		log.Fatal(er)
	}

	er = json.Unmarshal(content, &albums)

	c.IndentedJSON(http.StatusOK, albums)
}
