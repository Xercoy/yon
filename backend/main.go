package main

import (
	"flag"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	responseStr, err := prepareResponse()
	if err != nil {
		// add error method
		fmt.Fprintf(w, "Error")
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// CORS
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	//if r.Method == "OPTIONS" {
	//	return
	//}

	fmt.Fprint(w, responseStr)
}

func main() {
	addr := flag.String("address", ":8065", "host address")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.ListenAndServe(*addr, nil)
}
