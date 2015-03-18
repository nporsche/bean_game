package model

type Bean struct {
	Id        uint64  `json:"id"`
	State     uint8   `json:"state"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
