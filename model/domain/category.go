package domain

import "time"

type Category struct {
	Id        int `json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
