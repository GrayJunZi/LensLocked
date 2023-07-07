package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func chiExerciseHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	w.Write([]byte(fmt.Sprintf("userID: %v", userID)))
}

func main() {
	r := chi.NewRouter()
	r.Route("/chiExercise", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/{userID}", chiExerciseHandler)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000 ...")
	http.ListenAndServe(":3000", r)
}
