package main

import (
	"net/http"
)

func main() {

}

func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
