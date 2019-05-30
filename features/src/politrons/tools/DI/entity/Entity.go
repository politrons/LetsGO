package entity

//User for our system
type User struct {
	email       string
	password    string
	name        string
	dateOfBirth string
}

type SaveUser interface {
	Save(user User) error
}

//UseCase will store the dependency of [SaveUser] implemented in [Repository]
type UseCase struct {
	userSaver SaveUser
}

/*
Constructor that receive an interface [SaveUser] which in reality is the implementation of the [Repository]
*/
func NewUseCase(userSaver SaveUser) UseCase {
	return UseCase{userSaver}
}

/**
Function that use the internal dependency
*/
func (u UseCase) RegisterUser(user User) error {
	err := u.userSaver.Save(user)
	if err != nil {
		return err
	}
	return nil
}
