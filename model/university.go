package model

type University struct {
	Number             int    `gorm:"number" json:"-"`
	Name               string `gorm:"name" json:"name"`
	ID                 int64  `gorm:"id" json:"id"`
	CompetentAuthority string `gorm:"competent_authority" json:"-"`
	Location           string `gorm:"location" json:"-"`
	SchoolLevel        string `gorm:"school_level" json:"-"`
	Remark             string `gorm:"remark" json:"-"`
}

type UniversityResponse struct {
	Name string `gorm:"name" json:"name"`
	ID   int64  `gorm:"id" json:"id"`
}

type UniversityResponse2 struct {
	SchoolID   int64  `json:"id" gorm:"school_id"`
	SchoolName string `json:"name" gorm:"school_name"`
}

type Campus struct {
	Number      int64   `gorm:"column:number;PRIMARY_KEY;AUTO_INCREMENT"`
	ID          string  `gorm:"column:id;unique"`
	SchoolID    int64   `json:"school_id"`
	SchoolName  string  `json:"school_name"`
	Name        string  `json:"name"`
	FullName    string  `json:"full_name"`
	Address     string  `json:"address"`
	Tel         string  `json:"tel"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	AddressCode int     `json:"address_code"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	District    string  `json:"district"`
}

type CampusResponse struct {
	ID         string `json:"ID"`
	SchoolID   int64  `json:"schoolID"`
	SchoolName string `json:"schoolName"`
	Name       string `json:"name"`
	FullName   string `json:"fullName"`
	Address    string `json:"address"`
}
