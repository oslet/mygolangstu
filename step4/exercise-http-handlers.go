package main

import (
	"log"
	"net/http"
)

type String string

type Struct struct {
	greeting string
	punct string
	who string
}

func (e *Struct) ServeHTTP(
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})) {
	fmt.Println(e)
}
    
func main() {
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
