package main

import (
	"log"
	"net/http"
	"strconv"
)

// func timeHandler(format string) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		tm := time.Now().Format(format)
// 		w.Write([]byte("The time is: " + tm))
// 	}
// 	return http.HandlerFunc(fn)
// }

func showValue(value int) http.Handler {
	sv := func(w http.ResponseWriter, r *http.Request) {
		// tm := time.Now().Format(format)
		// w.Write([]byte("The value is: " + value))
		w.Write([]byte(strconv.Itoa(value)))
	}
	return http.HandlerFunc(sv)
}

func main() {
	mux := http.NewServeMux()

	// th := timeHandler(time.RFC1123)
	// mux.Handle("/time", th)
	dv := showValue(2000)
	mux.Handle("/user", dv)

	log.Print("Listening on port 3000...")
	http.ListenAndServe(":3000", mux)
}

// REFERENCE: https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go
