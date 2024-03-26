package workspace

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

type workspaceRepository struct {
	db *sqlx.DB
	m  sync.RWMutex
}

func NewWorkspaceRepository(db *sqlx.DB) *workspaceRepository {
	return &workspaceRepository{db: db}
}

func (r *workspaceRepository) Create(workspaceName string, userId int) (int, error) {
	r.m.Lock()
	defer r.m.Unlock()

	tx, err := r.db.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	var workspaceId int
	query := "INSERT INTO workspace (name) VALUES ($1) RETURNING id"
	row := tx.QueryRow(query, workspaceName)
	err = row.Scan(&workspaceId)
	if err != nil {
		return -1, err
	}

	query = "INSERT INTO user_workspace (user_id, workspace_id) VALUES ($1, $2)"
	_, err = tx.Exec(query, userId, workspaceId)
	if err != nil {
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return workspaceId, nil
}

func (r *workspaceRepository) Get() {

}

func (r *workspaceRepository) Delete() {

}

func (r *workspaceRepository) AddUser(workspaceId int, userId int) error {
	r.m.Lock()
	defer r.m.Unlock()

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := "INSERT INTO user_workspace (user_id, workspace_id) VALUES ($1, $2)"
	_, err = tx.Exec(query, userId, workspaceId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
