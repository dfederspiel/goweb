package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Legs bool   `json:"somethingElse"`
}

/*	DefaultHandler is the "default," or fallback handler for http requests.
*	It is called whenever there are no other handlers defined in your routes */
func DefaultHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, I am the default handler for this service.")
}

func AnimalHandler(w http.ResponseWriter, r *http.Request) {
	a := Animal{"Duchess", 4, true}
	jsonA, _ := json.Marshal(a)
	w.Write(jsonA)

}

func (a Animal) Speak(whatToSay string) string {
	return whatToSay
}
