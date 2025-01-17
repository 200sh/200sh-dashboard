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
	UpdatedAt  time.Time
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{db: db}
}

var NoUserFound = errors.New("user: no user found")

func (s *UserService) GetByProviderId(id string) (*User, error) {
	// Try to fetch the user from the db given a provider id
	row := s.db.QueryRow(
		`SELECT id, provider_id, provider, name, email, status, created_at, updated_at FROM user WHERE provider_id = ?`,
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
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NoUserFound
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserService) CreateUser(user *User) error {
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

func (s *UserService) CreateMonitor(monitor *Monitor) error {
	stmt, err := s.db.Prepare(`
	INSERT INTO monitor (user_id, url) VALUES (?, ?)
`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(monitor.UserId, monitor.Url)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetMonitors(user *User) ([]Monitor, error) {
	stmt, err := s.db.Prepare(`
	SELECT id, user_id, url, created_at, updated_at FROM monitor WHERE user_id = ?
`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	monitors := make([]Monitor, 0)

	for rows.Next() {
		var monitor Monitor
		err = rows.Scan(
			&monitor.Id,
			&monitor.UserId,
			&monitor.Url,
			&monitor.CreatedAt,
			&monitor.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		monitors = append(monitors, monitor)
	}

	return monitors, nil
}

// GetMonitor Fetches the monitor by the given id if it exists for that user
func (s *UserService) GetMonitor(id int, user *User) (*Monitor, error) {
	stmt, err := s.db.Prepare(`
	SELECT id, user_id, url, created_at, updated_at FROM monitor WHERE id = ? AND user_id = ?
`)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id, user.Id)

	var monitor Monitor

	err = row.Scan(
		&monitor.Id,
		&monitor.UserId,
		&monitor.Url,
		&monitor.CreatedAt,
		&monitor.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &monitor, nil
}
