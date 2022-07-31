package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "easy")
}

func main() {
	http.HandleFunc("/", Home)
	


	http.ListenAndServe(":3000", nil)
}
