package models

import (
	"errors"
	"time"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleEmployee Role = "employee"
)

type User struct {
	NIP          string    `gorm:"type:char(36);primaryKey"`
	FullName     string    `gorm:"type:varchar(255);not null"`
	Username     string    `gorm:"size:255;not null"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         Role      `gorm:"size:50;default:'user'"`
	Dob          time.Time `gorm:"not null" json:"dob"`
}

func ValidateRole(role Role) error {
	switch role {
	case RoleAdmin, RoleEmployee:
		return nil
	default:
		return errors.New("invalid role")
	}
}
