package main

import (
	"selfhosted/http"
)

func main() {
	s := http.NewServer(":4000")

	s.ListenAndServe()
}
