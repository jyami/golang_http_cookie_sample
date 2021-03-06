package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/najeira/randstr"
)

var (
	// キーの長さは 16, 24, 32 バイトのいずれかでなければならない。
	// (AES-128, AES-192 or AES-256)
	key   = []byte(randstr.String(16))
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	log.Printf("secret\n")
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		log.Printf("%s\n", err.Error())
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	log.Printf("%v\n", session)

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	session.Values["authenticated"] = false
	session.Save(r, w)

	fmt.Fprintln(w, "secret message...")
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Printf("login\n")
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	log.Printf("secret\n")
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		log.Printf("%s\n", err.Error())
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	log.Printf("%v\n", session)

	session.Values["authenticated"] = false
	//delete(session.Values, "authenticated")
	session.Options.MaxAge = -1
	session.Save(r, w)

	session, err = store.Get(r, "cookie-name")
	log.Printf("%v\n", session)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
