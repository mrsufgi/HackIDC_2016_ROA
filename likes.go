package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Like struct {
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}

var likesTable string = "likes"

func CreateLikesTable(c echo.Context) error {
	indices := []string{
		"UserID",
		"CommentID",
	}
	return createTable(likesTable, indices)
}

func likeComment(c echo.Context) error {
	data := make(map[string]string)
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	userId, commentId := "user_id", "comment_id"
	filterMap := map[string]string{
		"UserID":    data[userId],
		"CommentID": data[commentId],
	}
	ans, err := filterFromTable(likesTable, filterMap)

	arr := ans.([]interface{})
	// TODO add verification for user_id and comment_id
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

	return getCommentLikes(c)
}

func getCommentLikes(c echo.Context) error {
	filterMap := map[string]string{
		"CommentID": c.Param("id"),
	}

	ans, err := filterFromTable(likesTable, filterMap)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ans)
}
