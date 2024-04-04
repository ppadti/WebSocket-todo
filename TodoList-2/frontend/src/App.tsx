import "@patternfly/react-core/dist/styles/base.css";
import axios from "axios";
import React, { useEffect, useState } from "react";
import AddToForm from "./Components/AddToForm";
import { Todo } from "./Components/Model";
import SingleTodo from "./Components/SingleTodo";
import { w3cwebsocket } from "websocket";
import { Socket } from "dgram";
// import useWebsocket from "./hooks/useWebsocket";

type AddTodo = (text: string) => void;

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");

    ws.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log("received msg:", todos);
      setTodos(message);
    };

    // return () => {
    //   ws.onmessage = null;
    // };
  }, []);

  // async function fetchData() {
  //   const response = await axios.get("http://localhost:8080/todo");

  //   setTodos(response.data);
  // }

  const handleAdd: AddTodo = (todo: string) => {
    if (todo) {
      // const id = Math.floor(Math.random() * 1000);
      // console.log(id, todo);
      axios
        .post("http://localhost:8080/todo", {
          todo: todo,
        })
        .then((response: any) => {
          const ws = new WebSocket("ws://localhost:8080/ws");
          console.log(response.data);
          ws.onopen = () => {
            ws.send(JSON.stringify(response.data));
            // ws.close();
          };
          // console.log(response.data);
          // sendMessage(response.data);
          // fetchData().then((response) => {});
          // Handle data
        })
        .catch((error: any) => {
          console.log(error);
        });
    }
  };

  const handleDone = async (id: number, isDone: boolean, todo: string) => {
    // setTodos(
    //   todos.map((todo) =>
    //     todo.id === id ? { ...todo, isDone: !todo.isDone } : todo,
    //   ),
    // )
    await axios
      .put(`http://localhost:8080/todo/status/${id}`, { isDone: !isDone })
      .then((response) => {
        const ws = new WebSocket("ws://localhost:8080/ws");

        ws.onopen = () => {
          ws.send(
            JSON.stringify({
              id,
              todo,
              isDone,
            })
          );
          // ws.close();
        };
        // console.log(response);
        // fetchData();
        // Handle data
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleDelete = async (id: number, todo: string) => {
    await axios
      .delete(`http://localhost:8080/todo/${id}`)
      .then((response) => {
        const ws = new WebSocket("ws://localhost:8080/ws");
        console.log({
          id,
          todo,
          isDone: response.data.isDone,
        });
        ws.onopen = () => {
          ws.send(
            JSON.stringify({
              id,
              todo,
              isDone: false,
            })
          );
          // ws.close();
        };
        console.log(response);
        // fetchData();
        // Handle data
      })
      .catch((error) => {
        console.log(error);
      });

    // setTodos(todos.filter((todo) => todo.id !== id))
  };

  const handleEdit = async (id: number, editTodo: string) => {
    // setTodos(
    //   todos.map((todo) =>
    //     todo.id === id ? { ...todo, todo: editTodo } : todo,
    //   ),
    // )
    await axios
      .put(`http://localhost:8080/todo/${id}`, { todo: editTodo })
      .then((response: any) => {
        const ws = new WebSocket("ws://localhost:8080/ws");

        ws.onopen = () => {
          ws.send(
            JSON.stringify({
              id,
              todo: editTodo,
              isDone: false,
            })
          );
          // ws.close();
        };
        console.log(response);
        // fetchData();
        // Handle data
      })
      .catch((error: any) => {
        console.log(error);
      });
  };
  // useEffect(() => {
  //   // fetchData();
  // }, []);

  return (
    <>
      <AddToForm handleAdd={handleAdd} />
      {todos &&
        todos.map((todo) => (
          <SingleTodo
            todo={todo}
            key={todo.id}
            handleEdit={handleEdit}
            handleDone={handleDone}
            handleDelete={handleDelete}
          />
        ))}
    </>
  );
}

export default App;
