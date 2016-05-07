package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

var defaultCount int = 5

func getFeed(c echo.Context) error {
	count, _ := strconv.Atoi(c.Param(count))

	posts, err := fetchPosts(count)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//	debugPrinter("All posts", posts)
	postsMap := posts.([]interface{})
	for _, post := range postsMap {
		//		debugPrinter("Post in the loop", post)
		comments, err := fetchComments(post.(map[string]interface{})[id].(string), defaultCount)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		//		debugPrinter("All fetched comments", comments)

		//		debugPrinter("Post after updating with comments", post)
		commentsMap := comments.([]interface{})
		for _, comment := range commentsMap {
			//			debugPrinter("Comment inside comments loop", comment)
			likes, err := fetchCommentLikes(comment.(map[string]interface{})[id].(string))

			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			//			debugPrinter("Found likes", likes)
			comment.(map[string]interface{})["Likes"] = likes

			//			debugPrinter("comment after adding likes", comment)
		}
		post.(map[string]interface{})["Comments"] = comments
	}
	//	debugPrinter("Posts after updating all fields", posts)
	return c.JSON(http.StatusOK, posts)
}

func debugPrinter(title string, element interface{}) {
	fmt.Println(title)
	fmt.Printf("Print element as is: %s Type: %T\n", element, element)
	res, _ := json.Marshal(element)

	fmt.Printf("Print as JSON: %s Type: %T\n", res, res)
}
