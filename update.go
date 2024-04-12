package dbkit

import (
	"database/sql"
)

func (c *kit) UpdateScript(data UpdateInfo) (sql.Result, error) {
	stmt, err := c.db.Prepare(data.Script)
	if err != nil {
		return nil, err
	}

	args := make([]any, len(data.Args.Columns)+len(data.Args.Where))
	args = append(args, data.Args.Columns...)
	args = append(args, data.Args.Where...)

	return stmt.Exec(args...)
}
