package restaurantmodel

import (
	"errors"
	common2 "food_delivery/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common2.SQLModel `json:",inline"`
	Id               int            `json:"id"`
	Name             string         `json:"name"`
	Addr             string         `json:"address"`
	Logo             *common2.Image `json:"logo" gorm:"column:logo;"`
	// City string `json:"city"`
	// Lat float32 `json:"lat"`
	// Lng float32 `json:"lng"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common2.SQLModel `json:",inline"`
	Id               int            `json:"id" gorm:"column:id;" `
	Name             string         `json:"name" gorm:"column:name;"`
	Addr             string         `json:"address" gorm:"column:addr;"`
	Logo             *common2.Image `json:"logo" gorm:"column:logo;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {

	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}
