package decorator

import (
	"fmt"
	"testing"
)

func TestDecoratorPattern(t *testing.T) {
	var acousticGuitar Guitar = new(AcousticGuitar)
	var electricGuitar Guitar = CreateElectricGuitar(acousticGuitar)
	fmt.Println(acousticGuitar.Play())
	fmt.Println(electricGuitar.Play())

}

type AcousticGuitar struct {
	sound string
}

type ElectricGuitar struct {
	guitar Guitar
}

type Guitar interface {
	Play() string
}

func CreateElectricGuitar(guitar Guitar) ElectricGuitar {
	return ElectricGuitar{guitar}

}

func (eg ElectricGuitar) Play() string {
	return eg.guitar.Play() + " Amplified"
}

func (g AcousticGuitar) Play() string {
	return g.sound
}
