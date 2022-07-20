package components

import "gorm.io/gorm"

type AppContext interface {
	GetMySQLConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{ db: db }
}

func (ctx *appCtx) GetMySQLConnection() *gorm.DB {
	return ctx.db
}