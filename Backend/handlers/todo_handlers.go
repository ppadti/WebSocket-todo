package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/pushpa/database"
	"github.com/pushpa/models"
)

func GetTodos() []models.TODO {
    rows, err := database.DB.Query("SELECT * FROM todos LIMIT 100")
    if err != nil {
        log.Printf("error getting todos: %v", err)
        return nil
    }
    defer rows.Close()

    var todos []models.TODO
    for rows.Next() {
        var todo string
        var isDone bool
        var id int
        err := rows.Scan(&id, &todo, &isDone)
        if err != nil {
            log.Printf("error scanning todo: %v", err)
            continue
        }
        todos = append(todos, models.TODO{Action: "add",Id:id, Todo: todo, IsDone: isDone})
    }
    fmt.Println("todos", todos)
    return todos
}

func InsertTodo(todo string) {
    id := rand.New(rand.NewSource(int64(time.Now().Nanosecond()))).Int31n(1000)
    _, err := database.DB.Exec("INSERT INTO todos (id, todo, isDone) VALUES (?, ?, false)", id, todo)
    if err != nil {
        log.Printf("error inserting todo: %v", err)
    }
    fmt.Println("insertion succesful")
}

func UpdateTodoText(id int, newTodo string) {
    _, err := database.DB.Exec("UPDATE todos SET todo = ? WHERE id = ?", newTodo,id)
    if err != nil {
        log.Printf("error updating todo text: %v", err)
    }
}

func DeleteTodo(id int) {
    _, err := database.DB.Exec("DELETE FROM todos WHERE id = ?", id)
    if err != nil {
        log.Printf("error deleting todo: %v", err)
    }
}

func UpdateTodoIsDone(id int, isDone bool) {
    _, err := database.DB.Exec("UPDATE todos SET isDone = ? WHERE id = ?", isDone, id)
    if err != nil {
        log.Printf("error updating todo isDone: %v", err)
    }
}

