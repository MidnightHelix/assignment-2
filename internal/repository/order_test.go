package repository

import (
	"context"
	"errors"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MidnightHelix/assignment-2/internal/infrastructure/mocks"
	"github.com/MidnightHelix/assignment-2/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newMockGorm() (*gorm.DB, sqlmock.Sqlmock) {
	//db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}
	return gormDB, mock
}

func TestGetUsers(t *testing.T) {
	t.Run("error get orders", func(t *testing.T) {
		db, mock := newMockGorm()

		postgresMock := mocks.NewGormPostgres(t)
		postgresMock.On("GetConnection").Return(db)

		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT * FROM "orders"
		`)).WillReturnError(errors.New("some error"))

		userRepo := orderQueryImpl{db: postgresMock}
		res, err := userRepo.GetOrders(context.Background())
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(res))
	})

	t.Run("success get orders", func(t *testing.T) {
		db, mock := newMockGorm()

		postgresMock := mocks.NewGormPostgres(t)
		postgresMock.On("GetConnection").Return(db)

		orderRow := sqlmock.
			NewRows([]string{"id", "customer_name", "ordered_at"}).
			AddRow(1, "testing", time.Now())

		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT * FROM "orders"
		`)).WillReturnRows(orderRow)

		itemRow := sqlmock.
			NewRows([]string{"id", "name", "order_id"}).
			AddRow(1, "item_name", 1)

		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT * FROM "items" WHERE "items"."order_id" = $1
		`)).WithArgs(1).WillReturnRows(itemRow)

		userRepo := orderQueryImpl{db: postgresMock}
		res, err := userRepo.GetOrders(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
	})

}

func TestCreateOrder(t *testing.T) {
	t.Run("error create order", func(t *testing.T) {
		db, mock := newMockGorm()

		postgresMock := mocks.NewGormPostgres(t)
		postgresMock.On("GetConnection").Return(db)

		mock.ExpectExec(regexp.QuoteMeta(`
			INSERT INTO "orders" VALUES ($1, $2, $3)
		`)).WillReturnError(errors.New("some error"))

		userRepo := orderQueryImpl{db: postgresMock}
		order := model.Order{ID: 1, CustomerName: "test", OrderedAt: time.Now()}
		res, err := userRepo.CreateOrder(context.Background(), order)
		assert.NotNil(t, err)
		assert.Equal(t, model.Order{}, res)
	})

}

func TestDeleteOrder(t *testing.T) {
	t.Run("error deleting order", func(t *testing.T) {
		db, mock := newMockGorm()
		postgresMock := mocks.NewGormPostgres(t)
		postgresMock.On("GetConnection").Return(db)

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`
			DELETE FROM "orders" WHERE "id" = $1
		`)).WithArgs(123).WillReturnError(errors.New("some error"))
		mock.ExpectRollback()

		u := orderQueryImpl{db: postgresMock}
		err := u.DeleteOrder(context.Background(), 123)

		assert.NotNil(t, err)
	})

}
