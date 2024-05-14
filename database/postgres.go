package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"gitlab.com/miguelit0/toDoApp/models"

	_ "github.com/lib/pq" //Postgress works with this package
)

type PostgresRepository struct {
	db *sql.DB
}

// Create an instance of postgres database

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	// create connection with database
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

// Functions to implement the interface repository

func (repo *PostgresRepository) GetTasks(ctx context.Context) ([]models.Task, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM tasks")

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var tasks []models.Task
	for rows.Next() {
		task := models.Task{}

		if err = rows.Scan(&task.Id, &task.Title, &task.Description); err == nil {
			tasks = append(tasks, task)
		}

	}

	return tasks, err

}

func (repo *PostgresRepository) GetTaskById(ctx context.Context, id string) (*models.Task, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, title, description FROM tasks WHERE id = $1", id)

	fmt.Println("GET TASK BY UD")
	fmt.Println(rows)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var task = models.Task{}

	for rows.Next() {
		if err = rows.Scan(&task.Id, &task.Title, &task.Description); err == nil {
			return &task, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &task, nil

}

func (repo *PostgresRepository) CreateTask(ctx context.Context, task models.Task) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tasks (id, title, description) VALUES ($1, $2, $3)", task.Id, task.Title, task.Description)
	return err
}

func (repo *PostgresRepository) UpdateTask(ctx context.Context, task models.Task, taskId string) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE tasks SET title = $1, description = $2 WHERE id = $3", task.Title, task.Description, taskId)
	return err

}

func (repo *PostgresRepository) DeleteTask(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM tasks WHERE id = $1", id)
	return err
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
