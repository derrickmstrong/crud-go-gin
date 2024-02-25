package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Define the route for the /albums endpoint
	router.GET("/albums", getAlbums)
	router.GET("albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	// Start serving the application
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // Respond with the list of albums in JSON format
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id") // Get the ID value from the URL

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id { // If the album is found
			c.IndentedJSON(http.StatusOK, a) // Respond with the album in JSON format
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"}) // Respond with an error message in JSON format
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album // Create a new album from the request body

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil { // If there is an error, do not continue
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)            // append() returns a new slice with the new album added
	c.IndentedJSON(http.StatusCreated, newAlbum) // Respond with the new album in JSON format
}

// Example get request
// curl -X GET http://localhost:8080/albums

// Example get by id request
// curl -X GET http://localhost:8080/albums/1

// Example post request
// curl -X POST -d '{"id":"4","title":"The Modern Sound of Betty Carter","artist":"Betty Carter","price":49.99}' -H "Content-Type: application/json" http://localhost:8080/albums
