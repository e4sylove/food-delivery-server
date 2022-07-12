package restaurantstorage

import "gorm.io/gorm"

type SQLStorage struct {
	db *gorm.DB
}

func newSQLStore(db *gorm.DB) *SQLStorage {
	return &SQLStorage{
		db: db,
	}
}
