package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Remote   bool
	Salary   int
	Location string
	Link     string
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt.omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Remote    *bool     `json:"remote"`
	Salary    int       `json:"salary"`
	Location  string    `json:"location"`
	Link      string    `json:"link"`
}
