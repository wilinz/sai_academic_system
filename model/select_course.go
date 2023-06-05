package model

import "server_template/mytime"

type CourseSelection struct {
	ID        int64                  `json:"id"`
	CourseID  int64                  `json:"course_id"`
	StudentID int64                  `json:"student_id"`
	CreatedAt *mytime.CustomTime     `gorm:"type:datetime; not null" json:"created_at"`
	UpdatedAt *mytime.CustomTime     `gorm:"type:datetime; not null" json:"updated_at"`
	DeletedAt *mytime.CustomNullTime `gorm:"type:datetime; default:null" json:"-"`
	Course    Course                 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Student   Student                `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type CourseSelectionInfo struct {
	Course  `json:"course"`
	Student `json:"student"`
}
