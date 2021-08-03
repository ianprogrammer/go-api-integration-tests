package product

type Product struct {
	ID    string `gorm:"-"`
	Name  string
	Price int64
}
