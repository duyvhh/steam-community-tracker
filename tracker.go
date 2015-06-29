package tracker

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/trigger", handleTrigger)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "Hello, world")
}

func handleTrigger(rw http.ResponseWriter, req *http.Request) {
	resp, _ := http.Get("http://steamcommunity.com/market/listings/570/Scorching%20Talon")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
}
