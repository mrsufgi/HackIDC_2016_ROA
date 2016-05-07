package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Like struct {
	UserID    string `json:"UserID"`
	CommentID string `json:"CommentID"`
}

var likesTable string = "likes"

func createLikesTable() error {
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

	userId, commentId := "UserID", "CommentID"
	filterMap := map[string]string{
		"UserID":    data[userId],
		"CommentID": data[commentId],
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

	// Gotta set this so c.Param("id") will work in
	// getcommentLikes()
	c.SetParamNames("id")
	c.SetParamValues(data[commentId])

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
