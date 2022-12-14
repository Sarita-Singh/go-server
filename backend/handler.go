package backend

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world form backend")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func initializeRoutes() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/increment", incrementCounter)

}
func Run(addr string) {
	initializeRoutes()
	fmt.Println("Serving and listening at port:", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}
