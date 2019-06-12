package repositories

import (
	"time"
	"todo-api/database"
	"todo-api/models"
)

type Todo struct{}

func (t *Todo) Insert(mt models.Todo) error {
	_, err := database.DB.Exec(
		`INSERT INTO "TODO" (
			"USER_NAME",
			"CREATED_ON",
			"TITLE",
			"DESCRIPTION"
		) VALUES (
			:1, :2, :3, :4
		)`,
		mt.Username,
		time.Now(),
		mt.Title,
		mt.Description,
	)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) Delete(todoID int) error {
	_, err := database.DB.Exec(
		`DELETE FROM "TODO" WHERE "TODO_ID" = :1`,
		todoID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) SelectByUsername(username string) ([]models.Todo, error) {
	var ts []models.Todo
	rows, err := database.DB.Query(
		`SELECT
			"TODO_ID",
			"USER_NAME",
			"CREATED_ON",
			"TITLE",
			"DESCRIPTION"
		FROM "TODO"
		WHERE "USER_NAME" = :1`,
		username,
	)
	if err != nil {
		return ts, err
	}

	defer rows.Close()
	for rows.Next() {
		var t models.Todo
		err := rows.Scan(
			&t.TodoID,
			&t.Username,
			&t.CreatedOn,
			&t.Title,
			&t.Description,
		)
		if err != nil {
			return ts, err
		}
		ts = append(ts, t)
	}
	err = rows.Err()
	if err != nil {
		return ts, err
	}
	return ts, nil
}
