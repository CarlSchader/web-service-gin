package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums =[]album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
 
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	checkError(err)

	numericId, err := strconv.Atoi(albums[len(albums) - 1].ID)
	checkError(err)

	numericId += 1

	newAlbum.ID = strconv.Itoa(numericId)

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	checkError(err)

	if id >= len(albums) || id < 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "index out of range"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums[id - 1])
}

func main() {
	router := gin.Default()
	
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.GET("/albums/:id", getAlbumById)

	router.Run("localhost:8080")
}