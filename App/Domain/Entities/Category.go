package entities

import "time"

type Category struct {
	Id        int       `json:id gorm: "autoIncrement"`
	Name      string    `json:name`
	Slug      string    `json:slug`
	Status    string    `json:status`
	CreatedAt time.Time `json: created_at`
	UpdatedAt time.Time `json: updated_at`
}
