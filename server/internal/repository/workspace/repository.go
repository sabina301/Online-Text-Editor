package workspace

import "github.com/jmoiron/sqlx"

type workspaceRepository struct {
	db *sqlx.DB
}
