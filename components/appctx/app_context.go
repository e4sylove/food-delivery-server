package appctx

import (
	"food_delivery/components/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMySQLConnection() *gorm.DB
	GetSecret() string
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db       *gorm.DB
	secret   string
	provider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, secret string, provider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, secret: secret, provider: provider}
}

func (ctx *appCtx) GetMySQLConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetSecret() string {
	return ctx.secret
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.provider
}
