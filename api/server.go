package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id" binding:"required"`
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
} // postAlbums adds an album from JSON received in the request body.

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		fmt.Println("======================")
		fmt.Println(err)
		RespondWithValidationError(c)
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func postAlbumsWithoutResponseFunction(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		fmt.Println("#################")
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, nil)
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func GetApp() *gin.Engine {
	router := gin.Default()
	router.POST("/albums", postAlbums)
	router.POST("/albums2", postAlbumsWithoutResponseFunction)
	return router
}
