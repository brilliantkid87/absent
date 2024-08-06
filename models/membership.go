package models

import "time"

type Membership struct {
	MembershipID int       `json:"membership_id"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	Address      string    `json:"address"`
	IsActive     bool      `json:"is_active"`
	CreatedBy    string    `json:"created_by"`
	CreatedDate  time.Time `json:"created_date"`
	UpdateDate   time.Time `json:"update_date"`
	UpdateBy     string    `json:"update_by"`
}
