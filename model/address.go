package model

type Address struct {
	ID        int64  `gorm:"id" json:"id,omitempty"`
	Category  int    `gorm:"category" json:"category,omitempty"`
	Address   string `gorm:"address" json:"address,omitempty"`
	IsDefault *bool  `gorm:"is_default" json:"isDefault,omitempty"`
	Gender    int    `gorm:"gender" json:"gender,omitempty"`
	Name      string `gorm:"name" json:"name,omitempty"`
	Phone     string `gorm:"phone" json:"phone,omitempty"`
	Username  string `json:"-" gorm:"username"`
}
