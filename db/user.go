package db

import (
	"VK/models"
	"fmt"
)

func (database Database) AddUser(user *models.User) error {
	var id int
	query := `INSERT INTO users (name, balance) VALUES ($1, $2) RETURNING id`
	err := database.Connection.QueryRow(query, user.Name, 0).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (database Database) GetUserById(itemId int) (models.User, error) {
	user := models.User{}

	query := `SELECT * FROM users WHERE id = $1;`
	row := database.Connection.QueryRow(query, itemId)
	row.Scan(&user.ID, &user.Name, &user.Balance)
	err := row.Scan(&user.ID, &user.Name, &user.Balance)
	return user, err
}

func (database Database) UpdateUser(userId int, userData models.User) (models.User, error) {
	user := models.User{}
	query := `UPDATE users SET name=$1, balance=$2 WHERE id=$3 RETURNING id, name, balance;`
	err := database.Connection.QueryRow(query, userData.Name, userData.Balance, userId).Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		fmt.Println("no such user")
	}
	return user, nil
}

func (database Database) IncreaseBalance(userId int, change int) {
	user, _ := database.GetUserById(userId)
	user.Balance += change
	_, err := database.UpdateUser(userId, user)
	if err != nil {
		fmt.Println("increase balance error")
	}

}
