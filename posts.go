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

type Like struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

type Post struct {
	CreateTime int64             `json:"create_time"`
	EditTime   int64             `json:"edit_time"`
	CreatorID  string            `json:"creator_id"`
	Content    string            `json:"content"`
	Likes      map[string]string `json:"likes"`
}

var postsTable string = "posts"
var likesTable string = "likes"

func CreatePostsTable(c echo.Context) error {
	// TODO
	return nil
}

func CreateLikesTable(c echo.Context) error {
	// TODO
	return nil
}

func getLastPosts(c echo.Context) error {
	num, _ := strconv.Atoi(c.Param("num"))

	cur, err := r.DB("test").Table(postsTable).OrderBy(r.OrderByOpts{
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

func getAllPosts(c echo.Context) error {
	cur, err := r.DB(dbName).Table(postsTable).OrderBy(r.OrderByOpts{
		Index: r.Desc("CreateTime"),
	}).Run(session)

	if err != nil {
		return err
	}
	if cur == nil {
		errors.New("Error getting all posts. The cursor is nil!")
	}

	res, err := getAllDataFromCursor(cur)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func likePost(c echo.Context) error {
	data := make(map[string]string)
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	userId, postId := "user_id", "post_id"
	filterMap := map[string]string{
		"UserID": data[userId],
		"PostID": data[postId],
	}
	ans, err := filterFromTable(likesTable, filterMap)

	arr := ans.([]interface{})
	// TODO add verification for user_id and post_id
	if len(arr) == 0 {
		like := &Like{
			UserID: data[userId],
			PostID: data[postId],
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

	return getPostLikes(c)
}

func getPostLikes(c echo.Context) error {
	filterMap := map[string]string{
		"PostID": c.Param("id"),
	}

	ans, err := filterFromTable(likesTable, filterMap)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ans)
}

func createPost(c echo.Context) error {
	p := &Post{
		CreateTime: time.Now().Unix(),
		EditTime:   time.Now().Unix(),
	}

	if err := c.Bind(p); err != nil {
		return err
	}

	ans, err := insertToTable(postsTable, p)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, ans)
}

func getPost(c echo.Context) error {
	fmt.Println(c.Param("id"))
	ans, err := getFromTable(postsTable, c.Param("id"))
	if err != nil {
		return err
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
	// TODO filter data to contain only existing fields, nothing new
	res, err := updateFieldInTable(postsTable, data["id"].(string), data)
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

	UserID, ok := data["user_id"].(string)
	if !ok {
		return errors.New("user_id of the post creator must be supplied in order to delete a post")
	}
	PostID, ok := data["post_id"].(string)
	if !ok {
		return errors.New("post_id must be supplied in order to delete a post")
	}

	res, err := getFromTable(postsTable, PostID)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("No such post exists!")
	}
	if res.(map[string]interface{})["CreatorID"].(string) == UserID {
		removed, err := removeFromTable(postsTable, PostID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, removed)
	}

	return errors.New("Supplied user_id does not match the post's creator id")
}
