package service

import (
	"context"
	"gostore/entity"
	saError "gostore/helper/domain/errorModel"
	"gostore/middleware"
	"gostore/repo"
)

type shippingAddressService struct {
	Repo repo.ShippingAddressRepo
}

type ShippingAddressService interface {
	GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error)
	InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error
	UpdateShippingAddress(ctx context.Context, id int, sa *entity.ShippingAddress) error
	DeleteShippingAddress(ctx context.Context, id int) error
}

func NewShippingAddressService(repo repo.ShippingAddressRepo) ShippingAddressService {
	return &shippingAddressService{
		Repo: repo,
	}
}

func (u *shippingAddressService) GetShippingAddressByUserId(ctx context.Context) ([]entity.ShippingAddress, error) {
	return u.Repo.GetShippingAddressByUserId(ctx)
}

func (u *shippingAddressService) InsertShippingAddress(ctx context.Context, sa *entity.ShippingAddress) error {
	var (
		detailError = make(map[string]any)
	)

	if sa.RecipientName == "" {
		detailError["recipient_name"] = "this field is missing input"
	}

	if sa.Address == "" {
		detailError["address"] = "this field is missing input"
	}

	if sa.Phonenumber == "" {
		detailError["phonenumber"] = "this field is missing input"
	}

	if len(detailError) > 0 {
		return saError.ErrAddressInput.AttachDetail(detailError)
	}

	return u.Repo.InsertShippingAddress(ctx, sa)
}

func (u *shippingAddressService) UpdateShippingAddress(ctx context.Context, id int, sa *entity.ShippingAddress) error {
	var (
		userId      = ctx.Value(middleware.GOSTORE_USERID).(int)
		detailError = make(map[string]any)
	)

	if sa.RecipientName == "" {
		detailError["recipient_name"] = "this field is missing input"
	}

	if sa.Address == "" {
		detailError["address"] = "this field is missing input"
	}

	if sa.Phonenumber == "" {
		detailError["phonenumber"] = "this field is missing input"
	}

	if len(detailError) > 0 {
		return saError.ErrAddressInput.AttachDetail(detailError)
	}

	existSA, err := u.Repo.GetShippingAddressById(ctx, id)
	if err != nil {
		return err
	}

	// validation user updater is right user data
	if existSA.UserId != userId {
		return saError.ErrInvalidUser
	}
	err = u.Repo.UpdateShippingAddress(ctx, id, &existSA, sa)
	return err
}

func (u *shippingAddressService) DeleteShippingAddress(ctx context.Context, id int) error {
	userId := ctx.Value(middleware.GOSTORE_USERID).(int)
	existSA, err := u.Repo.GetShippingAddressById(ctx, id)
	if err != nil {
		return err
	}

	if existSA.UserId != userId {
		return saError.ErrInvalidUser
	}

	return u.Repo.DeleteShippingAddress(ctx, id)
}
