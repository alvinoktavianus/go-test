package common

import "time"

type HealthCheckDTO struct {
	BuildNumber uint16    `json:"build_number"`
	AppVersion  string    `json:"app_version"`
	AppName     string    `json:"app_name"`
	BuildTime   time.Time `json:"build_time"`
}
