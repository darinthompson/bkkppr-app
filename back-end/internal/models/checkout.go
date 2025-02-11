package models

import "time"

type Checkout struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	BookID uint `gorm:"not null" json:"book_id"`
	Book   Book `gorm:"foreignKey:BookID" json:"book"`

	UserID *uint `gorm:"null" json:"user_id,omitempty"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	// I HAVE BORROWER HERE BECAUSE THEY MIGHT NOT HAVE AN APP ACCOUNT
	BorrowerFirstName *string `gorm:"null" json:"borrower_first_name,omitempty"`
	BorrowerLastName  *string `gorm:"null" json:"borrower_last_name,omitempty"`
	BorrowerContact   *string `gorm:"null" json:"borrower_contact,omitempty"`

	CheckedOutAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"checked_out_at"`
	DueDate      *time.Time `json:"due_date,omitempty"`
	ReturnedAt   *time.Time `gorm:"null" json:"returned_at,omitempty"`
}
