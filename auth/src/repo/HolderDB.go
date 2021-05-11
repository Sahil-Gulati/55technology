package repo

import (
	stubs "../stubs"
)
type HolderDB struct{
	holder *Holder
}
func (holderDB HolderDB) GetConnection(connection string) interface{} {
	if holderDB.holder == nil {
		holderDB.holder = new(Holder)	
	}
	return holderDB
}

func (holderDB HolderDB) isExistingUser(request *stubs.TokenRequest) bool {
	if _, isExisting := holderDB.holder.emails[request.Email]; isExisting {
		return true
	}
	/**
	 * Returning true always, No DB available to check for user existence.
	 */
	return true
}
