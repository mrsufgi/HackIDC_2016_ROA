package main

//
//import (
//	"net/http"
//	"strconv"
//	"time"
//
//	r "github.com/dancannon/gorethink"
//	"github.com/labstack/echo"
//)
//
//type Post struct {
//	ID         int               `json:"post_id"`
//	CreateTime int64             `json:"create_time"`
//	EditTime   int64             `json:"edit_time"`
//	CreatorID  int               `json:"creator_id"`
//	Content    string            `json:"content"`
//	Likes      map[string]string `json:"likes"`
//}
//
//var postsTable string = "posts"
//
////var (
////	posts    = map[int]*Post{}
////	post_seq = 1
////)
//
//func getLastPosts(c echo.Context) error {
//	num, _ := strconv.Atoi(c.Param("num"))
//
//	cur, err := r.DB("test").Table(postsTable).OrderBy(r.OrderByOpts{
//		Index: r.Desc("CreateTime"),
//	}).Run(session)
//
//	if err != nil {
//		return err
//	}
//	res := make(chan interface{}, num)
//	var next interface{}
//	for i := 0; i < num; i++ {
//		if !cur.Next(&next) {
//			break
//		}
//		res <- next
//	}
//
//	close(res)
//	ans := make([]interface{}, len(res))
//	i := 0
//	for elem := range res {
//		ans[i] = elem
//		i++
//	}
//
//	return c.JSON(http.StatusOK, ans)
//}
//
//func getAllPosts(c echo.Context) error {
//	cur, err := r.DB("test").Table(postsTable).OrderBy(r.OrderByOpts{
//		Index: r.Desc("CreateTime"),
//	}).Run(session)
//	var res []interface{}
//	err = cur.All(&res)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, res)
//}
//
//func likePost(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//	u := new(user)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//
//	suid := strconv.Itoa(u.UserID)
//	if _, ok := posts[id].Likes[suid]; !ok {
//		posts[id].Likes[suid] = u.Name
//	} else {
//		delete(posts[id].Likes, suid)
//	}
//
//	return c.JSON(http.StatusOK, posts[id])
//}
//
//func createPost(c echo.Context) error {
//	p := &Post{
//		ID:         post_seq,
//		CreateTime: time.Now().Unix(),
//		EditTime:   time.Now().Unix(),
//	}
//	if err := c.Bind(p); err != nil {
//		return err
//	}
//
//	posts[p.ID] = p
//	post_seq++
//	return c.JSON(http.StatusCreated, p)
//}
//
//func getPost(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//	return c.JSON(http.StatusOK, posts[id])
//}
//
//func editPost(c echo.Context) error {
//	editor := new(user)
//	if err := c.Bind(editor); err != nil {
//		return err
//	}
//	p := new(Post)
//	if err := c.Bind(p); err != nil {
//		return err
//	}
//
//	id, _ := strconv.Atoi(c.Param("id"))
//	if editor.UserID == posts[id].CreatorID {
//		posts[id].Content = p.Content
//		posts[id].EditTime = time.Now()
//		return c.JSON(http.StatusOK, posts[id])
//	} else {
//		return c.JSON(http.StatusUnauthorized, posts[id])
//	}
//}
//
//func deletePost(c echo.Context) error {
//	creator := new(user)
//	if err := c.Bind(creator); err != nil {
//		return err
//	}
//	// TODO check post id same as user id
//	id, _ := strconv.Atoi(c.Param("id"))
//	delete(posts, id)
//	return c.NoContent(http.StatusNoContent)
//}
