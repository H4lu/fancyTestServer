package user


type User struct {
	Email string `json:"email"`,
	Password string `json:"password"`
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		var user User
		err := json.Unmarshal(r.Body, &user)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			fmt.Println("Error decoding json in registratoin", err)
			return
		}
		err = saveUser(user)
        if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			fmt.Println("Error to save new user", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "method not supported", http.Status)
	}

}

func handleLogin(w http.ResponseWriter, r *http.Request) {
 
}