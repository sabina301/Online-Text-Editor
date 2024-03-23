package workspace

import (
	"github.com/jmoiron/sqlx"
)

type workspaceRepository struct {
	db *sqlx.DB
}

func NewWorkspaceRepository(db *sqlx.DB) *workspaceRepository {
	return &workspaceRepository{db: db}
}

func (r *workspaceRepository) Create() {

}

func (r *workspaceRepository) Get() {

}

func (r *workspaceRepository) Delete() {

}

func (r *workspaceRepository) AddUser() {

}
