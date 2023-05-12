package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/players", PlayerServer)
	http.ListenAndServe(":8080", nil)
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "30")
}
