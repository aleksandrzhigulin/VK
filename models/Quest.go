package models

import (
	"fmt"
	"net/http"
)

type Quest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

type QuestList struct {
	Quests []Quest `json:"quests"`
}

func (q *QuestList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Quest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (quest *Quest) Bind(request *http.Request) error {
	if quest.Name == "" || quest.Cost <= 0 {
		fmt.Errorf("Unable to create quest")
	}
	return nil
}
