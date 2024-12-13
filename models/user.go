package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserStatus int

const (
	UserStatusNotActive = iota
	UserStatusActive
	UserStatusBanned
)

var userStateName = map[UserStatus]string{
	UserStatusNotActive: "not-active",
	UserStatusActive:    "active",
	UserStatusBanned:    "banned",
}

func (us UserStatus) String() string {
	return userStateName[us]
}

func UserStatusFromString(status string) (UserStatus, error) {
	for k, v := range userStateName {
		if v == status {
			return k, nil
		}
	}
	return UserStatusNotActive, fmt.Errorf("status '%s' could not be found in UserStatus", status)
}

type User struct {
	Id         int
	ProviderId string
	Provider   string
	Email      string
	Name       string
	Status     UserStatus
	CreatedAt  time.Time
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{db: db}
}

var NoUserFound = errors.New("user: no user found")

func (s UserService) GetByProviderId(id string) (*User, error) {
	// Try to fetch the user from the db given a provider id
	row := s.db.QueryRow(
		`SELECT id, provider_id, provider, name, email, status, created_at FROM user WHERE provider_id = ?`,
		id,
	)

	var user User

	err := row.Scan(
		&user.Id,
		&user.ProviderId,
		&user.Provider,
		&user.Name,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NoUserFound
		}
		return nil, err
	}

	return &user, nil
}

func (s UserService) CreateUser(user *User) error {
	stmt, err := s.db.Prepare(`
	INSERT INTO user (provider_id, provider, name, email, status, created_at) 
	VALUES (?, ?, ?, ?, ?, ?)
`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.ProviderId, user.Provider, user.Name, user.Email, user.Status, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
