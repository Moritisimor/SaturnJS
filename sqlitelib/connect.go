package sqlitelib

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
)

func Connect(path string) (map[string]any, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return map[string]any{}, err
	}

	return map[string]any{
		"startTransaction": StartTransaction(db),
		"execute": Execute(db),
		"querySingle": QuerySingle(db),
		"query": Query(db),
		"close": db.Close,
	}, nil
}
