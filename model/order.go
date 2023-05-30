package model

import (
	"database/sql/driver"
	"encoding/json"
)

type OrderJson struct {
	ID          int64     `json:"-"`
	AddressID   int       `json:"addressId"`
	GenderLimit int       `json:"genderLimit"`
	Time        TimeLimit `json:"time"`
	//TotalReward    int64            `json:"totalReward"`
	//TotalSpend     int64            `json:"totalSpend"`
	//TotalAmount    int64            `json:"totalAmount"`
	ExpressOrders  []*ExpressOrder  `json:"expressOrders"`
	OtherOrders    []*OtherOrder    `json:"otherOrders"`
	PrintOrders    []*PrintOrder    `json:"printOrders"`
	TakeawayOrders []*TakeawayOrder `json:"takeawayOrders"`
}

type Order struct {
	ID        int64  `gorm:"id" json:"-"`                 //
	AddressID int    `gorm:"address_id" json:"addressID"` //
	IsPaid    bool   `gorm:"is_paid" json:"isPaid"`       //
	Username  string `gorm:"username" json:"-"`           //
	OrderInfo
}

type OrderInfo struct {
	Address     string `gorm:"address" json:"address"`
	OrderNumber string `gorm:"column:order_number;unique" json:"orderNumber"`
	GenderLimit int    `gorm:"gender_limit" json:"genderLimit"`
	StartTime   int64  `gorm:"start_time" json:"startTime"`
	EndTime     int64  `gorm:"end_time" json:"endTime"`
	TotalReward int64  `gorm:"total_reward" json:"totalReward"`
	TotalSpend  int64  `gorm:"total_spend" json:"totalSpend"`
	TotalAmount int64  `gorm:"total_amount" json:"totalAmount"`
	OrderTime   int64  `gorm:"order_time" json:"orderTime"`
	SchoolID    int    `json:"schoolID" gorm:"column:school_id;index" json:"schoolID"`
	CampusID    string `json:"campusID" gorm:"column:campus_id;index" json:"campusID"`
}

type TimeLimit struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

type ExpressOrder struct {
	ID               int64       `json:"-"`
	Desc             string      `json:"desc"`
	ExpressAddressID int         `json:"expressAddressId"`
	ExpressCompany   string      `json:"expressCompany"`
	ExpressNumber    string      `json:"expressNumber"`
	Images           FileUrlList `json:"images"`
	PackageSizeLevel int         `json:"packageSizeLevel"`
	PickupCode       string      `json:"pickupCode"`
	PickupLocation   string      `json:"pickupLocation"`
	Reward           int64       `json:"reward"`
	Username         string      `json:"-" gorm:"username"`
	OrderID          int64       `json:"-" gorm:"order_id"`
}
type TakeawayOrder struct {
	ID                    int64  `json:"-"`
	Desc                  string `json:"desc"`
	EstimatedDeliveryTime int64  `json:"estimatedDeliveryTime"`
	PhoneSuffix           string `json:"phoneSuffix"`
	PickupLocation        string `json:"pickupLocation"`
	Reward                int64  `json:"reward"`
	TakeawayType          int    `json:"takeawayType"`
	Username              string `json:"-" gorm:"username"`
	OrderID               int64  `json:"-" gorm:"order_id"`
}

type PrintOrder struct {
	ID       int64       `json:"-"`
	Desc     string      `json:"desc"`
	Files    FileUrlList `json:"files"`
	Reward   int64       `json:"reward"`
	Spend    int64       `json:"spend"`
	Username string      `json:"-" gorm:"username"`
	OrderID  int64       `json:"-" gorm:"order_id"`
}

type OtherOrder struct {
	ID       int64       `json:"-"`
	Desc     string      `json:"desc"`
	Images   FileUrlList `json:"images"`
	Reward   int64       `json:"reward"`
	Spend    int64       `json:"spend"`
	Username string      `json:"-" gorm:"username"`
	OrderID  int64       `json:"-" gorm:"order_id"`
}

type FileUrlList []string

func (p FileUrlList) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *FileUrlList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}
