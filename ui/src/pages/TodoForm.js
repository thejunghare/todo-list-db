import React, { useState } from "react";
import axios from "axios";

const TodoForm = () => {
  const [todo, setTodo] = useState("");

  const handleaddtodo = async () => {
    try {
      const response = await axios.post("http://localhost:8080/create", {
        name: todo,
      });
      if (response.status === 201) {
        localStorage.setItem("token", response.data.token);
        alert("Todo added successfully!");
      } else {
        alert("Failed to add todo. Status code: " + response.status);
      }
    } catch (error) {
      alert("An error occurred: " + error.message);
    }
  };

  return (
    <div>
      <h1>Add Todo</h1>
      <input
        type="text"
        name="item"
        placeholder="Todo name"
        value={todo}
        onChange={(e) => setTodo(e.target.value)}
      />
      <button onClick={handleaddtodo}>Add</button>
    </div>
  );
};

export default TodoForm;
