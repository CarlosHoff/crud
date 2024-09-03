package shcemas

import (
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
