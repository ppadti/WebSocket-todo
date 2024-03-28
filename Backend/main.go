package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) 
var broadcast = make(chan TODO)         


type TODO struct {
    Id          int    `json:"id"`
    Action      string `json:"action"`
    Todo        string `json:"todo"`
    IsDone      bool   `json:"isDone,omitempty"`
}


var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("mysql", "root:PASSWORD/DATABASE?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal(err)
    }
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
		return true
    },
}

func main() {
    initDB()
    defer db.Close()

    http.HandleFunc("/ws", handleConnections)

    go handleMessages()

    log.Println("Server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
   	fmt.Println("connection initiated")
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    clients[ws] = true

   
    for {
        var msg TODO
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            delete(clients, ws)
            break
        }
       
        switch msg.Action {
        case "add":
            insertTodo(msg.Todo)
        case "edit":
            updateTodoText(msg.Id, msg.Todo)
        case "delete":
            deleteTodo(msg.Id)
        case "handleIsDone":
            updateTodoIsDone(msg.Id, msg.IsDone)
        case "get":
            getTodos()
        }
        
        broadcastTodos()
    }
}

func handleMessages() {
    for {
        <-broadcast
    }
}

func broadcastTodos() {
    todos := getTodos()
    for client := range clients {
        fmt.Println("write msg:", todos)
        err := client.WriteJSON(todos)
        if err != nil {
            log.Printf("error: %v", err)
            client.Close()
            delete(clients, client)
        }
    }
}

func getTodos() []TODO {
    rows, err := db.Query("SELECT * FROM todos LIMIT 100")
    if err != nil {
        log.Printf("error getting todos: %v", err)
        return nil
    }
    defer rows.Close()

    var todos []TODO
    for rows.Next() {
        var todo string
        var isDone bool
        var id int
        err := rows.Scan(&id, &todo, &isDone)
        if err != nil {
            log.Printf("error scanning todo: %v", err)
            continue
        }
        todos = append(todos, TODO{Action: "add",Id:id, Todo: todo, IsDone: isDone})
    }
    fmt.Println("todos", todos)
    return todos
}

func insertTodo(todo string) {
    id := rand.New(rand.NewSource(int64(time.Now().Nanosecond()))).Int31n(1000)
    _, err := db.Exec("INSERT INTO todos (id, todo, isDone) VALUES (?, ?, false)", id, todo)
    if err != nil {
        log.Printf("error inserting todo: %v", err)
    }
    fmt.Println("insertion succesful")
}

func updateTodoText(id int, newTodo string) {
    _, err := db.Exec("UPDATE todos SET todo = ? WHERE id = ?", newTodo,id)
    if err != nil {
        log.Printf("error updating todo text: %v", err)
    }
}

func deleteTodo(id int) {
    _, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
    if err != nil {
        log.Printf("error deleting todo: %v", err)
    }
}

func updateTodoIsDone(id int, isDone bool) {
    _, err := db.Exec("UPDATE todos SET is_done = ? WHERE id = ?", isDone, id)
    if err != nil {
        log.Printf("error updating todo isDone: %v", err)
    }
}


