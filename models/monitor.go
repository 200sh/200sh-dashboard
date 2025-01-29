package models

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/internal/repository"
	"net/url"
	"time"
)

type Monitor struct {
	Id        int
	UserId    int64
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Monitor) Validate() error {
	if _, err := url.Parse(m.Url); err != nil {
		return fmt.Errorf("invalid monitor URL: %w", err)
	}
	return nil
}

func (m *Monitor) UpdateUrl(newUrl string) error {
	parsed, err := url.Parse(newUrl)
	if err != nil {
		return err
	}
	m.Url = parsed.String()
	m.UpdatedAt = time.Now()
	return nil
}

func FromDBMonitor(dbMonitor repository.Monitor) *Monitor {
	return &Monitor{
		Id:        int(dbMonitor.ID),
		UserId:    dbMonitor.UserID,
		Url:       dbMonitor.Url,
		CreatedAt: dbMonitor.CreatedAt.Time,
		UpdatedAt: dbMonitor.UpdatedAt.Time,
	}
}
