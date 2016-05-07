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

type Post struct {
	createTime  int64
	editTime    int64
	creatorId   string
	creatorName string
	title       string
	imageUrl    string
}

var postsTable string = "posts"

func createPostsTable() error {
	indices := []string{
		createTime,
		editTime,
		creatorId,
		creatorName,
		title,
		imageUrl,
	}
	return createTable(postsTable, indices)
}

func getLastPosts(c echo.Context) error {
	num, _ := strconv.Atoi(c.Param(count))

	cur, err := r.DB(dbName).Table(postsTable).OrderBy(r.OrderByOpts{
		Index: r.Desc(createTime),
	}).Limit(num).Run(session)

	if err != nil {
		return err
	}

	res, err := getAllDataFromCursor(cur)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func getAllPosts(c echo.Context) error {
	cur, err := r.DB(dbName).Table(postsTable).OrderBy(r.OrderByOpts{
		Index: r.Desc(createTime),
	}).Run(session)

	if err != nil {
		return err
	}
	if cur == nil {
		return c.JSON(http.StatusNoContent,
			errors.New("Error getting all posts. The cursor is nil!"))
	}

	res, err := getAllDataFromCursor(cur)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func getPostComments(c echo.Context) error {
	filterMap := map[string]string{
		postId: c.Param(id),
	}

	ans, err := filterFromTable(commentsTable, filterMap)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ans)
}

func createPost(c echo.Context) error {
	p := &Post{
		createTime: time.Now().Unix(),
		editTime:   time.Now().Unix(),
	}

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	ans, err := insertToTable(postsTable, p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusCreated, ans)
}

func getPost(c echo.Context) error {
	fmt.Println(c.Param(id))
	ans, err := getFromTable(postsTable, c.Param(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(ans)
	return c.JSON(http.StatusOK, ans)
}

func editPost(c echo.Context) error {
	data := make(map[string]interface{})
	err := c.Bind(&data)
	if err != nil {
		return err
	}
	data[editTime] = time.Now().Unix()
	// TODO filter data to contain only existing fields, nothing new
	res, err := updateFieldInTable(postsTable, data[id].(string), data)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func deletePost(c echo.Context) error {
	data := make(map[string]interface{})
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	UserID, ok := data[userId].(string)
	if !ok {
		return errors.New("UserID of the post creator must be supplied in order to delete a post")
	}
	PostID, ok := data[postId].(string)
	if !ok {
		return errors.New("PostID must be supplied in order to delete a post")
	}

	res, err := getFromTable(postsTable, PostID)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("No such post exists!")
	}
	if res.(map[string]interface{})[creatorId].(string) == UserID {
		removed, err := removeFromTable(postsTable, PostID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, removed)
	}

	return errors.New("Supplied UserID does not match the post's creator id")
}
