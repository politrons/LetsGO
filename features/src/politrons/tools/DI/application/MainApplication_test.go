package application

import (
	"fmt"
	"politrons/tools/DI/domain"
	"politrons/tools/DI/infrastructure"
	"testing"
)

/*
Entry point of your program, here we would receive a [Command] or [Query] and we would
start interacting with [domain] in case of Commands, and directly with the [infrastructure] in case of Queries.

Normally in DDD the application layer has dependencies with [Domain] and [Infrastructure]

To achieve IOC(Inversion of control) we create an instance here of [UserRepository] from Infrastructure, and we
set as dependency into the [Domain] which expect the interface that he define.
Having this we achieve that the lower level of the architecture never has a dependency with the higher level.
*/
func TestDependencyInjectionInDDD(t *testing.T) {
	userRepository := infrastructure.NewUserRepository()
	entityAggregateRoot := domain.CreateEntityAggregateRoot(userRepository)
	channelEither := entityAggregateRoot.RegisterUser("Politrons", "pol@gmail.com", "changeme")
	either := <-channelEither
	if either.IsLeft() {
		fmt.Println("Error persisting user:", either.Get().(error))
	} else {
		if (either.IsTypeOf(domain.User{})) {
			fmt.Println("User persisted:", either.Get().(domain.User))
		} else {
			fmt.Println("Wrong type returned in Either:", either.Get())
		}
	}
}
