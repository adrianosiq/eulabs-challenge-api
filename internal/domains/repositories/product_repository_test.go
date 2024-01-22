package repositories

import (
	"database/sql/driver"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
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
		expectedSQL := "SELECT (.+) FROM `products` WHERE `products`.`deleted_at` IS NULL"
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

		productRepository := NewProductRepository(db)
		products, err := productRepository.GetAll()

		assert.NoError(t, err)
		assert.True(t, len(products) == 2)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		db, mock := NewMockDB()
		rows := sqlmock.NewRows([]string{"id", "title", "description", "price", "created_at", "updated_at", "deleted_at"})
		expectedSQL := "SELECT (.+) FROM `products` WHERE `products`.`deleted_at` IS NULL"
		mock.ExpectQuery(expectedSQL).WillReturnRows(rows)

		productRepository := NewProductRepository(db)
		products, err := productRepository.GetAll()

		assert.NoError(t, err)
		assert.Empty(t, products)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "SELECT (.+) FROM `products` WHERE `products`.`deleted_at` IS NULL"
		mock.ExpectQuery(expectedSQL).WillReturnError(fmt.Errorf("some error"))

		productRepository := NewProductRepository(db)
		_, err := productRepository.GetAll()

		assert.Error(t, err)
	})
}

func TestCreate(t *testing.T) {
	var mockCreateProduct = &models.Product{
		Title:       "Charmander",
		Description: "It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.",
		Price:       1093.45,
	}

	t.Run("should return an product", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "INSERT INTO `products` (.+) VALUES (.+)"
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		productRepository := NewProductRepository(db)
		product, err := productRepository.Create(mockCreateProduct)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), product.ID)
		assert.Equal(t, "Charmander", product.Title)
		assert.Contains(t, product.Description, "It has a preference")
		assert.Equal(t, 1093.45, product.Price)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "INSERT INTO `products` (.+) VALUES (.+)"
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectCommit()

		productRepository := NewProductRepository(db)
		_, err := productRepository.Create(mockCreateProduct)

		assert.Error(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("should return the product", func(t *testing.T) {
		db, mock := NewMockDB()
		row := sqlmock.NewRows([]string{"id", "title", "description", "price", "created_at", "updated_at", "deleted_at"}).
			AddRow(1, "Bulbasaur", "There is a plant seed on its back right from the day this Pokémon is born. The seed slowly grows larger.", 99.99, time.Now(), time.Now(), nil)
		expectedSQL := "SELECT (.+) FROM `products` WHERE `products`.`id` = (.+) AND `products`.`deleted_at` IS NULL"
		mock.ExpectQuery(expectedSQL).WillReturnRows(row)

		productRepository := NewProductRepository(db)
		product, err := productRepository.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), product.ID)
		assert.Equal(t, "Bulbasaur", product.Title)
		assert.Contains(t, product.Description, "There is a plant seed")
		assert.Equal(t, 99.99, product.Price)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "SELECT (.+) FROM `products` WHERE `products`.`id` = (.+) AND `products`.`deleted_at` IS NULL"
		mock.ExpectQuery(expectedSQL).WillReturnError(fmt.Errorf("some error"))

		productRepository := NewProductRepository(db)
		_, err := productRepository.GetByID(1)

		assert.Error(t, err)
	})
}

func TestUpdate(t *testing.T) {
	var mockUpdateProduct = &models.Product{
		ID:          1,
		Title:       "Charmander",
		Description: "It has a preference for hot things. When it rains, steam is said to spout from the tip of its tail.",
		Price:       1093.45,
	}

	t.Run("should return the product", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "UPDATE `products` SET .+"
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		productRepository := NewProductRepository(db)
		product, err := productRepository.Update(mockUpdateProduct)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), product.ID)
		assert.Equal(t, "Charmander", product.Title)
		assert.Contains(t, product.Description, "It has a preference")
		assert.Equal(t, 1093.45, product.Price)
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "UPDATE `products` SET .+"
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectCommit()

		productRepository := NewProductRepository(db)
		_, err := productRepository.Update(mockUpdateProduct)

		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "UPDATE `products` (.+) WHERE id = (.+) AND `products`.`deleted_at` IS NULL"
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		productRepository := NewProductRepository(db)
		err := productRepository.Delete(1)

		assert.Nil(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("should return an error", func(t *testing.T) {
		db, mock := NewMockDB()
		expectedSQL := "UPDATE `products` (.+) WHERE id = (.+) AND `products`.`deleted_at` IS NULL"
		mock.ExpectBegin()
		mock.ExpectExec(expectedSQL).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectCommit()

		productRepository := NewProductRepository(db)
		err := productRepository.Delete(1)

		assert.Error(t, err)
	})
}
