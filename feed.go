package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getFeed(c echo.Context) {
	count, _ := strconv.Atoi(c.Param(count))

	_, err := fetchPosts(count)
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest)
	}
}
