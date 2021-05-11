package services

import (
	"errors"
	"context"
	repo "../repo"
	stubs "../stubs"
	logger "../logger"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)
type ProductsServer struct{
	stubs.UnimplementedProductServiceServer
	Connection repo.Connection
	database *repo.Database
	Logger *logger.Logger
}
func (ps ProductsServer) Initiate() ProductsServer {
	ps.database = &repo.Database{Connection: ps.Connection}
	return ps
}
func (ps ProductsServer) GetProducts(ctx context.Context, in *emptypb.Empty) (*stubs.Products, error) {
	return ps.database.GetAllProducts(), nil
}
func (ps ProductsServer) CreateProduct(ctx context.Context, in *stubs.Product) (*stubs.Product, error) {
	if product, isCreated := ps.database.CreateProduct(in); isCreated {
		return product, nil
	}
	return nil, errors.New("Failed to add product!");
}
func (ps ProductsServer) DeleteProduct(ctx context.Context, in *stubs.ProductDelete) (*stubs.ProductDelete, error) {
	if isDeleted := ps.database.DeleteProductById(in); isDeleted {
		return in, nil
	}
	return nil, errors.New("Failed to delete product!");
}