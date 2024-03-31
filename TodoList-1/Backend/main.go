package main

import (
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"github.com/pushpa/database"
	"github.com/pushpa/handlers"
	"github.com/pushpa/models"
)

var clients = make(map[*websocket.Conn]bool) 
var broadcast = make(chan models.TODO)         

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
		return true
    },
}

func main() {
    database.InitDB()
    defer database.DB.Close()

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
        var msg models.TODO
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            delete(clients, ws)
            break
        }
       
        switch msg.Action {
        case "add":
            handlers.InsertTodo(msg.Todo)
        case "edit":
            handlers.UpdateTodoText(msg.Id, msg.Todo)
        case "delete":
            handlers.DeleteTodo(msg.Id)
        case "handleIsDone":
            handlers.UpdateTodoIsDone(msg.Id, msg.IsDone)
        case "get":
            handlers.GetTodos()
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
    todos := handlers.GetTodos()
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


