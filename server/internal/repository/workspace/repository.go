package workspace

import "github.com/jmoiron/sqlx"

type WorkspaceRepository interface {
	Create()
	Get()
	Delete()
	AddUser()
}

type workspaceRepository struct {
	db *sqlx.DB
}
