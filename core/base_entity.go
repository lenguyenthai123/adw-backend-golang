package core

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

// BaseEntity struct with common fields for almost all entities
type BaseEntity struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// BeforeCreate is a hook that generates a UUID before creating an entity.
// It ensures that the ID field is populated with a unique identifier.
func (record *BaseEntity) BeforeCreate(*gorm.DB) (err error) {
	// Check id field is empty
	if record.ID == "" {
		// Assign the generated UUID to the ID field
		record.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	}

	return nil
}
