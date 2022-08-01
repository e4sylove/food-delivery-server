package uploadprovider

import (
	"context"
	"food_delivery/common"
)

type UploadProvider interface {
	SaveFileUpLoaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
