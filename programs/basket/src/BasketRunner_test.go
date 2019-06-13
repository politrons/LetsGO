package src

import (
	"application/resources"
	"testing"
)

/*
type testStruct struct {
	Test string
}

func parseGhPost(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	t := testStruct{}
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
		println("Error")
	}
	println("good")

	fmt.Println(t.Test)

	jsonResponse, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	response.Header().Set("Content-Type", "application/jsonResponse")
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write(jsonResponse)
}*/

func TestRun(t *testing.T) {
	resources.Main()
	/*	mux := http.NewServeMux()
		mux.HandleFunc("/parser", parseGhPost)
		http.ListenAndServe(":8080", mux)*/
}
