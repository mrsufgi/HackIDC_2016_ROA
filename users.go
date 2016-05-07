package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"

	"net/http"
	_ "strconv"

	"github.com/labstack/echo"
)

type (
	user struct {
		ID       string `gorethink:"id,omitempty"`
		Password string `gorethink:"password"`
		Username string `gorethink:"username"`
	}
)

var (
	users = map[string]*user{}
)

func userUniq(username string) bool {
	res, err := r.Table("Users").GetAllByIndex("username", username).Run(session)
	if err != nil {
		fmt.Print(err)
		return false
	}

	if res.IsNil() {
		fmt.Print("Entry wasn't found")
		return true
	}
	defer res.Close()
	return false
}

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if userUniq(u.Username) {

		// db
		r.Table("Users").Insert(u).RunWrite(session)

		return c.JSON(http.StatusCreated, u)
	}

	return echo.NewHTTPError(http.StatusConflict)
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id := c.Param("id")
	users[id].Username = u.Username
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
