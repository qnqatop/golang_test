package dto

import "time"

type GoodsCreateRequest struct {
	Name string `json:"name"`
}

type GoodsResponse struct {
	ID          int       `json:"id"`
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Removed     bool      `json:"removed"`
	CreateAt    time.Time `json:"create_at"`
}

type GoodsUpdateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GoodsDeleteResponse struct {
	ID      int  `json:"id"`
	Removed bool `json:"removed"`
}
