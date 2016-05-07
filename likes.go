package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Like struct {
	UserId    string `json:"UserId" gorethink:"UserId"`
	CommentId string `json:"CommentId" gorethink:"CommentId"`
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

	filterMap := map[string]string{
		userId:    data[userId],
		commentId: data[commentId],
	}
	ans, err := filterFromTable(likesTable, filterMap)

	arr := ans.([]interface{})
	// TODO add verification for UserID and CommentID
	if len(arr) == 0 {
		like := &Like{
			UserId:    data[userId],
			CommentId: data[commentId],
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
	c.SetParamValues(data[commentId])

	return getCommentLikes(c)
}

func getCommentLikes(c echo.Context) error {
	ans, err := fetchCommentLikes(c.Param(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return c.JSON(http.StatusOK, ans)
}

func fetchCommentLikes(id string) (interface{}, error) {
	filterMap := map[string]string{
		commentId: id,
	}

	return filterFromTable(likesTable, filterMap)
}
