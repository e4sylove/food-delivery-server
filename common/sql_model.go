package common

import (
	"fmt"
	"time"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1;"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

func (sqlModel *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(sqlModel.Id), int(dbType), 1)
	sqlModel.FakeId = &uid

	fmt.Println(sqlModel.FakeId)
	fmt.Println(sqlModel.FakeId.GetLocalID())

}

func (sqlModel *SQLModel) GetRealId() {
	if sqlModel.FakeId == nil {
		return
	}

	sqlModel.Id = int(sqlModel.FakeId.GetLocalID())
}
