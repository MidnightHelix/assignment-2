package repository

import (
	"context"

	"github.com/MidnightHelix/assignment-2/internal/infrastructure"
	"github.com/MidnightHelix/assignment-2/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderQuery interface {
	GetOrders(ctx context.Context) ([]model.Order, error)
	GetOrdersByID(ctx context.Context, id uint64) (model.Order, error)
	DeleteOrder(ctx context.Context, id uint64) error
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	UpdateOrder(ctx context.Context, order model.Order, id uint64) (model.Order, error)
}

type OrderCommand interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
}

type orderQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewOrderQuery(db infrastructure.GormPostgres) OrderQuery {
	return &orderQueryImpl{db: db}
}

func (u *orderQueryImpl) GetOrders(ctx context.Context) ([]model.Order, error) {
	db := u.db.GetConnection()
	orders := []model.Order{}
	if err := db.
		WithContext(ctx).
		Table("orders").
		Preload("Items").
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (u *orderQueryImpl) GetOrdersByID(ctx context.Context, id uint64) (model.Order, error) {
	db := u.db.GetConnection()
	order := model.Order{}
	if err := db.
		WithContext(ctx).
		Table("orders").
		Where("id = ?", id).
		Find(&order).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (u *orderQueryImpl) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("orders").
		Save(&order).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (u *orderQueryImpl) UpdateOrder(ctx context.Context, order model.Order, id uint64) (model.Order, error) {
	db := u.db.GetConnection()
	if err := db.
		Session(&gorm.Session{FullSaveAssociations: true}).
		WithContext(ctx).
		Table("orders").
		Save(&order).Error; err != nil {
		return model.Order{}, err
	}
	return order, nil
}

func (u *orderQueryImpl) DeleteOrder(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		Session(&gorm.Session{FullSaveAssociations: true}).
		WithContext(ctx).
		Table("orders").
		Where("id = ?", id).
		Select(clause.Associations).
		Delete(&model.Order{ID: id}).Error; err != nil {
		return err
	}
	return nil
}
