package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rpratama-dev/sample-api/src/models"
	"github.com/rpratama-dev/sample-api/src/service"
)

type PostController struct {
	httpClient *service.HttpClient
}

func newPostController(hc *service.HttpClient) *PostController {
	return &PostController {
		httpClient: hc,
	}
}

func IndexPost(c echo.Context) error {
	apiService := service.APIService()
	postController := newPostController(apiService)

	dataPosts, err := postController.httpClient.GetPost()
	if err != nil {
		return err
	}
	// Will execute after returning last code executed
	defer dataPosts.Close()

	var posts []models.Post
	err = models.ConvertJSONToPosts(dataPosts, &posts)

	if err != nil {
		// Handle error if the JSON decoding fails
		return err
	}

	// Return the posts as JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get post",
		"data": posts,
	})

}

func StorePost(c echo.Context) error {
	// Create a new instance of the User struct
	postBody := new(models.Post)

	// Bind the JSON body to the User struct
	if err := c.Bind(postBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Convert the Post object to JSON
	jsonData, _ := json.Marshal(postBody)

	apiService := service.APIService()
	postController := newPostController(apiService)

	dataPost, err := postController.httpClient.StorePost(jsonData)
	if (err != nil) {
		fmt.Println("error", err);
	}
	// Will execute after returning last code executed
	defer dataPost.Close()

	var post models.Post
	err = models.ConvertJSONToPost(dataPost, &post)

	if err != nil {
		// Handle error if the JSON decoding fails
		return err
	}
	// Return the post as JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success insert new post",
		"data": post,
	})
}

func ShowPost(c echo.Context) error {
	apiService := service.APIService()
	postController := newPostController(apiService)

	postId := c.Param("id")
	dataPosts, err := postController.httpClient.GetPostById(postId)
	if err != nil {
		return err
	}
	// Will execute after returning last code executed
	defer dataPosts.Close()

	var posts models.Post
	err = models.ConvertJSONToPost(dataPosts, &posts)
	if (posts.ID == 0) {
		// Return the posts as JSON response
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Item you'r looking for doesn't exists",
			"data": nil,
		})
	}
	if err != nil {
		// Handle error if the JSON decoding fails
		return err
	}

	// Return the posts as JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get post by id",
		"data": posts,
	})
}

func DestroyPost(c echo.Context) error {
	apiService := service.APIService()
	postController := newPostController(apiService)

	postId := c.Param("id")
	dataPosts, err := postController.httpClient.DestroyPostById(postId)
	if err != nil {
		return err
	}
	// Will execute after returning last code executed
	defer dataPosts.Close()
	// Return the posts as JSON response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete post by id",
		"data": nil,
	})
}
