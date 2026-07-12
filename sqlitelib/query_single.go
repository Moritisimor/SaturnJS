package sqlitelib

import "fmt"

func QuerySingle(db Queryer) func(sqlText string, params ...any) ([]any, error) {
	return func(sqlText string, params ...any) ([]any, error) {
		rows, err := db.Query(sqlText, params...)
		if err != nil {
			return []any{}, err
		}

		cols, err := rows.Columns()
		if err != nil {
			return []any{}, err
		}

		results := make([]any, len(cols))
		pointers := make([]any, len(cols))
		for idx := range results {
			pointers[idx] = &results[idx]
		}

		scans := 0
		for rows.Next() {
			if err := rows.Scan(pointers...); err != nil {
				return results, err
			}

			scans++
		}

		if scans > 1 {
			return results, fmt.Errorf("query returned more than one result")
		}

		return results, rows.Err()
	}
}
