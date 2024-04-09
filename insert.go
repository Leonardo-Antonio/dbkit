package dbkit

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

func (c *kit) getInfoToStruct(data any, returnColumns, returnArgs bool) (columns []string, args []any) {
	values := reflect.ValueOf(data)
	meta := reflect.TypeOf(data)
	for i := 0; i < values.NumField(); i++ {
		tag := meta.Field(i).Tag.Get("db")
		field := values.Field(i)

		// valida si se validaran las columas, en caso si, se valida que tenga info en los tags
		if returnColumns && tag == "-" || len(tag) == 0 {
			continue
		}

		// valida si se retornara los args
		if returnArgs && field.IsZero() {
			continue
		}

		if returnColumns {
			columns = append(columns, tag)
		}

		if returnArgs {
			args = append(args, field.Interface())
		}
	}

	return
}

func (c *kit) NewItem(table string, data any) (sql.Result, error) {
	columns, args := c.getInfoToStruct(data, true, true)
	insert := fmt.Sprintf(`
		Insert Ignore Into %s (
			%s
		) Values (
			%s
		)
	`, table, strings.Join(columns, ", "), strings.TrimSuffix(strings.Repeat("?, ", len(args)), ", "))

	stmt, err := c.db.Prepare(insert)
	if err != nil {
		return nil, err
	}

	return stmt.Exec(args...)
}
