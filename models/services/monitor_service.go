package services

import (
	"context"
	"database/sql"
	"github.com/200sh/200sh-dashboard/internal/repository"
	"github.com/200sh/200sh-dashboard/models"
)

type MonitorService interface {
	Create(m *models.Monitor) error
	ListByUser(userID int64) ([]*models.Monitor, error)
	GetByIDAndUser(id int, userID int64) (*models.Monitor, error)
	Delete(monitorID int, userID int64) error
}

type monitorService struct {
	db   *sql.DB // For transactions
	repo *repository.Queries
}

func NewMonitorService(db *sql.DB, repo *repository.Queries) MonitorService {
	return &monitorService{db: db, repo: repo}
}

func (s *monitorService) GetByIDAndUser(id int, userID int64) (*models.Monitor, error) {
	p := repository.GetMonitorByUserIDAndMonitorIDParams{
		UserID: userID,
		ID:     int64(id),
	}
	dbMonitor, err := s.repo.GetMonitorByUserIDAndMonitorID(context.Background(), p)
	if err != nil {
		return nil, err
	}

	return models.FromDBMonitor(dbMonitor), nil
}

func (s *monitorService) ListByUser(userID int64) ([]*models.Monitor, error) {
	dbMonitors, err := s.repo.GetMonitorsByUserID(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	monitors := make([]*models.Monitor, len(dbMonitors))
	for _, m := range dbMonitors {
		monitors = append(monitors, models.FromDBMonitor(m))
	}

	return monitors, nil
}

func (s *monitorService) Create(m *models.Monitor) error {
	if err := m.Validate(); err != nil {
		return err
	}

	params := repository.CreateMonitorParams{
		UserID: m.UserId,
		Url:    m.Url,
	}

	dbMonitor, err := s.repo.CreateMonitor(context.Background(), params)
	if err != nil {
		return err
	}

	// Update the input monitor with the database values
	*m = *models.FromDBMonitor(dbMonitor)
	return nil
}

func (s *monitorService) Delete(monitorID int, userID int64) error {
	params := repository.DeleteMonitorParams{
		ID:     int64(monitorID),
		UserID: userID,
	}
	return s.repo.DeleteMonitor(context.Background(), params)
}
