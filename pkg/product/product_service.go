package product

type Service struct {
	Repository IProductRepository
}

type IProductService interface {
	GetById(id string) (Product, error)
	GetAll() ([]Product, error)
	Post(p Product) (Product, error)
	Update(id string, p Product) (Product, error)
	Delete(id string) error
}

func NewService(repository IProductRepository) *Service {

	return &Service{
		Repository: repository,
	}
}

func (ps *Service) GetById(id string) (Product, error) {
	return ps.Repository.SelectById(id)
}

func (ps *Service) GetAll() ([]Product, error) {
	return ps.Repository.SelectAll()
}

func (ps *Service) Post(p Product) (Product, error) {
	return ps.Repository.Insert(p)
}

func (ps *Service) Update(id string, p Product) (Product, error) {
	return ps.Repository.Update(id, p)
}

func (ps *Service) Delete(id string) error {
	return ps.Repository.Delete(id)
}
