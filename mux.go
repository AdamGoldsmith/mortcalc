package main

import (
	"log"
	"net/http"
	"time"
)

// REFERENCE: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go

// 1. THIS...

// type timeHandler struct {
// 	format string
// }

// func (th timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	tm := time.Now().Format(th.format)
// 	w.Write([]byte("The time is: " + tm))
// }

// func main() {
// 	mux := http.NewServeMux()

// 	// Initialise the timeHandler in exactly the same way we would any normal
// 	// struct.
// 	th := timeHandler{format: time.RFC1123}

// 	// Like the previous example, we use the mux.Handle() fnction to register
// 	// this with our ServeMux.
// 	mux.Handle("/time", th)

// 	log.Print("Listening on port 3000...")
// 	http.ListenAndServe(":3000", mux)
// }

// ...CAN BE RE-WRITTEN AS: (to reduce code)

// 2. BUT THIS...

// func timeHandler(w http.ResponseWriter, r *http.Request) {
// 	tm := time.Now().Format(time.RFC1123)
// 	w.Write([]byte("The time is: " + tm))
// }

// func main() {
// 	mux := http.NewServeMux()

// 	// Convert the timeHandler function to a http.HandlerFunc type.
// 	th := http.HandlerFunc(timeHandler)

// 	// And add it to the ServeMux.
// 	mux.Handle("/time", th)
// 	// mux.Handle("/user", "adam")

// 	log.Print("Listening on port 3000...")
// 	http.ListenAndServe(":3000", mux)
// }

// ...CAN BE RE-WRITTEN AS: (to reduce code and use built-in handlers)

// 3. BUT THIS...

// func timeHandler(w http.ResponseWriter, r *http.Request) {
// 	tm := time.Now().Format(time.RFC1123)
// 	w.Write([]byte("The time is: " + tm))
// }

// func main() {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/time", timeHandler)

// 	log.Print("Listening on port 3000...")
// 	http.ListenAndServe(":3000", mux)
// }

// ...CAN BE RE-WRITTEN AS: (to allow passing in variables to handlers)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()

	th := timeHandler(time.RFC1123)
	mux.Handle("/time", th)

	log.Print("Listening on port 3000...")
	http.ListenAndServe(":3000", mux)
}
