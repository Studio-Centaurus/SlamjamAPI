package models

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	HomeTeamID uint
	HomeTeam   Team `gorm:"foreignKey:HomeTeamID;references:ID"`
	AwayTeamID uint
	AwayTeam   Team `gorm:"foreignKey:AwayTeamID;references:ID"`
	HomeScore  int
	AwayScore  int
}
