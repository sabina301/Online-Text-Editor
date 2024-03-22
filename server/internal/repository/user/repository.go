package user

import (
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
)

type userRepository struct {
	db *sqlx.DB
	m  sync.RWMutex
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(name string, passwordHash string) (int, error) {
	log.Println("REP")
	r.m.Lock()
	defer r.m.Unlock()

	var id int
	query := "INSERT INTO users (name, password_hash) VALUES ($1, $2) RETURNING id"
	row := r.db.QueryRow(query, name, passwordHash)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *userRepository) Get() {

}

func (r *userRepository) Delete() {

}
