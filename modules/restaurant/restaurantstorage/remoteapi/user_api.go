package remoteapi

import (
	"context"
	"errors"
	"fmt"
	"food_delivery/common"
	"log"

	"github.com/go-resty/resty/v2"
)

type userAPI  struct {
	client *resty.Client
	serviceURL string
}

func NewUserAPI(serviceURL string) *userAPI {
	return &userAPI{
		client: resty.New(),
		serviceURL: serviceURL,
	}
}

func(user *userAPI) GetUsers(ctx context.Context, ids []int)([]common.SimpleUser, error) {

	type requestUserParam struct {
		Ids []int `json:"ids"`
	}

	type responseUser struct {
		Data []common.SimpleUser `json:"data"`
	}

	var result responseUser

	resp, err := user.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestUserParam{Ids: ids}).
		SetResult(&result).
		Post(fmt.Sprintf("%s/%s", user.serviceURL, "internal/get-users-by-ids"))

	
	if err != nil {
		log.Println(err)
		return nil, err
	}


	if !resp.IsSuccess() {
		log.Println(resp.RawResponse)
		return nil, errors.New("cannot call api get users")
	}
	
	for i := range result.Data {
		result.Data[i].GetRealId()
	}

	return result.Data, nil
} 