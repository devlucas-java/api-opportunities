package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Remote   bool
	Salary   int
	Location string
	Link     string
}
