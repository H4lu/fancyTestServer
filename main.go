package main

import "net/http"

func main() {
	hub := NewHub()
	go hub.run()
	http.HandleFunc("/register", handleRegistration)
	http.HandleFunc("/login", handleLogin)

}
