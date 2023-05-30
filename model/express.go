package model

import (
	"database/sql/driver"
	"encoding/json"
)

//type ExpressOrder struct {
//	ID               int64       `json:"-" gorm:"id"`
//	AddressID        int64       `gorm:"address_id" json:"addressID,omitempty"`
//	Desc             string      `gorm:"desc" json:"desc,omitempty"`
//	ExpressCompany   string      `gorm:"express_company" json:"expressCompany,omitempty"`
//	ExpressNumber    string      `gorm:"express_number" json:"expressNumber,omitempty"`
//	ExpressAddressID int64       `gorm:"express_address_id" json:"expressAddressID,omitempty"`
//	PictureList      PictureList `gorm:"column:picture_list;type:text" json:"pictureList,omitempty"`
//	PickUpLocation   string      `gorm:"pick_up_location" json:"pickUpLocation,omitempty"`
//	Reward           string      `gorm:"reward" json:"reward,omitempty"`
//	SMTPUsername         string      `json:"-" gorm:"username"`
//}

type PictureList []string

func (p PictureList) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *PictureList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}
