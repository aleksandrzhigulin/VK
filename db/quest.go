package db

import (
	"VK/models"
	"fmt"
)

func (database Database) AddQuest(quest *models.Quest) error {
	var id int
	query := `INSERT INTO quests (name, cost) VALUES ($1, $2) RETURNING id`
	err := database.Connection.QueryRow(query, quest.Name, quest.Cost).Scan(&id)
	if err != nil {
		return err
	}
	quest.ID = id
	return nil
}

func (database Database) GetQuestCost(quest_id int) int {
	var cost int
	query := `SELECT cost FROM quests WHERE id = $1`
	err := database.Connection.QueryRow(query, quest_id).Scan(&cost)
	if err != nil {
		fmt.Println("error")
	}
	return cost
}
