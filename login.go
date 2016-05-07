package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"

	_ "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo/engine/standard"
	"net/http"
	"time"
)

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	//if username == "jon" && password == "shhh!" {
	res, _ := r.Table("Users").GetAllByIndex("username", username).Run(session)

	if !res.IsNil() {
		var pass map[string]interface{}
		err := res.One(&pass)
		if err != nil {
			fmt.Printf("Error scanning database result: %s", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if password == pass["password"] {
			// Create token
			token := jwt.New(jwt.SigningMethodHS256)

			// Set claims
			token.Claims["name"] = "Jon Snow"
			token.Claims["admin"] = true
			token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	name := user.Claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
