package userservice

import (
	"context"
	"food_delivery/common"
	"food_delivery/components/tokenprovider"
	"food_delivery/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginService struct {
	store         LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginService(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginService {
	return &loginService{
		store:         storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. Find user, email
// 2. Hash password from input and compare with password in db
// 3. Provider: issue JWT Token for client
// 4. Access Token and Refresh Token
// 5. Return tokens.

func (service *loginService) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {

	user, err := service.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passwordHashed := service.hasher.Hash(data.Password + user.Salt)

	if passwordHashed != user.Password {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UID: user.Id,
		URole:   user.Role,
	}

	accessToken, err := service.tokenProvider.Generate(payload, service.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := service.tokenProvider.Generate(payload, service.expiry*2)

	if err != nil {
		return nil, common.ErrInternal(err)
	}
	
	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
