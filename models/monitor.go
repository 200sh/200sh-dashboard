package models

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/internal/repository"
	util "github.com/200sh/200sh-dashboard/internal/util"
	"net/http"
	"net/url"
	"strings"
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
		CreatedAt: dbMonitor.CreatedAt.Time,
		UpdatedAt: dbMonitor.UpdatedAt.Time,
	}
}

const (
	HttpIntervalDefault            = 5 * time.Minute
	HttpRetriesDefault             = 2
	HttpTimeoutDefault             = 48 * time.Second
	HttpExpectedStatusCodesDefault = "200-299"
	HttpMethodDefault              = http.MethodGet
)

type HttpMonitor struct {
	MonitorID           int
	Url                 string
	Interval            time.Duration
	Retries             int
	Timeout             time.Duration
	ExpectedStatusCodes []string
	HttpMethod          string
	HttpBody            *string
	HttpHeaders         *string
}

func FromDBHttpMonitor(dbHttpMonitor repository.HttpMonitor) *HttpMonitor {
	return &HttpMonitor{
		MonitorID:           int(dbHttpMonitor.MonitorID),
		Url:                 dbHttpMonitor.Url,
		Interval:            time.Duration(dbHttpMonitor.IntervalS) * time.Second,
		Retries:             int(dbHttpMonitor.Retries),
		Timeout:             time.Duration(dbHttpMonitor.TimeoutS) * time.Second,
		ExpectedStatusCodes: fromDBStringList(dbHttpMonitor.ExpectedStatusCodes),
		HttpMethod:          dbHttpMonitor.HttpMethod,
		HttpBody:            util.NullStringToPtr(dbHttpMonitor.HttpBody),
		HttpHeaders:         util.NullStringToPtr(dbHttpMonitor.HttpHeaders),
	}
}

func fromDBStringList(sl string) []string {
	//  by comma first
	parts := strings.Split(sl, ",")

	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if strings.Contains(p, "-") {
			// Split on '-' and generate the status codes between left and right value
			// return error if there is more than 2 values after the split.

		}

		out = append(out, p)
	}

	return out
}

// func toDBStringList(s []string) string {}
