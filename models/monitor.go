package models

import (
	"net/url"
	"time"
)

type Monitor struct {
	Id        int
	UserId    int
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(url url.URL, userId int) Monitor {
	t := time.Now()
	return Monitor{
		Id:        0, // Leave this empty for now as this will be created by the database later
		UserId:    userId,
		Url:       url.String(),
		CreatedAt: t,
		UpdatedAt: t,
	}
}

func (u *User) NewMonitor(url *url.URL) Monitor {
	t := time.Now()
	return Monitor{
		Id:        0, // Leave empty for now. Will be created by database
		UserId:    u.Id,
		Url:       url.String(),
		CreatedAt: t,
		UpdatedAt: t,
	}
}
