package categorymodel

import "food_delivery/common"

const EntityName = "Categories"


type Categories struct {
	common.SQLModel 					`json:",inline"`
	Description		string				`json:"description,omitempty" gorm:"column:description;"`
	Name            string              `json:"name" gorm:"column:name;"`
	Icon            *common.Image       `json:"logo,omitempty" gorm:"column:icon;"`
}

func(categories Categories) TableName() string {
	return "categories"
}

func (data *Categories) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeCategory)
	
	fakeId := common.NewUID(uint32(data.Id), common.DbTypeCategory, 1)
	data.FakeId = &fakeId
}
