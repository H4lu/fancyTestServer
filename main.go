package main

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":

	default:
		http.Error(w, "method not supported", http.Status)
	}

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

}

func main() {
	hub := NewHub()
	go hub.run()
	http.HandleFunc("/register", handleRegistration)
	http.HandleFunc("/login", handleLogin)

}
