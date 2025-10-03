package main

import "net/http"

func handleHttp(w http.ResponseWriter, r *http.Request) {
	queryTest := r.URL.Query().Get("test")
	response := "no query params"

	if queryTest != "" {
		response = queryTest
	}

	w.Write([]byte(response))
}

func main() {

	http.HandleFunc("/", handleHttp)

	http.ListenAndServe(":8080", nil)

}
