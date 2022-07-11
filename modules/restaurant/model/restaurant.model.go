package restaurantmodel

import "time"


type Restaurant struct {
	Id int `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Address string `json:"address"`
	Latitude float32 `json:"lat"`
	Longitude float32 `json:"lng"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}