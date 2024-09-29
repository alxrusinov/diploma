package use

import (
	"errors"

	"github.com/alxrusinov/diploma/internal/client"
	"github.com/alxrusinov/diploma/internal/customerrors"
	"github.com/alxrusinov/diploma/internal/model"
	"github.com/alxrusinov/diploma/internal/store"
)

type UsecaseInst struct {
	store  store.Store
	client *client.Client
}

func (useCase *UsecaseInst) CheckUserExists(user *model.User) (bool, error) {
	ok, err := useCase.store.FindUserByLogin(user)
	return ok, err
}

func (useCase *UsecaseInst) CreateUser(user *model.User) error {
	ok, err := useCase.store.CreateUser(user)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("user was not created")
	}

	return nil
}

func (useCase *UsecaseInst) UpdateUser(token *model.Token) (*model.Token, error) {
	resToken, err := useCase.store.UpdateUser(token)

	return resToken, err
}

func (useCase *UsecaseInst) CheckIsValidUser(user *model.User) (bool, error) {
	found, err := useCase.store.FindUserByLoginPassword(user)

	return found, err
}

func (useCase *UsecaseInst) UploadOrder(order *model.Order, login string) (*model.Order, error) {
	var resOrder *model.Order
	noOrderError := new(customerrors.NoOrderError)
	serverError := new(customerrors.ServerError)

	for {
		resOrder, err := useCase.client.GetOrderInfo(order.Number)

		if err != nil {
			if errors.As(err, &serverError) {
				continue
			}

			if errors.As(err, &noOrderError) {
				return nil, err
			}

			return nil, err
		}

		ok, err := useCase.store.AddOrder(resOrder, login)

		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, errors.New("order was not uploaded")
		}

		break

	}

	return resOrder, nil
}

func (useCase *UsecaseInst) GetOrders(login string) ([]model.OrderResponse, error) {
	orders, err := useCase.store.GetOrders(login)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (useCase *UsecaseInst) GetBalance(login string) (*model.Balance, error) {
	order, err := useCase.client.GetOrderInfo(login)

	if err != nil {
		return nil, err
	}

	balance := &model.Balance{
		Current: float64(order.Accrual),
	}

	return balance, nil
}

func (useCase *UsecaseInst) GetWithdrawls(login string) ([]model.Balance, error) {
	return make([]model.Balance, 0), nil
}

func CreateUsecase(store store.Store) Usecase {
	return &UsecaseInst{store: store, client: new(client.Client)}
}
