package service

import (
	"context"

	"github.com/MidnightHelix/assignment-2/internal/model"
	"github.com/MidnightHelix/assignment-2/internal/repository"
)

type OrderService interface {
	GetOrders(ctx context.Context) ([]model.Order, error)
	GetOrdersById(ctx context.Context, id uint64) (model.Order, error)
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	UpdateOrder(ctx context.Context, order model.Order, id uint64) (model.Order, error)
	DeleteOrder(ctx context.Context, id uint64) error
}

type orderServiceImpl struct {
	repo repository.OrderQuery
}

func NewOrderService(repo repository.OrderQuery) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (u *orderServiceImpl) GetOrders(ctx context.Context) ([]model.Order, error) {
	users, err := u.repo.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}

func (u *orderServiceImpl) GetOrdersById(ctx context.Context, id uint64) (model.Order, error) {
	order, err := u.repo.GetOrdersByID(ctx, id)
	if err != nil {
		return model.Order{}, err
	}
	return order, err
}

func (u *orderServiceImpl) CreateOrder(ctx context.Context, req model.Order) (model.Order, error) {
	order := model.Order{
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
		Items:        req.Items,
	}

	// store to db
	res, err := u.repo.CreateOrder(ctx, order)
	if err != nil {
		return model.Order{}, err
	}
	return res, err
}

func (u *orderServiceImpl) UpdateOrder(ctx context.Context, order model.Order, id uint64) (model.Order, error) {
	res, err := u.repo.UpdateOrder(ctx, order, id)
	if err != nil {
		return model.Order{}, err
	}
	return res, err
}

func (u *orderServiceImpl) DeleteOrder(ctx context.Context, id uint64) error {
	err := u.repo.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
