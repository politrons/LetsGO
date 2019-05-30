package infrastructure

import (
	"fmt"
	. "politrons/monads"
	"politrons/tools/DI/domain"
)

type DAO interface {
	Save(user domain.User) chan Either
}

//Type that it will Implement the interface DAO
type DAOImpl struct {
	//It should contain the ORM to persist the User
}

//Constructor of the the DAOImpl
func NewDAO() DAOImpl {
	return DAOImpl{}
}

/*
Implementation of the interface [DAO] and type [DAOImpl]
Access to database normally take time and it's blocking, so it's more efficient make it async
*/
func (dao DAOImpl) Save(user domain.User) chan Either {
	channel := make(chan Either)
	go func() {
		fmt.Println("Persisting user:", user)
		user.Id = "newId"
		channel <- Right{user}
	}()
	return channel
}
