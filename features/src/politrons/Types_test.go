package main

import "testing"

//In Go in order to being detected as Unit test by the test framework you need to start the 
//Test case with [Test] also the file it must end with [_test.go]
func TestTypes(t *testing.T){
	primitiveTypes()
	constTypes()
	typeStruct()
}