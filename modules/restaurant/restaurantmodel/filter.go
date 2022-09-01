package restaurantmodel


type Filter struct {
	CityID int `json:"city_id,omitempty" form:"city_id"`
	OwnerId int `json:"owner_id,omitempty" form:"owner_id"`
}