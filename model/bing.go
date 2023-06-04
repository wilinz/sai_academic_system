package model

type BingImage struct {
	Url     string `json:"url"`
	EndDate string `json:"enddate"`
}

type BingResponse struct {
	Images []BingImage `json:"images"`
}
