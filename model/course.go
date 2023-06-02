package model

import "server_template/mytime"

type Course struct {
	ID         int64                  `json:"id"`
	CourseName string                 `json:"course_name"`
	CourseNo   string                 `json:"course_no"`
	Credits    float64                `json:"credits"`
	Teacher    string                 `json:"teacher"`
	Grade      int                    `json:"grade"`
	Room       string                 `json:"room"`
	Capacity   int                    `json:"capacity"`
	Selected   int                    `json:"selected"`
	CreatedAt  *mytime.CustomTime     `gorm:"type:datetime; not null" json:"created_at"`
	UpdatedAt  *mytime.CustomTime     `gorm:"type:datetime; not null" json:"updated_at"`
	DeletedAt  *mytime.CustomNullTime `gorm:"type:datetime; default:null" json:"-"`
}
