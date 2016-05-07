package main

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
	"github.com/mewben/config-echo"
	//_ "github.com/lib/pq"
	//"github.com/mewben/db-go-env"
	_ "fmt"
	"log"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo/engine/standard"
	_ "golang.org/x/net/websocket"
)

var (
	createTime  = "createTime"
	editTime    = "editTime"
	creatorId   = "creatorId"
	creatorName = "creatorName"
	title       = "title"
	imageUrl    = "imageUrl"
	postId      = "postId"
	count       = "count"
	id          = "id"
	userId      = "userId"
	commentId   = "commentId"
	username    = "userName"
	password    = "password"
	content     = "content"
)

// Initialize Port and DB Connection config
func init() {
	type Config struct {
		SERVERPORT string
		//DB   db.Database
	}

	configFile, err := os.Open("env.json")
	if err != nil {
		panic(err)
	}

	var devConfig Config
	jsonParser := json.NewDecoder(configFile)

	if err = jsonParser.Decode(&devConfig); err != nil {
		panic(err)
	}

	// setup postgres db connection
	//db.Setup(devConfig.DB)

	// setup port
	// This sets the global Port string
	// If you set an environment variable DATABASE_URL,
	// it sets Mode = "prod" and uses the env Port instead
	config.Setup(devConfig.SERVERPORT)
}

var session *r.Session
var dbName string = "RoastMe"

func createTables() {
	createLikesTable()
	createCommentsTable()
	createPostsTable()
	createUsersTable()
}

func main() {
	se, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: dbName,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	session = se

	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
	app.Use(middleware.Gzip())
	app.Use(middleware.Static("public"))

	// Users Routes
	app.POST("/users", createUser)
	app.GET("/users/:id", getUser)
	app.PUT("/users/:id", updateUser)
	app.DELETE("/users/:id", deleteUser)

	// Comments Routes
	app.GET("/comment/:id/likes", getCommentLikes)
	app.GET("/comment/:id", getComment)
	app.GET("/comment/all", getAllComments)
	app.GET("/comment/head/:count", getLastComments)
	app.GET("/comment/top_rated/:count", getTopComments)
	app.POST("/comment/create", createComment)
	app.POST("/comment/delete", deleteComment)
	app.POST("/comment/like", likeComment)
	app.POST("/comment/edit", editComment)

	//app.GET("/feed", standard.WrapHandler(websocket.Handler(feedHandler)))
	// Login route
	app.POST("/login", login)

	// Unauthenticated route
	app.GET("/", accessible)

	// Restricted group
	b := app.Group("/restricted")
	b.Use(middleware.JWTAuth([]byte("secret")))
	b.GET("", restricted)

	createTables()

	app.Run(standard.New(config.Port))
}
