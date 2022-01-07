package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseCreatedAt struct {
	CreatedAt time.Time
}

type BaseUpdatedAt struct {
	UpdatedAt time.Time
}

type BaseDeleteAt struct {
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Pk struct {
	ID uint `gorm:"primarykey"`
}

type BaseModel struct {
	Pk
	BaseCreatedAt
	BaseUpdatedAt
	BaseDeleteAt
}
