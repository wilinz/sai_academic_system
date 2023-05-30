package model

type LocationResult struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Count     int    `json:"count"`
	RequestID string `json:"request_id"`
	Data      []Data `json:"data"`
	Region    Region `json:"region"`
}
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type AdInfo struct {
	Adcode   int    `json:"adcode"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}
type Data struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Address  string         `json:"address"`
	Tel      string         `json:"tel"`
	Category string         `json:"category"`
	Type     int            `json:"type"`
	Location LocationResult `json:"location"`
	Distance int            `json:"_distance"`
	AdInfo   AdInfo         `json:"ad_info"`
}

type Region struct {
	Title string `json:"title"`
}

type GetLocationArguments struct {
	Keyword   string `schema:"keyword"`
	Boundary  string `schema:"boundary"`
	Filter    string `schema:"filter"`
	PageSize  int    `schema:"page_size"`
	PageIndex int    `schema:"page_index"`
	Key       string `schema:"key"`
}
