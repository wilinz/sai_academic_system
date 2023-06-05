package model

import (
	"server_template/base_type"
	"server_template/mytime"
)

type Student struct {
	ID        int64                    `json:"id"`
	Majors    string                   `json:"majors"`
	Grade     int                      `json:"grade"`
	StudentNo string                   `json:"student_no"`
	Name      string                   `json:"name"`
	Username  base_type.NullableString `json:"username"`
	CreatedAt *mytime.CustomTime       `gorm:"type:datetime; not null" json:"created_at"`
	UpdatedAt *mytime.CustomTime       `gorm:"type:datetime; not null" json:"updated_at"`
	DeletedAt *mytime.CustomNullTime   `gorm:"type:datetime; default:null" json:"-"`
}
