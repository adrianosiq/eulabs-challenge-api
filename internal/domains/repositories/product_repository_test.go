package repositories

import (
	"database/sql/driver"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestGetAll(t *testing.T) {
	t.Run("should return a list the products", func(t *testing.T) {
		db, mock := NewMockDB()
		values := [][]driver.Value{
			{1, "Bulbasaur", "There is a plant seed on its back right from the day this Pokémon is born. The seed slowly grows larger.", 99.99, time.Now(), time.Now(), nil},
			{2, "Charmander", "It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.", 1093.45, time.Now(), time.Now(), nil},
		}
		rows := sqlmock.NewRows([]string{"id", "title", "description", "price", "created_at", "updated_at", "deleted_at"}).AddRows(values...)
		mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

		productRepository := NewProductRepository(db)
		products, err := productRepository.GetAll()

		assert.NoError(t, err)
		assert.True(t, len(products) == 2)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		db, mock := NewMockDB()
		rows := sqlmock.NewRows([]string{"id", "title", "description", "price", "created_at", "updated_at", "deleted_at"})
		mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

		productRepository := NewProductRepository(db)
		products, err := productRepository.GetAll()

		assert.NoError(t, err)
		assert.Empty(t, products)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock := NewMockDB()
		mock.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("some error"))

		productRepository := NewProductRepository(db)
		_, err := productRepository.GetAll()

		assert.Error(t, err)
	})
}
