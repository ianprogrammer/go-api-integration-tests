package product

type Service struct {
	Repository IProductRepository
}

type IProductService interface {
	GetById(id uint) (Product, error)
	GetAll() ([]Product, error)
	Post(p Product) (Product, error)
	Update(id uint, p Product) (Product, error)
	Delete(id uint) error
}

func NewService(repository IProductRepository) *Service {

	return &Service{
		Repository: repository,
	}
}

func (ps *Service) GetById(id uint) (Product, error) {
	return ps.Repository.SelectById(id)
}

func (ps *Service) GetAll() ([]Product, error) {
	return ps.Repository.SelectAll()
}

func (ps *Service) Post(p Product) (Product, error) {
	return ps.Repository.Insert(p)
}

func (ps *Service) Update(id uint, p Product) (Product, error) {
	return ps.Repository.Update(id, p)
}

func (ps *Service) Delete(id uint) error {
	return ps.Repository.Delete(id)
}
