package userservice

import "context"

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) error
}