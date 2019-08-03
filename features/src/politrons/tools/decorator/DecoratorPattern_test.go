package decorator

import (
	"fmt"
	"testing"
)

/**
Decorator pattern is about to wrap a functionality that we have originally and add an extra functionality but always using same
interface as the previous feature and invoking also the previous feature, adding the new one.

Here we have the functionality of one [Acoustic guitar] which emit a sound when we invoke function [Play].
But then we want to extend this functionality and make the sound amplified. So we need an electric guitar.
But we already have an acoustic guitar with strings and body, there's no way I create a new guitar from scratch.

I just need force when I create the [ElectricGuitar] that they pass me the old [AcousticGuitar] then
both guitar implement the interface Guitar and both implement the function [Play]
In the [AcousticGuitar] we just extract the sound. But in [ElectricGuitar] since is [decorating] the other guitar
we just invoke the [Play] of that guitar in our Play function and we extend the sound amplifying it!! Let's Rock!!!
*/
func TestDecoratorPattern(t *testing.T) {
	var acousticGuitar Guitar = new(AcousticGuitar)
	var electricGuitar Guitar = CreateElectricGuitar(acousticGuitar)
	fmt.Println(acousticGuitar.Play())
	fmt.Println(electricGuitar.Play())

}

//   Types
//------------
type AcousticGuitar struct {
	sound string
}

type ElectricGuitar struct {
	guitar Guitar
}

type Guitar interface {
	Play() string
}

/*
This function wrap the guitar into a new type ElectricGuitar decorator
*/
func CreateElectricGuitar(guitar Guitar) ElectricGuitar {
	return ElectricGuitar{guitar}
}

//Implementation of the [AcousticGuitar] just extract the sound from the guitar
func (g AcousticGuitar) Play() string {
	return g.sound
}

/*
Implementation of the [ElectricGuitar] since we decorate the [AcousticGuitar] we invoke the [Play]
of the AcousticGuitar and we add the Amplified feature.
*/
func (eg ElectricGuitar) Play() string {
	return eg.guitar.Play() + " Amplified"
}
