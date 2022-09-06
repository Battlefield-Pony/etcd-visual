package model

type BaseEntity struct {
	Id              string `json:"id,omitempty" gorm:"primary"`
	CreatedBy       string `json:"createdBy,omitempty"`
	CreatedTime     int64  `json:"createdTime" gorm:"autoCreateTime"`
	LastUpdatedBy   string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedTime int64  `json:"lastUpdatedTime" gorm:"autoUpdateTime"`
	Version         string `json:"version,omitempty"`
	RouteGroupId    string `json:"routeGroupId,omitempty"`
}
