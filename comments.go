package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo"
)

type Comment struct {
	CreateTime  int64  `json:"create_time"`
	EditTime    int64  `json:"edit_time"`
	CreatorID   string `json:"creator_id"`
	CreatorName string `json:"creator_name"`
	PostID      string `json:"post_id"`
	Content     string `json:"content"`
}

var commentsTable string = "comments"

func createCommentsTable(c echo.Context) error {
	indices := []string{
		"CreateTime",
		"EditTime",
		"CreatorID",
		"CreatorName",
		"PostID",
		"Content",
	}
	return createTable(commentsTable, indices)
}

func getLastComments(c echo.Context) error {
	num, _ := strconv.Atoi(c.Param("num"))

	cur, err := r.DB(dbName).Table(commentsTable).OrderBy(r.OrderByOpts{
		Index: r.Desc("CreateTime"),
	}).Run(session)

	if err != nil {
		return err
	}

	res, err := getDataFromCursor(cur, num)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func getAllComments(c echo.Context) error {
	cur, err := r.DB(dbName).Table(commentsTable).OrderBy(r.OrderByOpts{
		Index: r.Desc("CreateTime"),
	}).Run(session)

	if err != nil {
		return err
	}
	if cur == nil {
		errors.New("Error getting all comments. The cursor is nil!")
	}

	res, err := getAllDataFromCursor(cur)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func createComment(c echo.Context) error {
	p := &Comment{
		CreateTime: time.Now().Unix(),
		EditTime:   time.Now().Unix(),
	}

	if err := c.Bind(p); err != nil {
		return err
	}

	ans, err := insertToTable(commentsTable, p)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, ans)
}

func getComment(c echo.Context) error {
	fmt.Println(c.Param("id"))
	ans, err := getFromTable(commentsTable, c.Param("id"))
	if err != nil {
		return err
	}

	fmt.Println(ans)
	return c.JSON(http.StatusOK, ans)
}

func editComment(c echo.Context) error {
	data := make(map[string]interface{})
	err := c.Bind(&data)
	if err != nil {
		return err
	}
	data["EditTime"] = time.Now().Unix()
	// TODO filter data to contain only existing fields, nothing new
	res, err := updateFieldInTable(commentsTable, data["id"].(string), data)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func deleteComment(c echo.Context) error {
	data := make(map[string]interface{})
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	UserID, ok := data["user_id"].(string)
	if !ok {
		return errors.New("user_id of the comment creator must be supplied in order to delete a comment")
	}
	CommentID, ok := data["comment_id"].(string)
	if !ok {
		return errors.New("comment_id must be supplied in order to delete a comment")
	}

	res, err := getFromTable(commentsTable, CommentID)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("No such comment exists!")
	}
	if res.(map[string]interface{})["CreatorID"].(string) == UserID {
		removed, err := removeFromTable(commentsTable, CommentID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, removed)
	}

	return errors.New("Supplied user_id does not match the comment's creator id")
}
