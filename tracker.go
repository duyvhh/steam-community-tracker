package tracker

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Hello, world")
}
