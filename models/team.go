package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamName    string
	Sport       string
	Members     []User   `gorm:"many2many:team_members;"`
	Leauges     []League `gorm:"many2many:team_league;"`
	HomeMatches []Match  `gorm:"foreignKey:HomeTeamID;references:ID"`
	AwayMatches []Match  `gorm:"foreignKey:AwayTeamID;references:ID"`
}
