package model

import "time"

type BaseEntity struct {
	Id              string    `json:"id,omitempty" gorm:"primary"`
	CreatedBy       string    `json:"created_by,omitempty"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by,omitempty"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
	Version         string    `json:"version,omitempty"`
	RouteGroupId    string    `json:"route_group_id,omitempty"`
}
