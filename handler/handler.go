package handler

import (
	"VK/db"
	"github.com/go-chi/chi"
	"net/http"
)

var dbInstance db.Database

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.Route("/user", users)
	router.Route("/quest", quests)
	return router
}
