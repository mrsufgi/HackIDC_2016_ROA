package main

import (
	r "github.com/dancannon/gorethink"

	"net/http"
	_ "strconv"

	_ "fmt"
	"log"

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
		//fmt.Print(err)
		return false
	}

	if res.IsNil() {
		//fmt.Print("Entry wasn't found")
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

func getUserByProp(index string, value string) (*r.Cursor, error) {
	res, err := r.Table("Users").GetAllByIndex(index, value).Run(session)
	return res, err
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	res, err := getUserByProp("id", id)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if res.IsNil() {
		log.Println(id + "not found")
		return echo.NewHTTPError(http.StatusNotFound)
	}

	var result map[string]interface{}
	err = res.One(&result)
	if err != nil {
		log.Printf("Error scanning database result: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, result)
}

func getUsernameByID(id string) string {
	res, err := getUserByProp("id", id)
	if err != nil {
		log.Println(err)
		return ""
	}

	if res.IsNil() {
		log.Println(id + "not found")
		return ""
	}

	var result map[string]interface{}
	err = res.One(&result)
	if err != nil {
		log.Printf("Error scanning database result: %s", err)
		return ""
	}
	return result["username"].(string)
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id := c.Param("id")
	res, err := r.Table("Users").Get(id).Update(u).RunWrite(session)
	if err != nil {
		log.Printf("Error updating database: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, res)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	_, err := r.Table("Users").Get(id).Delete().RunWrite(session)
	if err != nil {
		log.Print("Error deleting entry: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
