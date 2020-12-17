package main

import (
	"fmt"
	"net/http"
)

func get_request(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
	fmt.Println(req.Body)
}

func main() {
	http.HandleFunc("/", get_request)
	http.ListenAndServe(":8001", nil)
}
