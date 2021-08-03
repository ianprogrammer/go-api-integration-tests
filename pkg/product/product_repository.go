package product

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type IProductRepository interface {
	Insert(p Product) (Product, error)
	Delete(id string) error
	Update(id string, p Product) (Product, error)
	SelectAll() ([]Product, error)
	SelectById(id string) (Product, error)
}

func (pr *Repository) Insert(p Product) (Product, error) {
	if result := pr.DB.Save(&p); result.Error != nil {
		return Product{}, result.Error
	}
	return p, nil
}

func (pr *Repository) Delete(id string) error {
	if result := pr.DB.Delete(&Product{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *Repository) Update(id string, p Product) (Product, error) {
	product, err := pr.SelectById(id)

	if err != nil {
		return Product{}, err
	}

	if result := pr.DB.Model(&product).Updates(p); result.Error != nil {
		return Product{}, result.Error
	}
	return product, nil
}

func (pr *Repository) SelectAll() ([]Product, error) {
	var products []Product

	if result := pr.DB.Find(&products); result.Error != nil {
		return []Product{}, result.Error
	}
	return products, nil
}

func (pr *Repository) SelectById(id string) (Product, error) {
	var product Product

	if result := pr.DB.First(&product, id); result.Error != nil {
		return product, result.Error
	}

	return product, nil

}
