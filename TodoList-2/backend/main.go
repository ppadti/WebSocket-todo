package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/ppadti/todo-app/database"
	"github.com/rs/cors"
)

type TODO struct {

	Id     int    `json:"id"`
	Todo   string `json:"todo"`
	IsDone bool   `json:"isDone"`
}

// type Message struct {
// 	Content string `json:"conetnt"`
// }

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan TODO)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	// 	origin := r.Header.Get("Origin")

	// switch origin {
	// 	// Update this to HTTPS
	// case "https://localhost:8080":
	// 	return true
	// default:
	// 	return false
	// }
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connection initiated")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		var msg TODO
		err := conn.ReadJSON(&msg)
		// _, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			break
		}
		fmt.Println(msg)
		broadcastToClients()
	}
}

func handleMessages() {
    for {
        <-broadcast
    }
}

func broadcastToClients() {
	fmt.Println("websocket broadcast to client called")
		todos := getTodos()
		for client := range clients {
			fmt.Println("write message:", todos)
			err := client.WriteJSON(todos)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients,client)
			}
		}
	}


func getTodos() []TODO {
	rows, err := database.DB.Query("select * from todos")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	 var todos []TODO
	for rows.Next() {
		var todo TODO
		err := rows.Scan(&todo.Id, &todo.Todo, &todo.IsDone)
		if err != nil {
			fmt.Println(err)
		}
		todos = append(todos, todo)
	}
	return todos


}

func createTodos(w http.ResponseWriter, r *http.Request) {
	var todo TODO

	json.NewDecoder(r.Body).Decode(&todo)
	stmt, err := database.DB.Prepare("insert into todos values (?,?,0)")
	if err != nil {
		fmt.Println(err)
	}
	id := rand.New(rand.NewSource(int64(time.Now().Nanosecond()))).Int31n(1000)

	res, err := stmt.Exec(id, todo.Todo)
	if err != nil {
		fmt.Println(err)
	}

	todoId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	todo.Id = int(todoId)
	
	json.NewEncoder(w).Encode(todo)

	fmt.Println("create todo result", todo)
}

func deleteTodos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]

	stmt, err := database.DB.Prepare("delete from todos where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(todoId)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(res)

	// fmt.Fprintf(w, "todo deleted successfullys")
}

func updateTodos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	var todo TODO
	json.NewDecoder(r.Body).Decode(&todo)

	stmt, err := database.DB.Prepare("update todos set todo=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(todo.Todo, todoId)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(todo)
	// fmt.Fprintf(w, "todo updated successfullys")
}

func updateTodoStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	var todo TODO
	json.NewDecoder(r.Body).Decode(&todo)

	stmt, err := database.DB.Prepare("update todos set isDone=? WHERE id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(todo.IsDone, todoId)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(todo.Id)

	// fmt.Fprintf(w, "todo  status updated successfullys")
}

func main() {
	database.Connect()
	fmt.Println("listening to 8080")
	router := mux.NewRouter()

	  go handleMessages()

	// go func() {
	// 	todos :=[]TODO{
	// 		{Id: 1, Todo: "task", IsDone: false},
	// 	}
		
		
	// }()
	// ro.Handle("/todo", getTodos)
	router.HandleFunc("/todo", createTodos).Methods("POST")
	router.HandleFunc("/todo/{id}", deleteTodos).Methods("DELETE")
	router.HandleFunc("/todo/{id}", updateTodos).Methods("PUT")
	router.HandleFunc("/todo/status/{id}", updateTodoStatus).Methods("PUT")
	router.HandleFunc("/ws", handleWebSocket)
	go broadcastToClients()
	c := cors.AllowAll().Handler(router)
	
	http.ListenAndServe(":8080", c)
}
