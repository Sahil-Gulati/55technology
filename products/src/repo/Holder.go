package repo;

import (
	stubs "../stubs"
)

type Holder struct{
	productsMap map[string][]interface{}
	products []*stubs.Product
}