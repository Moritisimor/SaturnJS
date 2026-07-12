package sqlitelib

func Execute(db Queryer) func(cmd string, params ...any) (map[string]any, error) {
	return func(cmd string, params ...any) (map[string]any, error) {
		res, err := db.Exec(cmd, params...)
		if err != nil {
			return map[string]any{}, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return map[string]any{}, err
		}

		rows, err := res.RowsAffected()
		if err != nil {
			return map[string]any{}, err
		}

		return map[string]any{
			"insertRowId":  id,
			"rowsAffected": rows,
		}, nil
	}
}
