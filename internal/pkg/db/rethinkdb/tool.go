package rethinkdb

import "gopkg.in/rethinkdb/rethinkdb-go.v6"

func (r *Store) getDatabases() ([]string, error) {
	c, err := rethinkdb.DBList().CoerceTo("array").Run(r.client)
	if err != nil {
		return nil, err
	}

	list, err := getArrayString(c)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *Store) getTables() ([]string, error) {
	c, err := rethinkdb.DB("shortlink").TableList().CoerceTo("array").Run(r.client)
	if err != nil {
		return nil, err
	}

	list, err := getArrayString(c)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func getArrayString(c *rethinkdb.Cursor) ([]string, error) {
	z, err := c.Interface()
	if err != nil {
		return nil, err
	}

	var tbsI = z.([]interface{})
	var tbs = make([]string, len(tbsI))

	for i, v := range tbsI {
		tbs[i] = v.(string)
	}

	return tbs, nil
}
