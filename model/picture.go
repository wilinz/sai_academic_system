package model

type Picture struct {
	ID         int64  `gorm:"id"`
	Url        string `gorm:"url"`
	UploadUser string `gorm:"upload_user"`
	ViewUser   string `gorm:"view_user"`
}

type UploadParameters struct {
	Filename string `form:"filename"`
}
