package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body := `<h1>Golang webapp running in a Docker container for the EDEM DevSecOps course</h1> <br> Now, rest a bit <br> <iframe src="https://chromedino.com/" frameborder="0" scrolling="no" width="100%" height="100%" loading="lazy"></iframe>
	<style type="text/css">iframe { position: absolute; width: 100%; height: 100%; z-index: 999; }</style>`
	fmt.Fprintf(w, body)
}

func main() {
	log.Print("Hello world webapp started.")

	http.HandleFunc("/", handler)

	port := "8080"

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}