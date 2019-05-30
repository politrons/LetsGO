package infrastructure

import "politrons/tools/DI/entity"

//Implementation type for the interface SaveUser
type UserRepository struct {
	response entity.User
}

//Constructor of the the UserRepository
func NewUserRepository() UserRepository {
	return UserRepository{}
}

/*
Here we apply IOC and we implement the interface [SaveUser] avoiding to have a dependency with
Higher level [infrastructure]
*/
func (repository UserRepository) Save(user entity.User) error {
	println("User persisted:", user)
	repository.response = user
	return nil
}
