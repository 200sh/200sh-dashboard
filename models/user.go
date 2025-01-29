package models

import (
	"errors"
	"fmt"
	"github.com/200sh/200sh-dashboard/internal/repository"
	"strconv"
	"strings"
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
	Id         int64
	ProviderId string
	Provider   string
	Email      string
	Name       string
	Status     UserStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("user email cannot be empty")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

func (u *User) Activate() {
	u.Status = UserStatusActive
	u.UpdatedAt = time.Now()
}

func (u *User) IsActive() bool {
	return u.Status == UserStatusActive
}

func FromDBUser(dbUser repository.User) *User {
	status, _ := UserStatusFromString(strconv.Itoa(int(dbUser.Status)))
	return &User{
		Id:         dbUser.ID,
		ProviderId: dbUser.ProviderID,
		Provider:   dbUser.Provider,
		Email:      dbUser.Email,
		Name:       dbUser.Name,
		Status:     status,
		CreatedAt:  dbUser.CreatedAt.Time,
		UpdatedAt:  dbUser.UpdatedAt.Time,
	}
}
