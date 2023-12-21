package web

import "time"

type CategoryResponse struct {
	Id int `json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
