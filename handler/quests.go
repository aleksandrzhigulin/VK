package handler

import (
	"VK/models"
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func quests(router chi.Router) {
	router.Post("/add", addQuest)
	router.Route("/completed/{user_id}/{quest_id}", func(router chi.Router) {
		router.Use(QuestContext)
		router.Get("/", completeQuest)
	})
}

func completeQuest(writer http.ResponseWriter, request *http.Request) {
	userId := request.Context().Value("user_id").(string)
	questId := request.Context().Value("quest_id").(string)
	questIdInteger, _ := strconv.Atoi(questId)
	userIdInteger, _ := strconv.Atoi(userId)

	// Check if task is already completed
	checkQuery := `SELECT * FROM quests_users WHERE user_id = $1 AND quest_id = $2`
	var checkUID int
	var checkQID int
	dbInstance.Connection.QueryRow(checkQuery, userIdInteger, questIdInteger).Scan(&checkUID, &checkQID)
	if (checkUID != 0) && (checkQID != 0) {
		render.Render(writer, request, ErrAlreadyCompleted)
		return
	}
	cost := dbInstance.GetQuestCost(questIdInteger)
	dbInstance.IncreaseBalance(userIdInteger, cost)

	query := `INSERT INTO quests_users (user_id, quest_id) VALUES ($1, $2)`
	dbInstance.Connection.QueryRow(query, userIdInteger, questIdInteger)
	return
}

func QuestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_id := chi.URLParam(r, "user_id")
		quest_id := chi.URLParam(r, "quest_id")
		ctx := context.WithValue(r.Context(), "user_id", user_id)
		ctx = context.WithValue(ctx, "quest_id", quest_id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func addQuest(writer http.ResponseWriter, request *http.Request) {
	quest := &models.Quest{}
	render.Bind(request, quest)
	dbInstance.AddQuest(quest)
	render.Render(writer, request, quest)
}
