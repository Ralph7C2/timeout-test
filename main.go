package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeout, err := strconv.Atoi(r.URL.Query().Get("timeout"))
		if err != nil {
			fmt.Fprintln(w, "Invalid tiemout")
			return
		}
		time.Sleep(time.Duration(timeout) * time.Second)
		fmt.Fprintln(w, "Hello, World!")
	}))
	log.Fatal(http.ListenAndServe(":10000", nil))
}
