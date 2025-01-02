package models

import (
	"time"
)

type ListCraftRequest struct {
	Name   string `json:"name"`
	Class  string `json:"class"`
	Status string `json:"status"`
}

type SpaceCraft struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"" json:"name" validate:"required"`
	Class     string     `gorm:"" json:"class" validate:"required"`
	Crew      int        `gorm:"" json:"crew"`
	Image     string     `gorm:"" json:"image"`
	Status    string     `gorm:"" json:"status" validate:"required"`
	Value     int        `gorm:"" json:"value" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	Armaments []Armament `gorm:"foreignKey:CraftID" json:"armaments"`
}

type Armament struct {
	ID       int    `gorm:"primaryKey" json:"-"`
	CraftID  int    `gorm:"foreignKey:ID" json:"-"`
	Title    string `gorm:"" json:"title" validate:"required"`
	Quantity int    `gorm:"" json:"quantity" validate:"required"`
}
