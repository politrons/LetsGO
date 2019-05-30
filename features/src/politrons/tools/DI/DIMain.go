package DI

import (
	"politrons/tools/DI/entity"
	"politrons/tools/DI/infrastructure"
)

func main() {
	userRepository := infrastructure.NewUserRepository()
	useCase := entity.NewUseCase(userRepository)
	useCase.RegisterUser(entity.User{})
}
