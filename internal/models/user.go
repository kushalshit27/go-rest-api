package models

import "time"

type Users struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Age       uint       `json:"age,omitempty"`
	Role      string     `json:"role"`
	CreatedAt *time.Time `json:"created"`
	UpdatedAt *time.Time `json:"updated,omitempty"`
	Status    bool       `json:"status"`
}
