package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"totality/internal/app"

	"github.com/go-chi/chi"
)

func main() {
	port := "4000"

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Get("/users", userList)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func userList(w http.ResponseWriter, r *http.Request) {
	userIDs := r.URL.Query().Get("ids")
	userID := []int{}

	for _, a := range strings.Split(userIDs, ",") {
		b, _ := strconv.Atoi(a)
		if b != 0 {
			userID = append(userID, b)
		}
	}

	app.UserList(w, userID)
}
