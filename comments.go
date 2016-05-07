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
	CreateTime  int64  `json:"CreateTime"`
	EditTime    int64  `json:"EditTime"`
	CreatorID   string `json:"CreatorID"`
	CreatorName string `json:"CreatorName"`
	PostID      string `json:"PostID"`
	Content     string `json:"Content"`
}

var commentsTable string = "comments"

func createCommentsTable() error {
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
	count, _ := strconv.Atoi(c.Param("count"))

	cur, err := r.DB(dbName).Table(commentsTable).OrderBy(r.OrderByOpts{
		Index: r.Desc("CreateTime"),
	}).Limit(count).Run(session)

	if err != nil {
		return err
	}

	res, err := getAllDataFromCursor(cur)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func getCommentsByLikes(c echo.Context) error {
	// TODO
	return nil
}

func getAllComments(c echo.Context) error {
	cur, err := r.DB(dbName).Table(commentsTable).OrderBy(r.OrderByOpts{
		Index: r.Desc("CreateTime"),
	}).Run(session)

	if err != nil {
		fmt.Printf("Failed to get ordered comments. Error: $s\n", err.Error())
		return err
	}
	if cur == nil {
		errors.New("Error getting all comments. The cursor is nil!")
	}

	res, err := getAllDataFromCursor(cur)
	if err != nil {
		fmt.Printf("Failed to get all data from cursor. Error: $s\n", err.Error())
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func createComment(c echo.Context) error {
	comm := &Comment{
		CreateTime: time.Now().Unix(),
		EditTime:   time.Now().Unix(),
	}

	if err := c.Bind(comm); err != nil {
		return err
	}

	fmt.Printf("Comment to create: %s", comm)
	ans, err := insertToTable(commentsTable, comm)
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

	UserID, ok := data["UserID"].(string)
	if !ok {
		return errors.New("UserID of the comment creator must be supplied in order to delete a comment")
	}
	CommentID, ok := data["CommentID"].(string)
	if !ok {
		err := errors.New("CommentID must be supplied in order to delete a comment")
		fmt.Println(err.Error())
		return err
	}

	res, err := getFromTable(commentsTable, CommentID)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if res == nil {
		err := errors.New("No such comment exists!")
		fmt.Println(err.Error())
		return err
	}
	if res.(map[string]interface{})["CreatorID"].(string) == UserID {
		removed, err := removeFromTable(commentsTable, CommentID)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		return c.JSON(http.StatusOK, removed)
	}

	err = errors.New("Supplied UserID does not match the comment's creator id")
	fmt.Println(err.Error())
	return err
}
