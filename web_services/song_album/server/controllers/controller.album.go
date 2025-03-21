package controllers

import (
	"context"
	"go-practice/models"
	"go-practice/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Generate random UUID
func GenerateUUID() string {
	return uuid.New().String()
}

func getMongoService(c *gin.Context) (*services.MongoService, bool) {
	mongoInstance, err := services.NewMongoService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to create mongo service"})
		return nil, false
	}
	return mongoInstance, true
}

// GetAlbums - handles GET requests to retrieve the list of albums.
// It responds with a JSON-encoded list of albums and an HTTP status code 200 (OK).
// @param c *gin.Context - the context for the current request, provided by the Gin framework.
func GetAlbums(c *gin.Context) {
	mongoInstance, ok := getMongoService(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	filter := bson.M{"isdeleted": bson.M{"$ne": true}}
	cursor, err := mongoInstance.FindAll("albums", filter)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve albums"})
		return
	}
	defer cursor.Close(context.Background())

	var albums []models.Album
	if err = cursor.All(context.Background(), &albums); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to decode albums"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album
	newAlbum.ID = GenerateUUID()
	newAlbum.IsDeleted = false

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	mongoInstance, ok := getMongoService(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	_, err := mongoInstance.Insert("albums", newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to add album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id parameter sent by the client.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// check if ID is valid. safety check.
	if !services.IsValidUUID(id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album ID"})
		return
	}

	sanitizedUUID, err := services.SanitizeUUID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mongoInstance, ok := getMongoService(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	var album models.Album
	filter := bson.M{"id": sanitizedUUID}
	err = mongoInstance.FindOne("albums", filter).Decode(&album)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

// UpdateAlbum updates an existing album with the provided ID.
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	// check if ID is valid. safety check.
	if !services.IsValidUUID(id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album ID"})
		return
	}
	sanitizedUUID, err := services.SanitizeUUID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedAlbum models.Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	mongoInstance, ok := getMongoService(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	filter := bson.M{"id": sanitizedUUID}
	update := bson.M{"$set": updatedAlbum}

	result, err := mongoInstance.Update("albums", filter, update)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to update album"})
		return
	}

	if result.MatchedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedAlbum)
}

// DeleteAlbum marks an album as deleted with the provided ID.
func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	// check if ID is valid. safety check.
	if !services.IsValidUUID(id) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album ID"})
		return
	}
	sanitizedUUID, err := services.SanitizeUUID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	mongoInstance, ok := getMongoService(c)
	if !ok {
		return
	}
	defer mongoInstance.Disconnect()

	filter := bson.M{"id": sanitizedUUID}
	update := bson.M{"$set": bson.M{"isdeleted": true}}

	result, err := mongoInstance.Update("albums", filter, update)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to delete album"})
		return
	}

	if result.MatchedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album marked as deleted"})
}
