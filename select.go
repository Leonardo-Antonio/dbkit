package dbkit

import (
	"database/sql"
)

type kit struct {
	db *sql.DB
}

func newKit(db *sql.DB) *kit {
	return &kit{db: db}
}

func (c *kit) SelectItems(data QueryInfo) ResponseQuery {
	stmt, err := c.db.Prepare(data.Query)
	if err != nil {
		return responseQueryErr(err)
	}

	rows, err := stmt.Query(data.Args...)
	if err != nil {
		return responseQueryErr(err)
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		err := c.mapScan(rows, row)
		if err != nil {
			return responseQueryErr(err)
		}
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		return responseQueryErr(err)
	}

	return responseQuerySuccess("ok", result)
}

func (c *kit) mapScan(r *sql.Rows, dest map[string]interface{}) error {
	// ignore r.started, since we needn't use reflect for anything.
	columns, err := r.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	for i := range values {
		values[i] = new(interface{})
	}

	err = r.Scan(values...)
	if err != nil {
		return err
	}

	for i, column := range columns {
		switch v := (*(values[i].(*interface{}))).(type) {
		case []byte:
			dest[column] = string(v)
		default:
			dest[column] = *(values[i].(*interface{}))
		}
	}

	return r.Err()
}
