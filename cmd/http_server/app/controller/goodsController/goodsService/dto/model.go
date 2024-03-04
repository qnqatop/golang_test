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
	ID        int  `json:"id"`
	ProjectId int  `json:"project_id"`
	Removed   bool `json:"removed"`
}

var GoodsResponseList struct {
	Meta struct {
		Total   int `json:"total"`
		Removed int `json:"removed"`
		Limit   int `json:"limit"`
		Offset  int `json:"offset"`
	} `json:"meta"`
	Goods []GoodsResponse `json:"goods"`
}
