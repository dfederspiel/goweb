package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*DefaultHandler is the "default," or fallback handler for http requests.
*	It is called whenever there is no other handler defined */
func DefaultHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "this is just plain text that we want to send back in our response. we can use ")
}

func AnimalHandler(w http.ResponseWriter, r *http.Request) {
	a := Animal{"Duchess", 4, true}
	jsonA, _ := json.Marshal(a)
	w.Write(jsonA)

}

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Legs bool   `json:"somethingElse"`
}

func (a Animal) Speak(whatToSay string) string {
	return whatToSay
}
