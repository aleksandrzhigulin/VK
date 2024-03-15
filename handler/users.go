package handler

import (
	"VK/models"
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func users(router chi.Router) {
	router.Post("/add", createUser)
	router.Route("/history/{userId}", func(router chi.Router) {
		router.Use(UserContext)
		router.Get("/", History)
	})
	router.Route("/{userId}", func(router chi.Router) {
		router.Use(UserContext)
		router.Get("/", getUserInfo)
	})
}

func UserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "userId")
		userIdInteger, _ := strconv.Atoi(userId)
		ctx := context.WithValue(r.Context(), "id", userIdInteger)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	user := &models.User{}
	render.Bind(request, user)
	dbInstance.AddUser(user)
	render.Render(writer, request, user)
}

func History(writer http.ResponseWriter, r *http.Request) {
	list := &models.QuestList{}
	id := r.Context().Value("id").(int)
	query := `SELECT * FROM quests_users INNER JOIN public.quests q on q.id = quests_users.quest_id WHERE user_id = $1`
	rows, err := dbInstance.Connection.Query(query, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		qId := 0
		uId := 0
		var quest models.Quest
		err := rows.Scan(&uId, &qId, &quest.ID, &quest.Name, &quest.Cost)
		if err != nil {
			fmt.Println(err.Error())
		}
		list.Quests = append(list.Quests, quest)
	}
	render.Render(writer, r, list)
}

func getUserInfo(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(int)
	user, err := dbInstance.GetUserById(id)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	render.Render(w, r, &user)
}
