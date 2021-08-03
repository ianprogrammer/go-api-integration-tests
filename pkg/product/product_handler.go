package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductRequest struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Handler struct {
	E              *echo.Echo
	ProductService IProductService
}

func RegisterProductHandlers(e *echo.Echo, productService IProductService) {
	ph := &Handler{
		E:              e,
		ProductService: productService,
	}
	ph.E.POST("/products", ph.saveProduct)
	ph.E.GET("/products/:id", ph.getByIdProduct)
	ph.E.GET("/products", ph.getAllProduct)
	ph.E.PUT("/products/:id", ph.updateProduct)
	ph.E.DELETE("/products/:id", ph.deleteProduct)
}

func (ph *Handler) saveProduct(c echo.Context) error {

	p := new(ProductRequest)

	if err := c.Bind(p); err != nil {
		return err
	}

	product, err := ph.ProductService.Post(Product{
		Name:  p.Name,
		Price: p.Price,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func (ph *Handler) getAllProduct(c echo.Context) error {

	products, err := ph.ProductService.GetAll()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, products)
}

func (ph *Handler) getByIdProduct(c echo.Context) error {
	id := c.Param("id")

	product, err := ph.ProductService.GetById(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)

}

func (ph *Handler) updateProduct(c echo.Context) error {
	id := c.Param("id")
	p := new(ProductRequest)

	if err := c.Bind(p); err != nil {
		return err
	}

	product, err := ph.ProductService.Update(id, Product{
		Name:  p.Name,
		Price: p.Price,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func (ph *Handler) deleteProduct(c echo.Context) error {
	id := c.Param("id")

	if err := ph.ProductService.Delete(id); err != nil {
		return err
	}
	return c.String(http.StatusNoContent, "")
}
