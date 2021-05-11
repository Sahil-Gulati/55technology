package repo

import (
	stubs "../stubs"
)
type Connection interface{
	GetConnection(string) interface{}
	isExistingUser(*stubs.TokenRequest) bool
}

type Database struct{
	Connection Connection
}
func (database Database) IsExistingUser(tokenRequest *stubs.TokenRequest) bool {
	return database.Connection.isExistingUser(tokenRequest)
}