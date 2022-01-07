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

type UUIDPk struct {
	ID string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
}

type BaseModel struct {
	UUIDPk
	BaseCreatedAt
	BaseUpdatedAt
	BaseDeleteAt
}
