package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":8000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	html := "<html><body><h1>Javier Castellanos <br>I like Not Monday<br>I dont like Mondays</h1><br><a href = '/greetings'>Greetings </a></body></html>"
	fmt.Fprintf(w, html)
}
