package appctx

import (
	"food_delivery/components/uploadprovider"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMySQLConnection() *gorm.DB
	GetSecret() string
	UploadProvider() uploadprovider.UploadProvider
	GetRedisConnection() *redis.Client
}

type appCtx struct {
	db       *gorm.DB
	secret   string
	provider uploadprovider.UploadProvider
	redis    *redis.Client
}

func NewAppContext(db *gorm.DB, secret string, provider uploadprovider.UploadProvider, redis *redis.Client) *appCtx {
	return &appCtx{
		db: db, 
		secret: secret, 
		provider: provider,
		redis: redis,
	}
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

func (ctx *appCtx) GetRedisConnection() *redis.Client {
	return ctx.redis
}
