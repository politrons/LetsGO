package infrastructure

import (
	"fmt"
	"politrons/tools/DI/domain"
)

type DAO interface {
	Save(user domain.User) (chan domain.User, error)
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
func (dao DAOImpl) Save(user domain.User) (chan domain.User, error) {
	channel := make(chan domain.User)
	go func() {
		fmt.Println("Persisting user:", user)
		user.Id = "newId"
		channel <- user
	}()
	return channel, nil
}
