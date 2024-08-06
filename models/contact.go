package models

type Contact struct {
	ContactID    int    `json:"contact_id"`
	MembershipID int    `json:"membership_id"`
	ContactType  string `json:"contact_type"`
	ContactValue string `json:"contact_value"`
	IsActive     bool   `json:"is_active"`
	CreatedDate  string `json:"created_date"`
	CreatedBy    string `json:"created_by"`
	UpdateDate   string `json:"update_date"`
	UpdateBy     string `json:"update_by"`
}
