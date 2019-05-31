package infrastructure

import (
	. "politrons/monads"
	"politrons/tools/DI/domain"
)

//Implementation type for the interface UserRepositoryImpl
type UserRepositoryImpl struct {
	dao DAO
}

/*
Constructor of the the UserRepositoryImpl.
Here repository has a dependency with DAOImpl which is the responsible to persist the [User]
here we dont have IOC since both Repo and DAO are in the same module.
*/
func NewUserRepository() UserRepositoryImpl {
	return UserRepositoryImpl{NewDAO()}
}

/*
Here we apply IOC and we implement the interface [UserRepositoryImpl] avoiding to have a dependency with
Higher level [infrastructure]
*/
func (repository UserRepositoryImpl) Save(user domain.User) chan Either {
	//Some logic of orchestration to persist the User.
	//Maybe multiple interactions with the DAO to have transactional
	return repository.dao.Save(user)
}
