package application

import (
	"fmt"
	"politrons/tools/DI/domain"
	"politrons/tools/DI/infrastructure"
	"testing"
)

func TestDependencyInjectionInDDD(t *testing.T) {
	userRepository := infrastructure.NewUserRepository()
	entityAggregateRoot := domain.CreateEntityAggregateRoot(userRepository)
	user, err := entityAggregateRoot.RegisterUser("Politrons", "pol@gmail.com", "changeme")
	if err != nil {
		fmt.Println("Error persisting user:", err)
	} else {
		fmt.Println("User persisted:", user)
	}
}
