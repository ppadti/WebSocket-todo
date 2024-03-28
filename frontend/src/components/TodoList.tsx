import React, { useState, useEffect } from "react";
import "./App.css";
import { Todo } from "../Models";
import { Button, Container, List, TextField } from "@mui/material";
import SingleTodo from "./SingleTodo";

const TodoList: React.FC = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [inputValue, setInputValue] = useState<string>("");

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data.toString());
      console.log("message", message);
      setTodos(message);
    };
  }, []);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (inputValue) {
      const ws = new WebSocket("ws://localhost:8080/ws");

      ws.onopen = () => {
        ws.send(JSON.stringify({ action: "add", todo: inputValue }));
      };

      setInputValue("");
    }
  };

  const handleDelete = (id: number) => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = () => {
      ws.send(JSON.stringify({ action: "delete", id }));
    };
  };

  const handleEdit = (id: number, newTodo: string) => {
    if (newTodo !== null) {
      const ws = new WebSocket("ws://localhost:8080/ws");

      ws.onopen = () => {
        ws.send(
          JSON.stringify({
            action: "edit",
            id,
            todo: newTodo,
          })
        );
      };
    }
  };

  return (
    <Container maxWidth="sm">
      <h1>Todo List</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <TextField
            label="Todo"
            value={inputValue}
            onChange={handleInputChange}
            variant="outlined"
            margin="normal"
            sx={{
              padding: "2px",
            }}
          />
        </div>
        <div>
          <Button
            type="submit"
            variant="contained"
            sx={{
              padding: "5px",
              marginBottom: "50px",
            }}
          >
            Add Todo
          </Button>
        </div>
      </form>
      <List>
        {todos &&
          todos.map((todo) => (
            <SingleTodo
              key={todo.id}
              todo={todo}
              handleDelete={handleDelete}
              handleUpdate={handleEdit}
            />
          ))}
      </List>
    </Container>
  );
};

export default TodoList;
