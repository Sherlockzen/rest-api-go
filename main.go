package main

import (
	"context"
	mongodb "example/rest-api-go/configuration/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// user type on mongoDB
type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Age  int32              `bson:"age"`
}

var albums []album

func main() {
	router := gin.Default()

	//Conection with MongoDB

	//testing insert into collection
	//user := User{
	//	Name: "hewerton",
	//	Age:  30,
	//}
	//
	//result, err := collection.InsertOne(context.Background(), user)
	//if err != nil {
	//	panic(err)
	//}

	//show the users in the userData collection

	router.GET("/users", getAlbums)
	router.POST("/users/post", postAlbums)
	router.GET("/users/:name", getName)
	router.DELETE("/users/:name", deleteName)

	// router.Run("localhost:8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router.Run("0.0.0.0:" + port)
}

func deleteName(c *gin.Context) {
	name := c.Param("name")
	collection := mongodb.InitConnection().Database("users").Collection("userData")

	res, err := collection.DeleteOne(context.Background(), bson.D{{"name", name}})
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, res)
}

func getName(c *gin.Context) {
	nam := c.Param("name")

	collection := mongodb.InitConnection().Database("users").Collection("userData")
	cur, err := collection.Find(context.Background(), bson.D{{"name", nam}})
	if err != nil {
		panic(err)
	}
	var result []User
	if err = cur.All(context.Background(), &result); err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, result)
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {

	albums = []album{}

	//content, er := os.ReadFile("data.json")
	//if er != nil {
	//	log.Fatal(er)
	//}
	//er = json.Unmarshal(content, &albums)

	collection := mongodb.InitConnection().Database("users").Collection("userData")
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	var result []User
	if err = cur.All(context.Background(), &result); err != nil {
		panic(err)
	}
	//fmt.Printf("Usuarios cadastrados: %v\n", result)

	c.IndentedJSON(http.StatusOK, result)
}

func postAlbums(c *gin.Context) {
	//var newAlbum album
	//
	//if err := c.BindJSON(&newAlbum); err != nil {
	//	return
	//}
	//
	//content, er := os.ReadFile("data.json")
	//if er != nil {
	//	log.Fatal(er)
	//}
	//
	//er = json.Unmarshal(content, &albums)
	//
	//albums := append(albums, newAlbum)
	//jsonAlbums, _ := json.Marshal(albums)
	//err := os.WriteFile("data.json", jsonAlbums, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//insert into collection
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		panic(err)
	}
	collection := mongodb.InitConnection().Database("users").Collection("userData")
	res, err := collection.InsertOne(context.Background(), newUser)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, res)
}
