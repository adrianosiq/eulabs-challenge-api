package repositories

import (
	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(product *models.Product) (*models.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) Delete(id int) error {
	var product models.Product
	err := r.db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}
