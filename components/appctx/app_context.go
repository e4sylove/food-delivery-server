package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMySQLConnection() *gorm.DB
	GetSecret() string
}

type appCtx struct {
	db     *gorm.DB
	secret string
}

func NewAppContext(db *gorm.DB, secret string) *appCtx {
	return &appCtx{db: db, secret: secret}
}

func (ctx *appCtx) GetMySQLConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecret() string {
	return ctx.secret
}
