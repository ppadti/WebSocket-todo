import React from "react";
import "./App.css";
import TodoList from "./components/TodoList";

const App: React.FC = () => {
  const ws = new WebSocket("ws://localhost:8080/ws");

  ws.onopen = () => {
    ws.send(JSON.stringify({ action: "get" }));
    ws.close();
  };

  return (
    <div className="App">
      <TodoList />
    </div>
  );
};

export default App;
