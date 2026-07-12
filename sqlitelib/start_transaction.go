package sqlitelib

import "database/sql"

func StartTransaction(db *sql.DB) func() (map[string]any, error) {
	return func() (map[string]any, error) {
		trans, err := db.Begin()
		if err != nil {
			return map[string]any{}, err
		}

		return map[string]any{
			"commit": trans.Commit,
			"rollback": trans.Rollback,
			"execute": Execute(trans),
			"query": Query(trans),
			"querySingle": QuerySingle(trans),
			"queryScalar": QueryScalar(trans),
		}, nil
	}
}
