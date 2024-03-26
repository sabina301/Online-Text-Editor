package user

import (
	"Online-Text-Editor/server/internal/model"
	"errors"
	"github.com/jmoiron/sqlx"
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

func (r *userRepository) Get(id int) (*model.UserInfo, error) {
	r.m.Lock()
	defer r.m.Unlock()

	var name string
	query := "SELECT name FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&name)
	if err != nil {
		return nil, err
	}
	var user model.UserInfo
	user.Name = name
	return &user, nil
}

func (r *userRepository) GetByUsername(username string, passwordHash string) (*model.UserEntity, error) {
	r.m.Lock()
	defer r.m.Unlock()

	query := "SELECT id FROM users WHERE name = $1 AND password_hash=$2"
	var user model.UserEntity
	err := r.db.Get(&user, query, username, passwordHash)
	if err != nil {
		return nil, errors.New("Cant get user")
	}
	user.Name = username
	user.Password = passwordHash
	return &user, nil
}
