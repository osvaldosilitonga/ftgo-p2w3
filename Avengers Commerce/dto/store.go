package dto

type Weather struct {
	CloudPct    any `json:"cloud_pct"`
	Temp        any `json:"temp"`
	FeelsLike   any `json:"feels_like"`
	Humidity    any `json:"humidity"`
	MinTemp     any `json:"min_temp"`
	MaxTemp     any `json:"max_temp"`
	WindSpeed   any `json:"wind_speed"`
	WindDegrees any `json:"wind_degrees"`
	Sunrise     any `json:"sunrise"`
	Sunset      any `json:"sunset"`
}

type StoreResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

type StoreDetailResponse struct {
	Code    int           `json:"code"`
	Store   StoreResponse `json:"store"`
	Weather Weather       `json:"weather"`
}
