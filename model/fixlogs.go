package model

import "github.com/jinzhu/gorm"

// Fixlogs logs
type Fixlogs struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Names       string `json:"names"`
	Institute   string `json:"institute"`
	Department  string `json:"department"`
	Agent       string `json:"agent"`
	Fixdetail   string `json:"fixdetail"`
	Fixstatus   string `json:"fixstatus"`
}
