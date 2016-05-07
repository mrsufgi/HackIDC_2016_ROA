package main

import (
	"errors"
	"fmt"

	r "github.com/dancannon/gorethink"
)

// Fetch an element from a given table. Returns a JSONifiable object
// if everything went well. Returns an error otherwise.
func getFromTable(table string, elemId string) (interface{}, error) {
	cur, err := r.DB(dbName).Table(table).Get(elemId).Run(session)
	if err != nil {
		return nil, err
	}

	return getNextFromCursor(cur)
}

// Insert element to table. Returns the newly inserted element, or error if
// something went wrong.
func insertToTable(table string, elem interface{}) (interface{}, error) {
	resp, err := r.DB(dbName).Table(table).Insert(elem).RunWrite(session)
	if err != nil {
		return nil, err
	}

	cur, err := r.DB(dbName).Table(table).Get(resp.GeneratedKeys[0]).Run(session)
	if err != nil {
		return nil, err
	}

	return getNextFromCursor(cur)
}

func removeFromTable(table string, elemId string) (i interface{}, err error) {
	cur, err := r.DB(dbName).Table(table).Get(elemId).Run(session)
	if err != nil {
		return nil, err
	}

	i, err = getNextFromCursor(cur)

	_, er := r.DB(dbName).Table(table).Get(elemId).Delete().RunWrite(session)
	if er != nil {
		return nil, er
	}

	return
}

func updateFieldInTable(table string, elemId string, field map[string]interface{}) (interface{}, error) {
	cur, err := r.DB(dbName).Table(table).Get(elemId).Update(field).Run(session)
	if err != nil {
		return nil, err
	}

	return getDataFromCursor(cur, 1)
}

func getNextFromCursor(c *r.Cursor) (interface{}, error) {
	var ans interface{}
	ok := c.Next(&ans)
	if !ok {
		return nil, errors.New("Error fetching next element from cursor")
	}

	return ans, nil
}

func getDataFromCursor(c *r.Cursor, num int) (interface{}, error) {
	if num < 0 {
		return getAllDataFromCursor(c)
	}

	res := make(chan interface{}, num)
	var next interface{}
	for i := 0; i < num; i++ {
		if !c.Next(&next) {
			break
		}
		res <- next
	}

	close(res)
	ans := make([]interface{}, len(res))
	i := 0
	for elem := range res {
		ans[i] = elem
		i++
	}

	return ans, nil
}

func getAllDataFromCursor(c *r.Cursor) (interface{}, error) {
	var res []interface{}
	err := c.All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func filterFromTable(table string, filter interface{}) (interface{}, error) {
	cur, err := r.DB(dbName).Table(table).Filter(filter).Run(session)
	if err != nil {
		return nil, err
	}

	ans, err := getAllDataFromCursor(cur)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func createTable(table string, indices []string) error {
	res, err := r.DB(dbName).TableCreate(table).RunWrite(session)
	if err != nil {
		return err
	}

	if res.TablesCreated == 1 {
		for _, str := range indices {
			res, err = r.DB(dbName).Table(table).IndexCreate(str).RunWrite(session)
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
	}

	return nil
}
