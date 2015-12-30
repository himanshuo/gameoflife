package models

type Player struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	AchievementPoints string `json:"achievementPoints"`
}
