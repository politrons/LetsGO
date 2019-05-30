package infrastructure

import (
	"fmt"
	"politrons/tools/DI/domain"
)

type DAO interface {
	Save(user domain.User) (domain.User, error)
}

//Type that it will Implement the interface DAO
type DAOImpl struct {
	//It should contain the ORM to persist the User
}

//Constructor of the the DAOImpl
func NewDAO() DAOImpl {
	return DAOImpl{}
}

//Implementation of the interface [DAO] and type [DAOImpl]
func (dao DAOImpl) Save(user domain.User) (domain.User, error) {
	fmt.Println("Persisting user:", user)
	user.Id = "newId"
	return user, nil
}
