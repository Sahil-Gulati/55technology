package services

import (
	"context"
	repo "../repo"
	stubs "../stubs"
	logger "../logger"
)
type AuthServer struct{
	stubs.UnimplementedAuthServiceServer
	Logger *logger.Logger
	Connection repo.HolderDB
	database *repo.Database
}
func (authServer AuthServer) Initiate() AuthServer {
	authServer.database = &repo.Database{Connection: authServer.Connection}
	return authServer
}
func (authServer AuthServer) GetToken(context context.Context, tokenRequest *stubs.TokenRequest) (*stubs.TokenResponse, error) {
	if authServer.database.IsExistingUser(tokenRequest) {
		return &stubs.TokenResponse{Token: authServer.generateToken()}, nil	
	}
	return &stubs.TokenResponse{}, nil
}

func (authServer AuthServer) generateToken() string {
	return JWT{}.GetToken()
}