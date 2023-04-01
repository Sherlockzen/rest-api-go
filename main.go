package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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

var albums []album

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums/post", postAlbums)

	// router.Run("localhost:8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router.Run("0.0.0.0:" + port)
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {

	albums = []album{}

	content, er := os.ReadFile("data.json")
	if er != nil {
		log.Fatal(er)
	}

	er = json.Unmarshal(content, &albums)

	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	content, er := os.ReadFile("data.json")
	if er != nil {
		log.Fatal(er)
	}

	er = json.Unmarshal(content, &albums)

	albums := append(albums, newAlbum)
	jsonAlbums, _ := json.Marshal(albums)
	err := os.WriteFile("data.json", jsonAlbums, 0666)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
