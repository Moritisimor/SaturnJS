package sqlitelib

import "fmt"

func QueryScalar(db Queryer) func(cmd string, params ...any) (any, error) {
	return func(cmd string, params ...any) (any, error) {
		rows, err := db.Query(cmd, params...)
		if err != nil {
			return nil, err
		}

		cols, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		if len(cols) > 1 {
			return nil, fmt.Errorf("query returned more than one column (not a scalar)")
		}

		scans := 0
		var result any

		for rows.Next() {
			rows.Scan(&result)

			scans++
		}

		if scans > 1 {
			return nil, fmt.Errorf("query returned more than one result")
		}

		return result, nil
	}
}
