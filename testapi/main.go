package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

type Error struct {
	Error string `json:"error"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist:" John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(context *gin.Context) {
	id := context.Query("id")

	if (id != "") {
		for _, album := range albums {
			if (album.ID == id) {
				context.JSON(http.StatusOK, album)
				return
			}
		}
		error := Error{ Error: "Album not found"}
		context.JSON(http.StatusNotFound, error)
		return
	}

	context.JSON(http.StatusOK, albums)
}

func getAlbum(context *gin.Context) {
	id := context.Param("id")

	if (id != "") {
		for _, album := range albums {
			if (album.ID == id) {
				context.JSON(http.StatusOK, album)
				return
			}
		}
	}
	error := Error{ Error: "Album ID not found"}
	context.JSON(http.StatusNotFound, error)
}

func postAlbum(context *gin.Context) {
	var newAlbum Album

	if err := context.BindJSON(&newAlbum); err != nil {
		log.Println(err)
		return
	}

	albums = append(albums, newAlbum)
	context.JSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbum)
	router.POST("/album", postAlbum)

	router.Run("localhost:8080")
}
