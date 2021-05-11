package repo

import (
	stubs "../stubs"
	utils "../utils"
)
type HolderDB struct{
	holder *Holder
}
func (holderDB HolderDB) GetConnection(connection string) interface{} {
	if holderDB.holder == nil {
		holderDB.holder = new(Holder)	
		holderDB.holder.productsMap = make(map[string][]interface{})
	}
	return holderDB
}
func (holderDB HolderDB) createProduct(product *stubs.Product) (*stubs.Product, bool) {
	unixNano := utils.ToString(utils.GetCurrentUnixNano())
	product.Id = unixNano
	holderDB.holder.productsMap[unixNano] = []interface{}{product, len(holderDB.holder.products)}
	holderDB.holder.products = append(holderDB.holder.products, product)	
	return product, true
}
func (holderDB HolderDB) deleteProductById(product *stubs.ProductDelete) bool {
	productId := product.Id
	if value, ok := holderDB.holder.productsMap[productId]; ok {
		index := value[1].(int)
		delete(holderDB.holder.productsMap, productId)
		holderDB.holder.products = append(holderDB.holder.products[:index], holderDB.holder.products[index+1:]...)
		return true
	}
	return false
}
func (holderDB HolderDB) getAllProducts() *stubs.Products {
	return &stubs.Products{
		Products: holderDB.holder.products,
	}
}
