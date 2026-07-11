package sqlitelib

import "database/sql"

func Query(db *sql.DB) func(sqlText string, params... any) ([][]any, error) {
	return func(sqlText string, params ...any) ([][]any, error) {
		rows, err := db.Query(sqlText, params...)
		if err != nil {
			return [][]any{}, err
		}

		cols, err := rows.Columns()
		if err != nil {
			return [][]any{}, err
		}

		results := [][]any{}
		for rows.Next() {
			thisRow := make([]any, len(cols))
			thisRowPointers := make([]any, len(cols))
			for i := range thisRow {
				thisRowPointers[i] = &thisRow[i]
			}

			if err := rows.Scan(thisRowPointers...); err != nil {
				return [][]any{}, err
			}

			results = append(results, thisRow)
		}

		return results, rows.Err()
	}
}
