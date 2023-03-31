package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/greetings", greetings)
	http.HandleFunc("/random", random)

	http.ListenAndServe(":8000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {

	html := "<html><body><h1>Javier Castellanos <br>I like Not Monday<br>I dont like Mondays</h1><br><a href = '/'>home </a> <a href = '/greetings'>Greetings </a> <a href = '/random'>random </a></body></html>"
	fmt.Fprintf(w, html)
}

func greetings(w http.ResponseWriter, r *http.Request) {

	time := time.Now()
	// timeHour := strconv.Itoa(time.Hour())
	// timeMinute := strconv.Itoa(time.Minute())

	timeString := time.Format("3:04 PM")

	day := time.Weekday().String()

	html2 := "<html><body><h1>Greetings <br> right now is " + timeString + " enjoy the rest of your " + day + "</h1> <a href = '/'>home </a> <a href = '/greetings'>Greetings </a> <a href = '/random'>random </a> </body></html>"

	fmt.Fprintf(w, html2)
}

func random(w http.ResponseWriter, r *http.Request) {

	quoteMap := make(map[int]string)

	// Add six key-value pairs to the map
	quoteMap[0] = "Let’s eat kids. Let’s eat, kids. Punctuation saves lives!"
	quoteMap[1] = "I try to keep an open mind about everything except grammar and punctuation."
	quoteMap[2] = "Similes are like metaphors."
	quoteMap[3] = "My life is a constant battle between wanting to correct grammar and wanting friends."
	quoteMap[4] = "Theiyr’re. Take that, Grammar Police."
	quoteMap[5] = "Quote Number 6"

	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(quoteMap))

	quote := quoteMap[index]

	html := "<html><body><h1>" + quote + "</h1><br><a href = '/'>home </a> <a href = '/greetings'>Greetings </a> <a href = '/random'>random </a></body></html>"
	fmt.Fprintf(w, html)
}
