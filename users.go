package main

import (
	_ "fmt"
	"log"
	"net/http"
	_ "strconv"

	r "github.com/dancannon/gorethink"
	"github.com/labstack/echo"
)

var usersTable string = "users"

type user struct {
	userId   string
	password string
	username string
}

func createUsersTable() error {
	indices := []string{
		userId,
		password,
		username,
	}
	return createTable(usersTable, indices)
}

func isUserUnique(userNameInput string) bool {
	res, err := r.Table(usersTable).GetAllByIndex(userNameInput, username).Run(session)
	if err != nil {
		//fmt.Print(err)
		return false
	}

	if res == nil {
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

	if isUserUnique(u.username) {
		// db
		r.Table(usersTable).Insert(u).RunWrite(session)

		return c.JSON(http.StatusCreated, u)
	}

	return echo.NewHTTPError(http.StatusConflict)
}

func getUserByProp(index string, value string) (*r.Cursor, error) {
	res, err := r.Table(usersTable).GetAllByIndex(index, value).Run(session)
	return res, err
}

func getUser(c echo.Context) error {
	idLocal := c.Param(id)

	res, err := getUserByProp(id, idLocal)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	if res.IsNil() {
		log.Println(idLocal + " not found")
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

func getUsernameByID(idInput string) string {
	res, err := getUserByProp(id, idInput)
	if err != nil {
		log.Println(err)
		return ""
	}

	if res.IsNil() {
		log.Println(idInput + " not found")
		return ""
	}

	var result map[string]interface{}
	err = res.One(&result)
	if err != nil {
		log.Printf("Error scanning database result: %s", err)
		return ""
	}
	return result[username].(string)
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	idLocal := c.Param(id)
	res, err := r.Table(usersTable).Get(idLocal).Update(u).RunWrite(session)
	if err != nil {
		log.Printf("Error updating database: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, res)
}

func deleteUser(c echo.Context) error {
	idLocal := c.Param(id)
	_, err := r.Table(usersTable).Get(idLocal).Delete().RunWrite(session)
	if err != nil {
		log.Print("Error deleting entry: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusNoContent)
}
