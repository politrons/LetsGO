package domain

import . "politrons/monads"

//Entity of the domain
type User struct {
	Id       string
	Name     string
	Password string
	Email    string
}

/*
Interface with Save method to be implemented. To achieve IOC we define the interface
in the [domain] layer where we dont want to have any external dependency.
The [infrastructure] it will have a dependency with the [domain] and it will implement
this interface
*/
type UserRepository interface {
	Save(user User) chan Either
}

//EntityAggregateRoot will store the implementation of [UserRepository] implemented in [Repository]
type EntityAggregateRoot struct {
	userRepository UserRepository
}

/*
Constructor that receive an interface [UserRepository] which in reality is the implementation of the [Repository]
*/
func CreateEntityAggregateRoot(userRepository UserRepository) EntityAggregateRoot {
	return EntityAggregateRoot{userRepository}
}

/*
Function extension of [EntityAggregateRoot] that use the internal dependency implementation to persist the user.
*/
func (entityAggregateRoot EntityAggregateRoot) RegisterUser(name string, email string, password string) chan Either {
	user := User{Name: "Politrons", Email: "pol@gmail.com", Password: "changeme"}
	return entityAggregateRoot.userRepository.Save(user)
}
