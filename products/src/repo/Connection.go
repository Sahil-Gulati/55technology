package repo

import (
	stubs "../stubs"
)
type Connection interface{
	GetConnection(string) interface{}
	createProduct(*stubs.Product) (*stubs.Product, bool)
	deleteProductById(*stubs.ProductDelete) bool
	getAllProducts() *stubs.Products
}

type Database struct{
	Connection Connection
}
func (database Database) CreateProduct(product *stubs.Product) (*stubs.Product, bool) {
	return database.Connection.createProduct(product)
}
func (database Database) DeleteProductById(product *stubs.ProductDelete) bool {
	return database.Connection.deleteProductById(product)
}
func (database Database) GetAllProducts() *stubs.Products {
	return database.Connection.getAllProducts()
}