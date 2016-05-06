package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// TODO remove after Ori pushes users.go
type User struct {
	Name string `json:"name"`
	ID   int64  `json:"user_id"`
}

type Likes struct {
	likes map[string]bool
}

func (set *Likes) like(uid int64) {
	// TODO check DB if uid is actual user id
	suid := strconv.FormatInt(uid, 10)
	set.likes[suid] = !set.likes[suid]
}

type Post struct {
	ID         int64     `json:"post_id"`
	CreateTime time.Time `json:"create_time"`
	EditTime   time.Time `json:"edit_time"`
	CreatorID  int64     `json:"creator_id"`
	Content    string    `json:"content"`
	Likes      Likes     `json:"likes"`
}

var (
	posts = map[int64]*Post{}
	seq   = int64(1)
)

func getAllPosts(c echo.Context) error {
	tmpPosts := map[string]*Post{}
	keys := make([]int64, 0, len(posts))
	for k := range posts {
		keys = append(keys, k)
	}

	for _, i := range keys {
		tmpPosts[strconv.FormatInt(i, 10)] = posts[i]
	}
	return c.JSON(http.StatusOK, tmpPosts)
}

func likePost(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	posts[id].Likes.like(u.ID)

	return c.JSON(http.StatusOK, posts[id])
}

func createPost(c echo.Context) error {
	p := &Post{
		ID:         seq,
		CreateTime: time.Now(),
	}
	if err := c.Bind(p); err != nil {
		return err
	}

	posts[p.ID] = p
	seq++
	return c.JSON(http.StatusCreated, p)
}

func getPost(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
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

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
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
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	delete(posts, id)
	return c.NoContent(http.StatusNoContent)
}
