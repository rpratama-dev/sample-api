package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rpratama-dev/sample-api/src/models"
)

// slice of User - users data
var users = []models.User{};

func IndexUser(c echo.Context) error {
	// render JSON response with success message and the new data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get users",
		"users": users,
	})
}

func StoreUser(c echo.Context) error {
	// binding data
	user := models.User{}
	c.Bind(&user)

	// if length of slice users less then 1 id = 1, else id +1 from last id in data slice users
	if len(users) < 1 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	// append new data to slice users
	users = append(users, user)

	// render JSON response with success message and the new data
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}

func ShowUser(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))
	message := "User Not Found"
	status := http.StatusNotFound
	var user *models.User
	// find data by id, then render data - JSON response
	for _, v := range users {
	fmt.Print("id",  v.Id == id)
		if v.Id == id {
			user = &v
			message = "Success user found"
			status = http.StatusOK
			break;
		}
	}

	return c.JSON(status, map[string]interface{}{
		"messages": message,
		"user": user,
	})
}

func UpdateUser(c echo.Context) error {
	// catch path param from url
	id, _ := strconv.Atoi(c.Param("id"))
	message := "User Not Found";

	// binding data
	input := models.User{}
	c.Bind(&input)

	var user *models.User = nil
	// find data and change the value, then render JSON response
	for i, v := range users {
		if v.Id == id {
			updatedUser := models.User{
				Id:      v.Id,
				Name:    input.Name,
				Email:   input.Email,
				Address: input.Address,
			}
			users[i] = updatedUser
			user = &users[i]
			message = "success updated user"
			break;
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": message,
		"user": user,
	})
}
