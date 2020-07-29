package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"test/model"
	"test/user"
)

// Login for user login
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var loginUser model.User
	res := make(map[string]string)

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		res["Message"] = "input erro"
		js, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
	}
	token, err := user.LogIn(&loginUser)
	if err != nil {
		res["Message"] = "internal error"
		js, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(js)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    token.Name,
		Value:   token.Token,
		Expires: token.Expire,
	})

	res["Message"] = "Success"
	js, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

// Register is for registering new users
func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var newUser model.User
	res := make(map[string]string)

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Fatal(err.Error())
	}

	if err := user.Register(&newUser); err != nil {
		res["Message"] = "Server error"
		js, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(js)
	}
}
