// ./models/admin.go
package models

import (
	"gorm.io/gorm"
	"time"
)

// Admin represents an administrative user with specific permissions.
type Admin struct {
	gorm.Model
	UserID      uint      `json:"user_id" gorm:"unique"`
	Permissions string    `json:"permissions"` // Define permission levels as needed
	AssignedAt  time.Time `json:"assigned_at"`
}
