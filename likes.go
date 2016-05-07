package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Like struct {
	userId    string
	commentId string
}

var likesTable string = "likes"

func createLikesTable() error {
	indices := []string{
		userId,
		commentId,
	}
	return createTable(likesTable, indices)
}

func likeComment(c echo.Context) error {
	data := make(map[string]string)
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	userIdLocal, commentIdLocal := userId, commentId
	filterMap := map[string]string{
		userId:    data[userIdLocal],
		commentId: data[commentIdLocal],
	}
	ans, err := filterFromTable(likesTable, filterMap)

	arr := ans.([]interface{})
	// TODO add verification for UserID and CommentID
	if len(arr) == 0 {
		like := &Like{
			userId:    data[userId],
			commentId: data[commentId],
		}
		_, err = insertToTable(likesTable, like)
	} else {
		// All this just to get the uid from the returned object :(
		idLocal := ans.([]interface{})[0].(map[string]interface{})[id].(string)
		_, err = removeFromTable(likesTable, idLocal)
	}

	if err != nil {
		return err
	}

	// Gotta set this so c.Param("id") will work
	// in getcommentLikes()
	c.SetParamNames("id")
	c.SetParamValues(data[commentIdLocal])

	return getCommentLikes(c)
}

func getCommentLikes(c echo.Context) error {
	filterMap := map[string]string{
		commentId: c.Param(id),
	}

	ans, err := filterFromTable(likesTable, filterMap)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ans)
}
