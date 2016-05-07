package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Like struct {
	UserID    string `json:"userID"`
	CommentID string `json:"commentID"`
}

var likesTable string = "likes"

func createLikesTable() error {
	indices := []string{
		"userID",
		"commentID",
	}
	return createTable(likesTable, indices)
}

func likeComment(c echo.Context) error {
	data := make(map[string]string)
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	userId, commentId := "userID", "commentID"
	filterMap := map[string]string{
		"userID":    data[userId],
		"commentID": data[commentId],
	}
	ans, err := filterFromTable(likesTable, filterMap)

	arr := ans.([]interface{})
	// TODO add verification for UserID and CommentID
	if len(arr) == 0 {
		like := &Like{
			UserID:    data[userId],
			CommentID: data[commentId],
		}
		_, err = insertToTable(likesTable, like)
	} else {
		// All this just to get the uid from the returned object :(
		id := ans.([]interface{})[0].(map[string]interface{})["id"].(string)
		_, err = removeFromTable(likesTable, id)
	}

	if err != nil {
		return err
	}

	// Gotta set this so c.Param("id") will work
	// in getcommentLikes()
	c.SetParamNames("id")
	c.SetParamValues(data[commentId])

	return getCommentLikes(c)
}

func getCommentLikes(c echo.Context) error {
	filterMap := map[string]string{
		"commentID": c.Param("id"),
	}

	ans, err := filterFromTable(likesTable, filterMap)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ans)
}
