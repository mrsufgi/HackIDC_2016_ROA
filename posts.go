package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	ID   string `json:"user_id"`
}

type Post struct {
	ID         int               `json:"post_id"`
	CreateTime time.Time         `json:"create_time"`
	EditTime   time.Time         `json:"edit_time"`
	CreatorID  string            `json:"creator_id"`
	Content    string            `json:"content"`
	Likes      map[string]string `json:"likes"`
}

var (
	posts = map[int]*Post{}
	seq   = 1
)

func getAllPosts(c echo.Context) error {
	tmpPosts := map[string]*Post{}
	keys := make([]int, 0, len(posts))
	for k := range posts {
		keys = append(keys, k)
	}

	for _, i := range keys {
		tmpPosts[strconv.Itoa(i)] = posts[i]
		fmt.Println(posts[i].Likes)
	}
	return c.JSON(http.StatusOK, tmpPosts)
}

func likePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	if posts[id] != nil {
		// TODO check DB if uid is actual user id
		_, ok := posts[id].Likes[u.ID]
		if !ok {
			posts[id].Likes[u.ID] = u.Name
		} else {
			delete(posts[id].Likes, u.ID)
		}
	}

	return c.JSON(http.StatusOK, posts[id])
}

func createPost(c echo.Context) error {
	p := &Post{
		ID:         seq,
		CreateTime: time.Now(),
		Likes:      map[string]string{},
	}
	if err := c.Bind(p); err != nil {
		return err
	}

	posts[p.ID] = p
	seq++
	return c.JSON(http.StatusCreated, p)
}

func getPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, posts[id])
}

func editPost(c echo.Context) error {
	editor := new(User)
	if err := c.Bind(editor); err != nil {
		return err
	}
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if editor.ID == posts[id].CreatorID {
		posts[id].Content = p.Content
		posts[id].EditTime = time.Now()
		return c.JSON(http.StatusOK, posts[id])
	} else {
		return c.JSON(http.StatusUnauthorized, posts[id])
	}
}

func deletePost(c echo.Context) error {
	creator := new(User)
	if err := c.Bind(creator); err != nil {
		return err
	}
	// TODO check post id same as user id
	id, _ := strconv.Atoi(c.Param("id"))
	delete(posts, id)
	return c.NoContent(http.StatusNoContent)
}
