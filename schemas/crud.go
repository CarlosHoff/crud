package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Crud struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type CrudResponse struct {
	ID       uint      `json:"id"`
	CreateDT time.Time `json:"createdt"`
	UpdateDT time.Time `json:"updatedt"`
	DeleteDT time.Time `json:"deletedt,omitempty"`
	Role     string    `json:"role"`
	Company  string    `json:"company"`
	Location string    `json:"localtion"`
	Remote   bool      `json:"remote"`
	Link     string    `json:"link"`
	Salary   int64     `json:"salary"`
}
