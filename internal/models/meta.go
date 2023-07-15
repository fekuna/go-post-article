package models

type Meta struct {
	TotalCount int  `json:"total_count"`
	TotalPages int  `json:"total_pages"`
	Page       int  `json:"page"`
	Size       int  `json:"size"`
	HasMore    bool `json:"has_more"`
}
