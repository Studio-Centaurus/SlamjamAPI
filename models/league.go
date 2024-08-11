package models

import "gorm.io/gorm"

type League struct {
	gorm.Model
	LeagueName string
	Sport      string
	Teams      []Team `gorm:"many2many:team_league;"`
}
